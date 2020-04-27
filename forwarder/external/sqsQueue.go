package external

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
)

func SqsQueueListener() {
	sqsService := createSqsService()

	queueUrl := getQueueUrl(sqsService)

	//start listening
	listenToQueue(sqsService, queueUrl)
}

func listenToQueue(sqsService *sqs.SQS, queueUrl *string) {
	for {
		log.Print("Calling SQS ReceiveMessage")
		queueMessages, err := sqsService.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl: queueUrl,
			MaxNumberOfMessages: aws.Int64(5),
			WaitTimeSeconds: aws.Int64(5),
		})
		if err != nil {
			log.Printf("AWS SQS queue messages weren't able to be retrieved; %+v", err)
		}

		for _, message := range queueMessages.Messages {
			go handleQueueMessage(message, sqsService, queueUrl)
		}
	}
}

func handleQueueMessage(message *sqs.Message, sqsService *sqs.SQS, queueUrl *string) {
	log.Print("Handling queue message")
	log.Print(*message.Body)

	//delete the message
	_, err := sqsService.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      queueUrl,
		ReceiptHandle: message.ReceiptHandle,
	})

	if err != nil {
		log.Printf("Failed to delete handled message")
	}
}

func getQueueUrl(sqsService *sqs.SQS) *string {
	queueUrlResult, err := sqsService.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String("email-conceal-forwarder-dev"),
	})

	if err != nil {
		log.Fatalf("AWS SQS queue URL wasn't able to be gotten; %+v", err)
	}

	return queueUrlResult.QueueUrl
}

func createSqsService() *sqs.SQS {
	awsSession, err := session.NewSession()

	if err != nil {
		log.Fatalf("AWS SQS Session failed to contruct; %+v", err)
	}

	return sqs.New(awsSession)
}
