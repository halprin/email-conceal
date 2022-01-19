package dynamodb

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/halprin/email-conceal/src/context"
	"github.com/halprin/email-conceal/src/external/lib/errors"
	"log"
	"strings"
)

type ActualEmailEntity struct {
	KeyBase
	Active bool `dynamodbav:"active"`
}

type SecretToActualEmailMapping struct {
	KeyBase
}

func (receiver DynamoDbGateway) AddUnprovedActualEmail(actualEmail string, ownershipSecret string) error {
	if sessionErr != nil {
		return errors.Wrap(sessionErr, "Error with the AWS session")
	}

	actualEmailEntity := ActualEmailEntity{
		Active: false,
	}
	actualEmailEntity.Primary = generateActualEmailKey(actualEmail)
	actualEmailEntity.Secondary = generateActualEmailKey(actualEmail)

	secretToActualEmailMapping := SecretToActualEmailMapping{}
	secretToActualEmailMapping.Primary = generateOwnershipSecretKey(ownershipSecret)
	secretToActualEmailMapping.Secondary = generateActualEmailKey(actualEmail)

	rollbackFromNewActual := func() {
		_ = batchDeleteItemsWithRollback([]interface{}{actualEmailEntity, secretToActualEmailMapping}, nil)
	}

	log.Printf("Writing new unproved actual e-mail, %s, to DynamoDB", actualEmail)
	return batchWriteItemsWithRollback([]interface{}{actualEmailEntity, secretToActualEmailMapping}, rollbackFromNewActual)
}

func (receiver DynamoDbGateway) GetActualEmailForSecret(secret string) (*string, error) {

	var environmentGateway context.EnvironmentGateway
	applicationContext.Resolve(&environmentGateway)
	tableName := environmentGateway.GetEnvironmentValue("TABLE_NAME")

	primary := generateOwnershipSecretKey(secret)

	log.Println("Looking up an ownership secret to actual e-mail mapping")

	mapping, err := getAllItemsForHashKey(primary, tableName)
	if err != nil {
		return nil, err
	}

	if len(mapping) == 0 {
		//no actual e-mail exists for the secret, return a nil string pointer
		return nil, nil
	}

	firstRawItem := mapping[0]

	var secretToActualEmailMapping SecretToActualEmailMapping
	err = dynamodbattribute.UnmarshalMap(firstRawItem, &secretToActualEmailMapping)
	if err != nil {
		return nil, err
	}

	actualEmail := strings.TrimPrefix(secretToActualEmailMapping.Secondary, actualEmailKeyPrefix)

	return &actualEmail, nil
}

func (receiver DynamoDbGateway) ActivateActualEmail(actualEmail string) error {
	var environmentGateway context.EnvironmentGateway
	applicationContext.Resolve(&environmentGateway)
	tableName := environmentGateway.GetEnvironmentValue("TABLE_NAME")

	actualEmailKey := generateActualEmailKey(actualEmail)

	rawItem, err := getItem(actualEmailKey, actualEmailKey, tableName)
	if err != nil {
		return err
	} else if rawItem == nil {
		return errors.New("Couldn't find actual e-mail in database")
	}

	keyEntity := KeyBase{
		Primary:   actualEmailKey,
		Secondary: actualEmailKey,
	}

	dynamoKeyMapping, err := dynamodbattribute.MarshalMap(keyEntity)
	if err != nil {
		return errors.Wrap(err, "Failed to marshal actual e-mail key")
	}

	updateExpressionBuilder := expression.Set(expression.Name("active"), expression.Value(true))
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

	log.Println("Activating actual e-mail in DynamoDB")
	_, err = dynamoService.UpdateItem(&updateItemInput)
	if err != nil {
		return errors.Wrap(err, "Failed to update item in DynamoDB")
	}

	return nil
}

var ownershipSecretKeyPrefix = "ownershipSecret#"

func generateOwnershipSecretKey(ownershipSecret string) string {
	return fmt.Sprintf("%s%s", ownershipSecretKeyPrefix, ownershipSecret)
}
