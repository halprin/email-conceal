package gateways

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/halprin/email-conceal/forwarder/context"
)

var awsSession, sessionErr = session.NewSession()
var sesService = ses.New(awsSession)

func AwsSesSendEmailGateway(email []byte, applicationContext context.ApplicationContext) error {
	if sessionErr != nil {
		return sessionErr
	}

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
