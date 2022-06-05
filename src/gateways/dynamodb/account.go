package dynamodb

import (
	"fmt"
	"github.com/halprin/email-conceal/src/external/lib/errors"
	"log"
)

type AccountEntity struct {
	KeyBase
	Password string `dynamodbav:"password"`
}

func (receiver DynamoDbGateway) AddAccount(emailUsername string, passwordHash string) error {
	if sessionErr != nil {
		return errors.Wrap(sessionErr, "Error with the AWS session")
	}

	actualEmailEntity := AccountEntity{
		Password: passwordHash,
	}
	actualEmailEntity.Primary = generateAccountKey(emailUsername)
	actualEmailEntity.Secondary = generateAccountKey(emailUsername)

	rollbackFromNewActual := func() {
		_ = batchDeleteItemsWithRollback([]interface{}{actualEmailEntity}, nil)
	}

	log.Printf("Writing new account, %s, to DynamoDB", emailUsername)
	return batchWriteItemsWithRollback([]interface{}{actualEmailEntity}, rollbackFromNewActual)
}

var accountPrefix = "account#"

func generateAccountKey(username string) string {
	return fmt.Sprintf("%s%s", accountPrefix, username)
}
