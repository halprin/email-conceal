package gateways

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/halprin/email-conceal/forwarder/context"
)

func AwsSesSendEmailGateway(email []byte, applicationContext context.ApplicationContext) error {
	awsSession, err := session.NewSession()
	if err != nil {
		return err
	}
	sesService := ses.New(awsSession)

	sendRawEmailInput := &ses.SendRawEmailInput{
		Source:       aws.String(applicationContext.EnvironmentGateway("FORWARDER_EMAIL")),
		Destinations: []*string{aws.String(applicationContext.EnvironmentGateway("RECEIVING_EMAIL"))},
		RawMessage:   &ses.RawMessage{
			Data: email,
		},
	}

	_, err = sesService.SendRawEmail(sendRawEmailInput)

	return err
}
