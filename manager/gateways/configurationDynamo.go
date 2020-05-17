package gateways

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/halprin/email-conceal/manager/context"
)

var dynamoService = dynamodb.New(awsSession)

type ConcealEmailMapping struct {
	Primary string
	Secondary string
}

func AddConcealedEmailToActualEmailMapping(concealPrefix string, actualEmail string, applicationContext context.ApplicationContext) error {
	if sessionErr != nil {
		return sessionErr
	}

	goItem := ConcealEmailMapping{
		Primary:   fmt.Sprintf("conceal-%s", concealPrefix),
		Secondary: fmt.Sprintf("email-%s", actualEmail),
	}

	dynamoItem, err := dynamodbattribute.MarshalMap(goItem)
	if err != nil {
		return err
	}

	putItemInput := &dynamodb.PutItemInput{
		TableName: aws.String(applicationContext.EnvironmentGateway("TABLE_NAME")),
		Item:      dynamoItem,
	}
	_, err = dynamoService.PutItem(putItemInput)
	if err != nil {
		return err
	}

	return nil
}

func DeleteConcealEmailToActualEmailMapping(concealPrefix string, actualEmail string, applicationContext context.ApplicationContext) error {
	if sessionErr != nil {
		return sessionErr
	}

	goItem := ConcealEmailMapping{
		Primary:   fmt.Sprintf("conceal-%s", concealPrefix),
		Secondary: fmt.Sprintf("email-%s", actualEmail),
	}

	dynamoItem, err := dynamodbattribute.MarshalMap(goItem)
	if err != nil {
		return err
	}

	deleteItemInput := &dynamodb.DeleteItemInput{
		TableName: aws.String(applicationContext.EnvironmentGateway("TABLE_NAME")),
		Key:       dynamoItem,
	}
	_, err = dynamoService.DeleteItem(deleteItemInput)
	if err != nil {
		return err
	}

	return nil
}
