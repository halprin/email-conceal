package gateways

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/halprin/email-conceal/forwarder/context"
	"log"
)

var sesService = ses.New(awsSession)

func AwsSesSendEmailGateway(email []byte, recipients []string, applicationContext context.ApplicationContext) error {
	if sessionErr != nil {
		return sessionErr
	}

	forwarderEmailPrefix := applicationContext.EnvironmentGateway("FORWARDER_EMAIL_PREFIX")
	domain := applicationContext.EnvironmentGateway("DOMAIN")

	recipientsPointers := make([]*string, 0, len(recipients))
	for _, recipient := range recipients {
		recipientsPointers = append(recipientsPointers, aws.String(recipient))
	}

	log.Printf("Fowarding email from e-mail %s@%s", forwarderEmailPrefix, domain)
	sendRawEmailInput := &ses.SendRawEmailInput{
		Source:       aws.String(fmt.Sprintf("%s@%s", forwarderEmailPrefix, domain)),
		Destinations: recipientsPointers,
		RawMessage: &ses.RawMessage{
			Data: email,
		},
	}

	_, err := sesService.SendRawEmail(sendRawEmailInput)

	return err
}
