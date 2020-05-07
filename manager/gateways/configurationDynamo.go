package gateways

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/halprin/email-conceal/manager/context"
	"github.com/halprin/email-conceal/manager/external/lib/errors"
	"strings"
)

var dynamoService = dynamodb.New(awsSession)

type ConcealEmailMapping struct {
	Primary string
	Secondary string
}

func AddConcealedEmailToActualEmailMapping(concealPrefix string, actualEmail string, applicationContext context.ApplicationContext) (string, error) {
	if sessionErr != nil {
		return "", sessionErr
	}

	putItemInput := &dynamodb.PutItemInput{
		TableName:                 aws.String(applicationContext.EnvironmentGateway("TABLE_NAME")),
		KeyConditionExpression:    expressionBuilder.KeyCondition(),
		ExpressionAttributeNames:  expressionBuilder.Names(),
		ExpressionAttributeValues: expressionBuilder.Values(),
	}
	queryOutput, err := dynamoService.PutItem(putItemInput)
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

	return strings.TrimPrefix(item.Secondary, "email-"), nil
}
