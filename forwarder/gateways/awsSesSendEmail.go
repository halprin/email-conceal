package gateways

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/halprin/email-conceal/forwarder/context"
	"github.com/halprin/email-conceal/forwarder/external/lib/errors"
)

func AwsSesSendEmailGateway(email []byte, applicationContext context.ApplicationContext) error {
	config := &aws.Config{
		Region: aws.String("us-east-1"),
	}
	session, _ := session.NewSession(config)
	sesService := ses.New(session)

	sendRawEmailInput := &ses.SendRawEmailInput{
		Source:       aws.String(applicationContext.EnvironmentGateway("FORWARDER_EMAIL")),
		Destinations: []*string{aws.String(applicationContext.EnvironmentGateway("RECEIVING_EMAIL"))},
		RawMessage:   &ses.RawMessage{
			Data: email,
		},
	}

	_, err := sesService.SendRawEmail(sendRawEmailInput)

	return err
}
