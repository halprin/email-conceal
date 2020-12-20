package awsSesSendEmail

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/halprin/email-conceal/src/context"
	"log"
)

var sesService = ses.New(awsSession)

var applicationContext = context.ApplicationContext{}

type AwsSesSendEmailGateway struct {}

func (receiver AwsSesSendEmailGateway) SendEmail(email []byte, recipients []string) error {
	if sessionErr != nil {
		return sessionErr
	}

	recipientsPointers := make([]*string, 0, len(recipients))
	for _, recipient := range recipients {
		recipientsPointers = append(recipientsPointers, aws.String(recipient))
	}

	var environmentGateway context.EnvironmentGateway
	applicationContext.Resolve(&environmentGateway)

	forwarderEmailPrefix := environmentGateway.GetEnvironmentValue("FORWARDER_EMAIL_PREFIX")
	domain := environmentGateway.GetEnvironmentValue("DOMAIN")
	forwarderEmailAddress := fmt.Sprintf("%s@%s", forwarderEmailPrefix, domain)

	log.Printf("Fowarding mail from %s", forwarderEmailAddress)
	sendRawEmailInput := &ses.SendRawEmailInput{
		Source:       aws.String(forwarderEmailAddress),
		Destinations: recipientsPointers,
		RawMessage: &ses.RawMessage{
			Data: email,
		},
	}

	_, err := sesService.SendRawEmail(sendRawEmailInput)

	return err
}
