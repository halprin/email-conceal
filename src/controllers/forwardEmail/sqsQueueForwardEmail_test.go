package forwardEmail

import (
	"fmt"
	"github.com/halprin/email-conceal/src/context"
	"github.com/halprin/email-conceal/src/external/lib/errors"
	"github.com/halprin/email-conceal/src/usecases/forwardEmail"
	"testing"
)

var sqsController = SqsQueueForwardController{}
var testAppContext = context.ApplicationContext{}

type TestForwardEmailUsecase struct {
	ForwardEmailUri         string
	ForwardEmailReturnError error
}

func (testUsecase *TestForwardEmailUsecase) ForwardEmail(url string) error {
	testUsecase.ForwardEmailUri = url

	return testUsecase.ForwardEmailReturnError
}

func TestSqsQueueForwardEmailFailsJsonParsing(t *testing.T) {

	message := `{"moof": "dogcow}`
	arguments := map[string]interface{}{
		"message": message, //no ending quote
	}

	testForwardEmailUsecase := TestForwardEmailUsecase{}
	testAppContext.Bind(func() forwardEmail.ForwardEmailUsecase {
		return &testForwardEmailUsecase
	})

	err := sqsController.ForwardEmail(arguments)

	if err == nil {
		t.Error("An error should have been returned from SqsQueueForwardEmail controller")
	}

	if testForwardEmailUsecase.ForwardEmailUri != "" {
		t.Errorf("The forward e-mail usecase was called when it shouldn't have")
	}
}

func TestSqsQueueForwardEmailFailsJsonNotBeingAsExpected(t *testing.T) {

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

	testForwardEmailUsecase := TestForwardEmailUsecase{}
	testAppContext.Bind(func() forwardEmail.ForwardEmailUsecase {
		return &testForwardEmailUsecase
	})

	arguments := map[string]interface{}{
		"message": message,
	}

	err := sqsController.ForwardEmail(arguments)

	if err != nil {
		t.Error("An error shouldn't have been returned from SqsQueueForwardEmail controller")
	}

	if testForwardEmailUsecase.ForwardEmailUri != expectedUrl {
		t.Errorf("The forward e-mail usecase didn't have the expected URL passed to it; expted %s", expectedUrl)
	}
}

func TestSqsQueueForwardEmailFailsTheUsecase(t *testing.T) {

	expectedErrorFromUsecase := errors.New("oops")
	testForwardEmailUsecase := TestForwardEmailUsecase{
		ForwardEmailReturnError: expectedErrorFromUsecase,
	}
	testAppContext.Bind(func() forwardEmail.ForwardEmailUsecase {
		return &testForwardEmailUsecase
	})

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
		"message": message,
	}

	err := sqsController.ForwardEmail(arguments)

	if err != expectedErrorFromUsecase {
		t.Error("A specific error should have been returned from SqsQueueForwardEmail controller")
	}

	if testForwardEmailUsecase.ForwardEmailUri != expectedUrl {
		t.Errorf("The forward e-mail usecase didn't have the expected URL passed to it; expted %s", expectedUrl)
	}
}

func TestSqsQueueForwardEmailIsSuccess(t *testing.T) {

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

	testForwardEmailUsecase := TestForwardEmailUsecase{}
	testAppContext.Bind(func() forwardEmail.ForwardEmailUsecase {
		return &testForwardEmailUsecase
	})

	arguments := map[string]interface{}{
		"message": message,
	}

	err := sqsController.ForwardEmail(arguments)

	if err != nil {
		t.Error("An  error shouldn't have been returned from SqsQueueForwardEmail controller")
	}

	if testForwardEmailUsecase.ForwardEmailUri != expectedUrl {
		t.Errorf("The forward e-mail usecase didn't have the expected URL passed to it; expted %s", expectedUrl)
	}
}
