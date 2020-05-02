package usecases

import (
	"bytes"
	"fmt"
	"github.com/halprin/email-conceal/forwarder/context"
	"log"
	"net/mail"
	"strings"
)

func ForwardEmailUsecase(url string, applicationContext context.ApplicationContext) error {
	//TODO: I may be copying `rawEmail` around, which could be 150 MB or whatever size big of an e-mail.  That would be bad.
	//But maybe not?  I believe I may be passing around a "slice", which internally is a pointer?
	log.Println("Reading the e-mail")
	rawEmail, err := applicationContext.ReadEmailGateway(url)
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

	log.Println("Changing the headers in e-mail")
	changeHeadersInEmail(email, applicationContext)

	log.Println("Reconstruct raw e-mail bytes")
	myTypeEmail := ByteSliceMessage(*email)
	modifiedRawEmail := myTypeEmail.ByteSlice()

	log.Println("Sending the e-mail")
	recipient := applicationContext.EnvironmentGateway("RECEIVING_EMAIL")
	err = applicationContext.SendEmailGateway(modifiedRawEmail, recipient)
	if err != nil {
		log.Printf("Sending the e-mail failed, %+v\n", err)
		return NewUnableToSendEmailError(err)
	}

	return nil
}

func emailFromRawBytes(rawEmail []byte) (*mail.Message, error) {
	return mail.ReadMessage(bytes.NewReader(rawEmail))
}

func changeHeadersInEmail(email *mail.Message, applicationContext context.ApplicationContext) {
	delete(email.Header, "Dkim-Signature")  //the signature is handled by the forwarding service, not us
	delete(email.Header, "Return-Path")  //don't continue on the return path, especially because it's probably not from a verified domain

	//get the "From" based header
	originalFrom := fromAddressOf(email)
	if originalFrom == nil {
		fmt.Println("E-mail doesn't have any from-based headers")
		originalFrom = &mail.Address{
			Name:    "Unknown Sender",
			Address: applicationContext.EnvironmentGateway("FORWARDER_EMAIL"),
		}
	}

	delete(email.Header, "From")
	delete(email.Header, "Sender")
	delete(email.Header, "Source")

	//change the From to the service, and the Reply-To to the original sender
	originalFromString := originalFrom.String()
	newFrom := mail.Address{
		Name:    originalFromString,
		Address: applicationContext.EnvironmentGateway("FORWARDER_EMAIL"),
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
