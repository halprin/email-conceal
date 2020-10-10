package gateways

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/halprin/email-conceal/manager/context"
	"github.com/halprin/email-conceal/manager/external/lib/errors"
	"log"
	"time"
)


var applicationContext = context.ApplicationContext{}

type DynamoDbGateway struct {}

var dynamoService = dynamodb.New(awsSession)

type ConcealEmailMapping struct {
	Primary   string  `dynamodbav:"primary"`
	Secondary string `dynamodbav:"secondary"`
}

type ConcealEmailEntity struct {
	Primary     string `dynamodbav:"primary"`
	Secondary   string `dynamodbav:"secondary"`
	Description *string `dynamodbav:"description"`
}

func (receiver DynamoDbGateway) AddConcealedEmailToActualEmailMapping(concealPrefix string, actualEmail string, description *string) error {
	if sessionErr != nil {
		return errors.Wrap(sessionErr, "Error with the AWS session")
	}

	//the primary entity
	entity := ConcealEmailEntity{
		Primary:   fmt.Sprintf("conceal-%s", concealPrefix),
		Secondary: fmt.Sprintf("conceal-%s", concealPrefix),
		Description: description,
	}

	//the mapping data for the conceal entity
	mapping := ConcealEmailMapping{
		Primary:   fmt.Sprintf("conceal-%s", concealPrefix),
		Secondary: fmt.Sprintf("email-%s", actualEmail),
	}

	rollbackFromNewConceal := func(applicationContext context.ApplicationContext) {
		_ = batchDeleteItemsWithRollback([]interface{}{entity, mapping}, nil, applicationContext)
	}

	log.Println("Writing new conceal mapping to DynamoDB")
	return batchWriteItemsWithRollback([]interface{}{entity, mapping}, rollbackFromNewConceal, applicationContext)
}

func (receiver DynamoDbGateway) DeleteConcealedEmailToActualEmailMapping(concealPrefix string) error {
	if sessionErr != nil {
		return errors.Wrap(sessionErr, "Error with the AWS session")
	}

	var environmentGateway context.EnvironmentGateway
	applicationContext.Resolve(&environmentGateway)
	tableName := environmentGateway.GetEnvironmentValue("TABLE_NAME")

	items, err := getAllItemsForHashKey(fmt.Sprintf("conceal-%s", concealPrefix), tableName)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed to get all the items for the hash key %s", concealPrefix))
	}

	for _, item := range items {
		deleteItemInput := &dynamodb.DeleteItemInput{
			TableName: aws.String(tableName),
			Key:       item,
		}

		_, err = dynamoService.DeleteItem(deleteItemInput)
		if err != nil {
			return errors.Wrap(err, "Failed to delete item in DynamoDB")
		}
	}

	return nil
}

func getAllItemsForHashKey(hashKey string, tableName string) ([]map[string]*dynamodb.AttributeValue, error) {
	keyCondition := expression.Key("primary").Equal(expression.Value(hashKey))
	keyBuilder := expression.NewBuilder().WithKeyCondition(keyCondition)
	expressionBuilder, err := keyBuilder.Build()
	if err != nil {
		return nil, err
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		KeyConditionExpression:    expressionBuilder.KeyCondition(),
		ExpressionAttributeNames:  expressionBuilder.Names(),
		ExpressionAttributeValues: expressionBuilder.Values(),
	}
	queryOutput, err := dynamoService.Query(queryInput)
	if err != nil {
		return nil, err
	}

	return queryOutput.Items, nil
}

func batchWriteItemsWithRollback(structsToWrite []interface{}, rollbackFunction func(context.ApplicationContext), applicationContext context.ApplicationContext) error {
	log.Println("Batch writing items")
	return batchInternal(structsToWrite, rollbackFunction, batchWrite, applicationContext)
}

func batchDeleteItemsWithRollback(structsToDelete []interface{}, rollbackFunction func(context.ApplicationContext), applicationContext context.ApplicationContext) error {
	log.Println("Batch deleting items")
	return batchInternal(structsToDelete, rollbackFunction, batchDelete, applicationContext)
}

const (
	batchWrite = "batchWrite"
	batchDelete = "batchDelete"
)

func batchInternal(structsToWrite []interface{}, rollbackFunction func(context.ApplicationContext), batchOperation string, applicationContext context.ApplicationContext) error {
	//convert the structs to dynamo attribute maps
	var dynamoItems []map[string]*dynamodb.AttributeValue

	for _, structToWrite := range structsToWrite {
		dynamoMapping, err := dynamodbattribute.MarshalMap(structToWrite)
		if err != nil {
			//return immediately without running the rollback function because we haven't even made a single DynamoDB call yet
			return err
		}

		dynamoItems = append(dynamoItems, dynamoMapping)
	}

	//convert the dynamo attribute maps to write requests (structs needed by the WriteBatchItem API)
	var writeRequests []*dynamodb.WriteRequest
	var writeRequest *dynamodb.WriteRequest

	for _, dynamoItem := range dynamoItems {
		if batchOperation == batchWrite {
			putRequest := &dynamodb.PutRequest{
				Item:  dynamoItem,
			}

			writeRequest = &dynamodb.WriteRequest{
				PutRequest: putRequest,
			}
		} else if batchOperation == batchDelete {
			deleteRequest := &dynamodb.DeleteRequest{
				Key:  dynamoItem,
			}

			writeRequest = &dynamodb.WriteRequest{
				DeleteRequest: deleteRequest,
			}
		}

		writeRequests = append(writeRequests, writeRequest)
	}

	//do last bit of construction for the BatchWriteItem API
	var environmentGateway context.EnvironmentGateway
	applicationContext.Resolve(&environmentGateway)
	tableName := environmentGateway.GetEnvironmentValue("TABLE_NAME")
	requestItems := map[string][]*dynamodb.WriteRequest{
		tableName: writeRequests,
	}

	//loop until all the remaining items have been written
	millisecondsToWait := 20

	for {
		batchWriteItemInput := &dynamodb.BatchWriteItemInput{
			RequestItems: requestItems,
		}

		batchWriteItemOutput, err := dynamoService.BatchWriteItem(batchWriteItemInput)
		if err != nil {
			//there was an error writing to DynamoDB
			log.Println("Failed to put/delete items in DynamoDB")
			if rollbackFunction != nil {
				log.Println("Calling rollback function")
				go rollbackFunction(applicationContext)
			}
			return errors.Wrap(err, "Failed to put/delete items in DynamoDB")
		}

		if len(batchWriteItemOutput.UnprocessedItems) > 0 {
			//there are still items to write, reset requestItems for the next pass
			log.Println("Unprocessed items remain, trying again with remaining items")
			requestItems = batchWriteItemOutput.UnprocessedItems
		} else {
			//no more items to write, break out
			log.Println("Done putting/deleting batch items to DynamoDB")
			break
		}

		//do an exponential back-off
		time.Sleep(time.Duration(millisecondsToWait) * time.Millisecond)
		millisecondsToWait *= 2
	}

	return nil
}
