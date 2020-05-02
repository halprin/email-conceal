package gateways

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/halprin/email-conceal/forwarder/context"
)

var sesService = ses.New(awsSession)

func AwsSesSendEmailGateway(email []byte, recipient string, applicationContext context.ApplicationContext) error {
	if sessionErr != nil {
		return sessionErr
	}

	sendRawEmailInput := &ses.SendRawEmailInput{
		Source:       aws.String(applicationContext.EnvironmentGateway("FORWARDER_EMAIL")),
		Destinations: []*string{aws.String(recipient)},
		RawMessage:   &ses.RawMessage{
			Data: email,
		},
	}

	_, err := sesService.SendRawEmail(sendRawEmailInput)

	return err
}
