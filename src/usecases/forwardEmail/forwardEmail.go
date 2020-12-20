package forwardEmail

import (
	"bytes"
	"fmt"
	"github.com/halprin/email-conceal/src/context"
	"log"
	"net/mail"
	"strings"
)


var applicationContext = context.ApplicationContext{}

type ForwardEmailUsecase interface {
	ForwardEmail(url string) error
}

type ForwardEmailUsecaseImpl struct {}

func (receiver ForwardEmailUsecaseImpl) ForwardEmail(url string) error {
	//TODO: I may be copying `rawEmail` around, which could be 150 MB or whatever size big of an e-mail.  That would be bad.
	//But maybe not?  I believe I may be passing around a "slice", which internally is a pointer?
	log.Println("Reading the e-mail")

	var emailReaderGateway ReadEmailGateway
	applicationContext.Resolve(&emailReaderGateway)

	rawEmail, err := emailReaderGateway.ReadEmail(url)
	if err != nil {
		log.Printf("Reading the e-mail failed, %+v\n", err)
		return NewUnableToReadEmailError(url, err)
	}

	log.Println("Parsing the e-mail")
	email, err := emailFromRawBytes(rawEmail)
	if err != nil {
		log.Printf("Parsing the e-mail failed, %+v\n", err)
		return NewUnableToParseEmailError(err)
	}

	var environmentGateway context.EnvironmentGateway
	applicationContext.Resolve(&environmentGateway)

	domain := environmentGateway.GetEnvironmentValue("DOMAIN")
	concealedRecipients := getConcealedRecipients(email, domain)
	log.Printf("Concealed recipients are %s", concealedRecipients)
	log.Println("Looking up actual recipients...")
	actualRecipients := getActualRecipients(concealedRecipients, domain)
	log.Printf("Actual recipients are %s", actualRecipients)
	if len(actualRecipients) == 0 {
		log.Println("No actual recipients to forward e-mail to")
		return nil
	}

	log.Println("Changing the headers in e-mail")
	changeHeadersInEmail(email)

	log.Println("Reconstruct raw e-mail bytes")
	myTypeEmail := ByteSliceMessage(*email)
	modifiedRawEmail := myTypeEmail.ByteSlice()

	log.Println("Sending the e-mail")
	err = applicationContext.SendEmailGateway(modifiedRawEmail, actualRecipients)
	if err != nil {
		log.Printf("Sending the e-mail failed, %+v\n", err)
		return NewUnableToSendEmailError(err)
	}

	return nil
}

func getActualRecipients(concealedRecipients []string, domain string) []string {
	recipientsStrings := make([]string, 0, len(concealedRecipients))

	for _, concealedRecipient := range concealedRecipients {
		concealedRecipientPrefix := strings.TrimSuffix(concealedRecipient, fmt.Sprintf("@%s", domain))

		actualRecipient, err := applicationContext.GetRealEmailForConcealPrefix(concealedRecipientPrefix)

		if err != nil {
			log.Printf("Unable to get actual recipient for concealed recipient %s due to error %+v", concealedRecipient, err)
			log.Println("Ignoring recipient")
			continue
		}

		recipientsStrings = append(recipientsStrings, actualRecipient)
	}

	return recipientsStrings
}

func getConcealedRecipients(email *mail.Message, domain string) []string {
	recipientsAddresses, _ := email.Header.AddressList("To")

	recipientsStrings := make([]string, 0, len(recipientsAddresses))

	for _, recipientAddress := range recipientsAddresses {
		if strings.HasSuffix(recipientAddress.Address, domain) {
			//it's our domain so we need to forward
			recipientsStrings = append(recipientsStrings, recipientAddress.Address)
		}
	}

	return recipientsStrings
}

func emailFromRawBytes(rawEmail []byte) (*mail.Message, error) {
	return mail.ReadMessage(bytes.NewReader(rawEmail))
}

func changeHeadersInEmail(email *mail.Message) {
	delete(email.Header, "Dkim-Signature")  //the signature is handled by the forwarding service, not us
	delete(email.Header, "Return-Path")  //don't continue on the return path, especially because it's probably not from a verified domain

	//construct the complete forwarder e-mail address
	var environmentGateway context.EnvironmentGateway
	applicationContext.Resolve(&environmentGateway)

	forwarderEmailPrefix := environmentGateway.GetEnvironmentValue("FORWARDER_EMAIL_PREFIX")
	domain := environmentGateway.GetEnvironmentValue("DOMAIN")
	forwarderEmailAddress := fmt.Sprintf("%s@%s", forwarderEmailPrefix, domain)

	//get the "From" based header
	originalFrom := fromAddressOf(email)
	if originalFrom == nil {
		log.Println("E-mail doesn't have any from-based headers")
		originalFrom = &mail.Address{
			Name:    "Unknown Sender",
			Address: forwarderEmailAddress,
		}
	}

	delete(email.Header, "From")
	delete(email.Header, "Sender")
	delete(email.Header, "Source")

	//change the From to the service, and the Reply-To to the original sender
	originalFromString := originalFrom.String()
	newFrom := mail.Address{
		Name:    originalFromString,
		Address: forwarderEmailAddress,
	}
	email.Header["From"] = []string{newFrom.String()}
	email.Header["Reply-To"] = []string{originalFromString}
}

func fromAddressOf(email *mail.Message) *mail.Address {
	originalFrom, _ := email.Header.AddressList("From")

	if originalFrom == nil {
		originalFrom, _ = email.Header.AddressList("Sender")
	}

	if originalFrom == nil {
		originalFrom, _ = email.Header.AddressList("Source")
	}

	if originalFrom == nil {
		return nil
	}

	return originalFrom[0]
}

type ByteSliceMessage mail.Message

func (email *ByteSliceMessage) ByteSlice() []byte {
	var rawEmailBuffer bytes.Buffer

	//write out the headers
	for headerKey, headerValueList := range email.Header {
		for _, header := range headerValueList {
			var fullHeader strings.Builder
			fullHeader.WriteString(headerKey)
			fullHeader.WriteString(": ")
			fullHeader.WriteString(header)
			fullHeader.WriteString("\r\n")

			rawEmailBuffer.WriteString(fullHeader.String())
		}
	}

	//write out the extra \r\n that designates the the beginning of the body
	rawEmailBuffer.WriteString("\r\n")

	rawEmailBuffer.ReadFrom(email.Body)

	return rawEmailBuffer.Bytes()
}
