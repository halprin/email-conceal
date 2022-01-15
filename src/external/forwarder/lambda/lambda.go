package lambda

import (
	goContext "context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/halprin/email-conceal/src/context"
	"github.com/halprin/email-conceal/src/controllers/forwardEmail"
	"log"
)

var applicationContext = context.ApplicationContext{}

func LambdaHandler(ctx goContext.Context, req events.SQSEvent) (events.SQSEventResponse, error) {
	var handleMessageFailures []events.SQSBatchItemFailure

	for _, sqsMessage := range req.Records {
		err := handleQueueMessage(sqsMessage)
		if err != nil {
			handleMessageFailures = append(handleMessageFailures, events.SQSBatchItemFailure{
				ItemIdentifier: sqsMessage.MessageId,
			})
		}
	}

	return events.SQSEventResponse{
		BatchItemFailures: handleMessageFailures,
	}, nil
}

func handleQueueMessage(message events.SQSMessage) error {
	log.Printf("Handling queue message %s", message.MessageId)

	arguments := map[string]interface{}{
		"message": message.Body,
	}

	var forwardEmailController forwardEmail.ForwardEmail
	applicationContext.Resolve(&forwardEmailController)

	err := forwardEmailController.ForwardEmail(arguments)
	if err != nil {
		//don't delete the message since we weren't able to handle it
		log.Printf("Failed to handle the message %s", message.MessageId)
		return err
	}

	return nil
}
