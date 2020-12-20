package dynamodb

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/halprin/email-conceal/src/context"
	"github.com/halprin/email-conceal/src/external/lib/errors"
	"github.com/halprin/email-conceal/src/usecases/concealEmail"
	"log"
	"strings"
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
		return concealEmail.ConcealEmailNotExistError{
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

func (receiver DynamoDbGateway) GetRealEmailAddressForConcealPrefix(concealPrefix string) (string, error) {
	if sessionErr != nil {
		return "", sessionErr
	}

	keyCondition := expression.Key("primary").Equal(expression.Value(fmt.Sprintf("conceal#%s", concealPrefix))).And(expression.Key("secondary").BeginsWith("email#"))
	keyBuilder := expression.NewBuilder().WithKeyCondition(keyCondition)
	expressionBuilder, err := keyBuilder.Build()
	if err != nil {
		return "", err
	}

	var environmentGateway context.EnvironmentGateway
	applicationContext.Resolve(&environmentGateway)

	queryInput := &dynamodb.QueryInput{
		TableName:                 aws.String(environmentGateway.GetEnvironmentValue("TABLE_NAME")),
		KeyConditionExpression:    expressionBuilder.KeyCondition(),
		ExpressionAttributeNames:  expressionBuilder.Names(),
		ExpressionAttributeValues: expressionBuilder.Values(),
	}
	queryOutput, err := dynamoService.Query(queryInput)
	if err != nil {
		return "", err
	}

	if *queryOutput.Count < 1 {
		return "", errors.New(fmt.Sprintf("No real e-mail for conceal prefix %s", concealPrefix))
	}

	firstItem := queryOutput.Items[0]
	item := ConcealEmailMapping{}
	err = dynamodbattribute.UnmarshalMap(firstItem, &item)
	if err != nil {
		return "", err
	}

	return strings.TrimPrefix(item.Secondary, "email#"), nil
}
