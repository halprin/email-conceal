package dynamodb

import (
	"github.com/halprin/email-conceal/src/external/lib/errors"
	"log"
)

type ActualEmailEntity struct {
	KeyBase
	OwnershipSecret string `dynamodbav:"ownershipSecret"`
	Active          bool   `dynamodbav:"active"`
}

func (receiver DynamoDbGateway) AddUnprovedActualEmail(actualEmail string, ownershipSecret string) error {
	if sessionErr != nil {
		return errors.Wrap(sessionErr, "Error with the AWS session")
	}

	actualEmailEntity := ActualEmailEntity{
		OwnershipSecret: ownershipSecret,
		Active:          false,
	}

	actualEmailEntity.Primary = generateSourceEmailKey(actualEmail)
	actualEmailEntity.Secondary = generateSourceEmailKey(actualEmail)

	rollbackFromNewActual := func() {
		_ = batchDeleteItemsWithRollback([]interface{}{actualEmailEntity}, nil)
	}

	log.Println("Writing new unproved actual e-mail to DynamoDB")
	return batchWriteItemsWithRollback([]interface{}{actualEmailEntity}, rollbackFromNewActual)
}
