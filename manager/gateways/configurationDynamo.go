package gateways

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/halprin/email-conceal/manager/context"
	"github.com/halprin/email-conceal/manager/external/lib/errors"
)

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

func AddConcealedEmailToActualEmailMapping(concealPrefix string, actualEmail string, description *string, applicationContext context.ApplicationContext) error {
	if sessionErr != nil {
		return errors.Wrap(sessionErr, "Error with the AWS session")
	}

	tableName := aws.String(applicationContext.Gateways().GetEnvironmentValue("TABLE_NAME"))

	//write the primary entity

	entity := ConcealEmailEntity{
		Primary:   fmt.Sprintf("conceal-%s", concealPrefix),
		Secondary: fmt.Sprintf("conceal-%s", concealPrefix),
		Description: description,
	}
	dynamoEntity, err := dynamodbattribute.MarshalMap(entity)
	if err != nil {
		return errors.Wrap(err, "Failed to marshal struct into a DynamoDB item")
	}

	putItemInput := &dynamodb.PutItemInput{
		TableName: tableName,
		Item:      dynamoEntity,
	}
	_, err = dynamoService.PutItem(putItemInput)
	if err != nil {
		return errors.Wrap(err, "Failed to put entity in DynamoDB")
	}

	//write the mapping data for the conceal entity
	mapping := ConcealEmailMapping{
		Primary:   fmt.Sprintf("conceal-%s", concealPrefix),
		Secondary: fmt.Sprintf("email-%s", actualEmail),
	}
	dynamoMapping, err := dynamodbattribute.MarshalMap(mapping)
	if err != nil {
		return errors.Wrap(err, "Failed to marshal struct into a DynamoDB item")
	}

	putItemInput = &dynamodb.PutItemInput{
		TableName: tableName,
		Item:      dynamoMapping,
	}
	_, err = dynamoService.PutItem(putItemInput)
	if err != nil {
		return errors.Wrap(err, "Failed to put item in DynamoDB")
	}

	return nil
}

func DeleteConcealedEmailToActualEmailMapping(concealPrefix string, applicationContext context.ApplicationContext) error {
	if sessionErr != nil {
		return errors.Wrap(sessionErr, "Error with the AWS session")
	}

	tableName := applicationContext.Gateways().GetEnvironmentValue("TABLE_NAME")

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
