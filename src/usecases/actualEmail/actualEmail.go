package actualEmail

import (
	"crypto/rand"
	"fmt"
	"github.com/halprin/email-conceal/src/context"
	"github.com/halprin/email-conceal/src/entities"
	"github.com/halprin/email-conceal/src/external/lib/errors"
	"github.com/halprin/email-conceal/src/usecases/forwardEmail"
	"github.com/jordan-wright/email"
	"math/big"
	"strings"
)

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
const secretKeyLength = 128

var applicationContext = context.ApplicationContext{}
var actualEmailConfigGateway ActualEmailConfigurationGateway
var sendEmailGateway forwardEmail.SendEmailGateway
var environmentGateway context.EnvironmentGateway

type ActualEmailUsecase interface {
	Add(actualEmail string) error
}

type ActualEmailUsecaseImpl struct{}

func (receiver ActualEmailUsecaseImpl) Init() {
	applicationContext.Resolve(&actualEmailConfigGateway)
	applicationContext.Resolve(&sendEmailGateway)
	applicationContext.Resolve(&environmentGateway)
}

func (receiver ActualEmailUsecaseImpl) Add(actualEmail string) error {
	err := entities.ValidateEmail(actualEmail)
	if err != nil {
		return err
	}

	secret, err := generateSecret()
	if err != nil {
		return errors.Wrap(err, "Unable to generate a new secret")
	}

	err = actualEmailConfigGateway.AddUnprovedActualEmail(actualEmail, secret)
	if err != nil {
		return errors.Wrap(err, "Failed to write the secret key associated with the actual e-mail address")
	}

	registrationEmailBytes, err := generateRegistrationEmail(actualEmail, secret)
	if err != nil {
		return errors.Wrap(err, "Failed to generate the registration e-mail")
	}

	err = sendEmailGateway.SendEmail(registrationEmailBytes, []string{actualEmail})
	if err != nil {
		return errors.Wrap(err, "Failed to send the confirmation e-mail")
	}

	return nil
}

func generateSecret() (string, error) {
	secretBuilder := strings.Builder{}

	for secretIndex := 0; secretIndex < secretKeyLength; secretIndex++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
		if err != nil {
			return "", err
		}

		runeVersionOfAlphabet := []rune(alphabet)

		secretBuilder.WriteRune(runeVersionOfAlphabet[randomIndex.Int64()])
	}

	return secretBuilder.String(), nil
}

func generateRegistrationEmail(receipient string, secret string) ([]byte, error) {
	domain := environmentGateway.GetEnvironmentValue("DOMAIN")
	forwarderEmailPrefix := environmentGateway.GetEnvironmentValue("FORWARDER_EMAIL_PREFIX")

	body := fmt.Sprintf("You have registered your e-mail, %s, for E-mail Conceal.  To complete registration, click the link.  Do nothing if you didn't register.\n"+
		"https://%s/v1/activateRegistration/%s", receipient, domain, secret)

	registrationEmail := email.NewEmail()
	registrationEmail.From = fmt.Sprintf("%s@%s", forwarderEmailPrefix, domain)
	registrationEmail.To = []string{receipient}
	registrationEmail.Subject = "E-mail Conceal Registration"
	registrationEmail.Text = []byte(body)

	registrationEmailBytes, err := registrationEmail.Bytes()
	if err != nil {
		return nil, err
	}

	return registrationEmailBytes, nil
}
