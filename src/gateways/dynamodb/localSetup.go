package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/halprin/email-conceal/src/context"
	"log"
	"time"
)

func localInit() {
	log.Println("Setting up local DynamoDB environment")

	localConfig := aws.Config{
		Endpoint: aws.String("http://dynamodb:8000"),
		Region:   aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
			AccessKeyID:     "AccessKeyID",
			SecretAccessKey: "SecretAccessKey",
		}),
	}

	awsSession, sessionErr = session.NewSession(&localConfig)
	dynamoService = dynamodb.New(awsSession)

	var environmentGateway context.EnvironmentGateway
	applicationContext.Resolve(&environmentGateway)
	tableName := environmentGateway.GetEnvironmentValue("TABLE_NAME")

	//retry a few times
	maxRetries := 3
	retry := 0
	for retry = 0; retry < maxRetries; retry++ {
		//check if the table exists
		_, err := dynamoService.DescribeTable(&dynamodb.DescribeTableInput{
			TableName: aws.String(tableName),
		})

		if err == nil {
			//the table exists
			log.Printf("Table %s already exists", tableName)
			break
		} else {
			//table probably doesn't exist, create it

			_, err := dynamoService.CreateTable(&dynamodb.CreateTableInput{
				TableName: aws.String(tableName),
				AttributeDefinitions: []*dynamodb.AttributeDefinition{
					{
						AttributeName: aws.String("primary"),
						AttributeType: aws.String("S"),
					},
					{
						AttributeName: aws.String("secondary"),
						AttributeType: aws.String("S"),
					},
				},
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("primary"),
						KeyType:       aws.String("HASH"),
					},
					{
						AttributeName: aws.String("secondary"),
						KeyType:       aws.String("RANGE"),
					},
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
					ReadCapacityUnits:  aws.Int64(5),
					WriteCapacityUnits: aws.Int64(5),
				},
			})

			if err == nil {
				log.Printf("Table %s created", tableName)
				break
			}

			//wait a bit
			time.Sleep(time.Second)
		}
	}

	if retry == maxRetries {
		log.Fatalf("Exhausted retries to create local DynamoDB table")
	}
	log.Println("Done setting up local DynamoDB environment")
}