package gateways

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/halprin/email-conceal/forwarder/context"
	"github.com/halprin/email-conceal/forwarder/external/lib/errors"
	"strings"
)

var dynamoService = dynamodb.New(awsSession)

type ConcealEmailMapping struct {
	Primary string
	Secondary string
}

func GetRealEmailForConcealPrefix(concealPrefix string, applicationContext context.ApplicationContext) (string, error) {
	keyCondition := expression.Key("primary").Equal(expression.Value(fmt.Sprintf("conceal-%s", concealPrefix))).And(expression.Key("secondary").BeginsWith("email-"))
	keyBuilder := expression.NewBuilder().WithKeyCondition(keyCondition)
	expressionBuilder, err := keyBuilder.Build()
	if err != nil {
		return "", err
	}

	queryInput := &dynamodb.QueryInput{
		TableName:                 aws.String(applicationContext.EnvironmentGateway("TABLE_NAME")),
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

	return strings.TrimPrefix(item.Secondary, "email-"), nil
}
