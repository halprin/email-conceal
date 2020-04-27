package controllers

import (
	"encoding/json"
	"github.com/halprin/email-conceal/forwarder/context"
	"log"
	"strings"
)

func SqsQueueForwardEmail(arguments map[string]interface{}, applicationContext context.ApplicationContext) error {
	messageJsonString := arguments["message"].(*string)
	var message map[string]interface{}

	err := json.Unmarshal([]byte(*messageJsonString), &message)
	if err != nil {
		log.Printf("Unable to unmarshal the JSON message; %+v", err)
		return err
	}

	messageRecords := message["Records"].([]interface{})
	for _, record := range messageRecords {
		s3 := record.(map[string]interface{})["s3"].(map[string]interface{})
		bucket := s3["bucket"].(map[string]interface{})["name"].(string)
		object := s3["object"].(map[string]interface{})["key"].(string)

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
