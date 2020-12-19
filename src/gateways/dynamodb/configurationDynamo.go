package dynamodb

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/halprin/email-conceal/src/context"
	"github.com/halprin/email-conceal/src/external/lib/errors"
	"github.com/halprin/email-conceal/src/usecases"
	"log"
	"time"
)


var applicationContext = context.ApplicationContext{}

type DynamoDbGateway struct {}

var dynamoService = dynamodb.New(awsSession)

type KeyBase struct {
	Primary   string  `dynamodbav:"primary"`
	Secondary string `dynamodbav:"secondary"`
}

type ConcealEmailEntity struct {
	KeyBase
	Description *string `dynamodbav:"description"`
}

type ConcealEmailMapping struct {
	KeyBase
}

func (receiver DynamoDbGateway) AddConcealedEmailToActualEmailMapping(concealPrefix string, actualEmail string, description *string) error {
	if sessionErr != nil {
		return errors.Wrap(sessionErr, "Error with the AWS session")
	}

	var concealDynamoDbKey = generateConcealEmailKey(concealPrefix)
	//the primary entity
	entity := ConcealEmailEntity{
		Description: description,
	}
	entity.Primary = concealDynamoDbKey
	entity.Secondary = concealDynamoDbKey

	//the mapping data for the conceal entity
	mapping := ConcealEmailMapping{}
	mapping.Primary = concealDynamoDbKey
	mapping.Secondary = generateSourceEmailKey(actualEmail)

	rollbackFromNewConceal := func() {
		_ = batchDeleteItemsWithRollback([]interface{}{entity, mapping}, nil)
	}

	log.Println("Writing new conceal mapping to DynamoDB")
	return batchWriteItemsWithRollback([]interface{}{entity, mapping}, rollbackFromNewConceal)
}

func (receiver DynamoDbGateway) DeleteConcealedEmailToActualEmailMapping(concealPrefix string) error {
	if sessionErr != nil {
		return errors.Wrap(sessionErr, "Error with the AWS session")
	}

	var environmentGateway context.EnvironmentGateway
	applicationContext.Resolve(&environmentGateway)
	tableName := environmentGateway.GetEnvironmentValue("TABLE_NAME")

	items, err := getAllItemsForHashKey(generateConcealEmailKey(concealPrefix), tableName)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed to get all the items for the hash key %s", concealPrefix))
	}

	for _, item := range items {
		key, err := convertItemToKey(item)
		if err != nil {
			return errors.Wrap(err, "Failed to delete item in DynamoDB due to unable to get key from item")
		}

		deleteItemInput := &dynamodb.DeleteItemInput{
			TableName: aws.String(tableName),
			Key:       key,
		}

		_, err = dynamoService.DeleteItem(deleteItemInput)
		if err != nil {
			return errors.Wrap(err, "Failed to delete item in DynamoDB")
		}
	}

	return nil
}

func (receiver DynamoDbGateway) UpdateConcealedEmail(concealPrefix string, description *string) error {
	if sessionErr != nil {
		return errors.Wrap(sessionErr, "Error with the AWS session")
	}

	var environmentGateway context.EnvironmentGateway
	applicationContext.Resolve(&environmentGateway)
	tableName := environmentGateway.GetEnvironmentValue("TABLE_NAME")

	concealEmailKey := generateConcealEmailKey(concealPrefix)
	item, err := getItem(concealEmailKey, concealEmailKey, tableName)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Failed to get conceal e-mail %s to update it", concealPrefix))
	} else if item == nil {
		return usecases.ConcealEmailNotExistError{
			ConcealEmailId: concealPrefix,
		}
	}

	keyEntity := KeyBase{
		Primary:   concealEmailKey,
		Secondary: concealEmailKey,
	}

	dynamoKeyMapping, err := dynamodbattribute.MarshalMap(keyEntity)
	if err != nil {
		return errors.Wrap(err, "Failed to marshal conceal e-mail key")
	}

	updateExpressionBuilder := expression.Set(expression.Name("description"), expression.Value(description))
	expressionBuilder, err := expression.NewBuilder().WithUpdate(updateExpressionBuilder).Build()
	if err != nil {
		return errors.Wrap(err, "Failed to make update expression")
	}


	updateItemInput := dynamodb.UpdateItemInput{
		TableName:                 aws.String(tableName),
		Key:                       dynamoKeyMapping,
		ExpressionAttributeNames:  expressionBuilder.Names(),
		ExpressionAttributeValues: expressionBuilder.Values(),
		UpdateExpression:          expressionBuilder.Update(),
	}

	log.Println("Updating conceal e-mail to DynamoDB")
	_, err = dynamoService.UpdateItem(&updateItemInput)
	if err != nil {
		return errors.Wrap(err, "Failed to update item in DynamoDB")
	}

	return nil
}

func convertItemToKey(item map[string]*dynamodb.AttributeValue) (map[string]*dynamodb.AttributeValue, error) {
	var keyEntity KeyBase
	err := dynamodbattribute.UnmarshalMap(item, &keyEntity)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to convert item to key due to unmarshalling")
	}

	key, err := dynamodbattribute.MarshalMap(keyEntity)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to convert item to key due to marshalling")
	}

	return key, nil
}

var concealEmailKeyPrefix = "conceal#"
var sourceEmailKeyPrefix = "email#"

func generateConcealEmailKey(concealPrefix string) string {
	return fmt.Sprintf("%s%s", concealEmailKeyPrefix, concealPrefix)
}

func generateSourceEmailKey(sourceEmail string) string {
	return fmt.Sprintf("%s%s", sourceEmailKeyPrefix, sourceEmail)
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

func getItem(hashKey string, sortKey string, tableName string) (map[string]*dynamodb.AttributeValue, error) {
	keyEntity := KeyBase{
		Primary: hashKey,
		Secondary: sortKey,
	}

	key, err := dynamodbattribute.MarshalMap(keyEntity)
	if err != nil {
		return nil, err
	}

	getInput := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       key,
	}

	getOutput, err := dynamoService.GetItem(getInput)
	if err != nil {
		return nil, err
	}

	return getOutput.Item, nil
}

func batchWriteItemsWithRollback(structsToWrite []interface{}, rollbackFunction func()) error {
	log.Println("Batch writing items")
	return batchInternal(structsToWrite, rollbackFunction, batchWrite)
}

func batchDeleteItemsWithRollback(structsToDelete []interface{}, rollbackFunction func()) error {
	log.Println("Batch deleting items")
	return batchInternal(structsToDelete, rollbackFunction, batchDelete)
}

const (
	batchWrite = "batchWrite"
	batchDelete = "batchDelete"
)

func batchInternal(structsToWrite []interface{}, rollbackFunction func(), batchOperation string) error {
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
			key, err := convertItemToKey(dynamoItem)
			if err != nil {
				//return immediately without running the rollback function because we haven't even made a single DynamoDB call yet
				return err
			}

			deleteRequest := &dynamodb.DeleteRequest{
				Key:  key,
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
				go rollbackFunction()
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
