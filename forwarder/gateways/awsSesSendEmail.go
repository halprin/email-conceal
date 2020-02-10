package gateways

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/halprin/email-conceal/forwarder/context"
)

func AwsSesSendEmailGateway(email []byte, applicationContext context.ApplicationContext) error {
	config := &aws.Config{
		Region: aws.String("us-east-1"),
	}
	session, _ := session.NewSession(config)
	sesService := ses.New(session)

	sendRawEmailInput := &ses.SendRawEmailInput{
		Source:       aws.String("___@_____.___"),
		Destinations: []*string{aws.String("___@_____.___")},
		RawMessage:   &ses.RawMessage{
			Data: email,
		},
	}

	result, err := sesService.SendRawEmail(sendRawEmailInput)

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(result)

	return nil
}
