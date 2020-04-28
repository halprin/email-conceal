package controllers

import (
	"encoding/json"
	"github.com/halprin/email-conceal/forwarder/context"
	"log"
	"strings"
)

type S3FileUploadedEvent struct {
	Records []struct {
		S3 struct{
			Bucket struct{
				Name string
			}
			Object struct{
				Key string
			}
		}
	}
}

func SqsQueueForwardEmail(arguments map[string]interface{}, applicationContext context.ApplicationContext) error {
	messageJsonString := arguments["message"].(*string)

	var message S3FileUploadedEvent
	err := json.Unmarshal([]byte(*messageJsonString), &message)
	if err != nil {
		log.Printf("Unable to unmarshal the JSON message; %+v", err)
		return err
	}

	messageRecords := message.Records
	for _, record := range messageRecords {
		s3 := record.S3
		bucket := s3.Bucket.Name
		object := s3.Object.Key

		url := constructS3Url(bucket, object)

		log.Println("URL to read e-mail from =", url)

		err = applicationContext.ForwardEmailUsecase(url)
		if err != nil {
			log.Printf("Unable to forward e-mail at %s due to error, %+v\n", url, err)
			return err
		}

		log.Printf("Forwarded e-mail at %s", url)
	}

	return nil
}

func constructS3Url(bucket string, object string) string {
	urlBuilder := strings.Builder{}

	urlBuilder.WriteString("s3://")
	urlBuilder.WriteString(bucket)
	urlBuilder.WriteString("/")
	urlBuilder.WriteString(object)

	return urlBuilder.String()
}
