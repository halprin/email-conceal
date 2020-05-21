package controllers

import (
	"fmt"
	"github.com/halprin/email-conceal/forwarder/context"
	"github.com/halprin/email-conceal/forwarder/external/lib/errors"
	"testing"
)

func TestSqsQueueForwardEmailFailsJsonParsing(t *testing.T) {
	appContext := context.TestApplicationContext{}

	message := `{"moof": "dogcow}`
	arguments := map[string]interface{}{
		"message": &message,  //no ending quote
	}

	err := SqsQueueForwardEmail(arguments, &appContext)

	if err == nil {
		t.Error("An error should have been returned from SqsQueueForwardEmail controllers")
	}

	if appContext.ReceivedForwardEmailUsecaseArguments != "" {
		t.Errorf("The forward e-mail usecase was called when it shouldn't have")
	}
}

func TestSqsQueueForwardEmailFailsJsonNotBeingAsExpected(t *testing.T) {
	appContext := context.TestApplicationContext{}

	key := "an_object.txt"
	message := fmt.Sprintf( //no bucket key
`{
	"Records": [{
		"s3": {
			"asdf": {
				"name": "a_bucket"
			},
			"object": {
				"key": "%s"
			}
		}
	}]
}`, key)
	expectedUrl := fmt.Sprintf("s3:///%s", key) // the lack of bucket because the bucket key was not correct in the above JSON

	arguments := map[string]interface{}{
		"message": &message,
	}

	err := SqsQueueForwardEmail(arguments, &appContext)

	if err != nil {
		t.Error("An error shouldn't have been returned from SqsQueueForwardEmail controllers")
	}

	if appContext.ReceivedForwardEmailUsecaseArguments != expectedUrl {
		t.Errorf("The forward e-mail usecase didn't have the expected URL passed to it; expted %s", expectedUrl)
	}
}

func TestSqsQueueForwardEmailFailsTheUsecase(t *testing.T) {
	appContext := context.TestApplicationContext{}
	expectedErrorFromUsecase := errors.New("oops")
	appContext.ReturnErrorForwardEmailUsecase = expectedErrorFromUsecase

	bucket := "a_bucket"
	key := "an_object.txt"
	message := fmt.Sprintf(
`{
	"Records": [{
		"s3": {
			"bucket": {
				"name": "%s"
			},
			"object": {
				"key": "%s"
			}
		}
	}]
}`, bucket, key)
	expectedUrl := fmt.Sprintf("s3://%s/%s", bucket, key)

	arguments := map[string]interface{}{
		"message": &message,
	}

	err := SqsQueueForwardEmail(arguments, &appContext)

	if err != expectedErrorFromUsecase {
		t.Error("A specific error should have been returned from SqsQueueForwardEmail controllers")
	}

	if appContext.ReceivedForwardEmailUsecaseArguments != expectedUrl {
		t.Errorf("The forward e-mail usecase didn't have the expected URL passed to it; expted %s", expectedUrl)
	}
}

func TestSqsQueueForwardEmailIsSuccess(t *testing.T) {
	appContext := context.TestApplicationContext{}

	bucket := "a_bucket"
	key := "an_object.txt"
	message := fmt.Sprintf(
		`{
	"Records": [{
		"s3": {
			"bucket": {
				"name": "%s"
			},
			"object": {
				"key": "%s"
			}
		}
	}]
}`, bucket, key)
	expectedUrl := fmt.Sprintf("s3://%s/%s", bucket, key)

	arguments := map[string]interface{}{
		"message": &message,
	}

	err := SqsQueueForwardEmail(arguments, &appContext)

	if err != nil {
		t.Error("An  error shouldn't have been returned from SqsQueueForwardEmail controllers")
	}

	if appContext.ReceivedForwardEmailUsecaseArguments != expectedUrl {
		t.Errorf("The forward e-mail usecase didn't have the expected URL passed to it; expted %s", expectedUrl)
	}
}
