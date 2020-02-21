package usecases

import (
	"bytes"
	"fmt"
	"github.com/halprin/email-conceal/forwarder/context"
	"github.com/halprin/email-conceal/forwarder/external/lib/errors"
	"testing"
)

func TestForwardEmailUsecaseWithFailingToReadEmail(t *testing.T) {
	errorFromGateway := errors.New("something bad happened")
	appContext := context.TestApplicationContext{
		ReturnFromReadEmailGateway:      nil,
		ReturnErrorFromReadEmailGateway: errorFromGateway,
	}

	testUrl := "https://email.com"
	err := ForwardEmailUsecase(testUrl, &appContext)

	if !errors.Is(err, NewUnableToReadEmailError(testUrl, errorFromGateway)) {
		t.Errorf("An UnableToReadEmailError should have been returned from ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}
}

func TestForwardEmailUsecaseWithFailingToParseEmail(t *testing.T) {
	badEmail := ` To: jobs@apple.com
From: moof@apple.com
Subject: bad T header

There is an initial space and that is bad.
`
	appContext := context.TestApplicationContext{
		ReturnFromReadEmailGateway: []byte(badEmail),
	}

	err := ForwardEmailUsecase("https://email.com", &appContext)

	if !errors.Is(err, NewUnableToParseEmailError(nil)) {
		t.Errorf("An UnableToParseEmailError should have been returned from ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}
}

func TestForwardEmailUsecaseWillRemoveCertainHeaders(t *testing.T) {
	dkimHeader := "Dkim-Signature"
	returnPathHeader := "Return-Path"

	email := `To: jobs@apple.com
From: moof@apple.com
Dkim-Signature: asdf
Return-Path: asdf
Subject: lol

Test e-mail.
`
	appContext := context.TestApplicationContext{
		ReturnFromReadEmailGateway: []byte(email),
	}

	err := ForwardEmailUsecase("https://email.com", &appContext)

	if err != nil {
		t.Errorf("No error should have been returned from the ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}

	rawForwardedEmail := appContext.ReceivedSendEmailGatewayArguments
	if bytes.Contains(rawForwardedEmail, []byte(dkimHeader)) {
		t.Errorf("Header %s was not removed from the e-mail; it should have been", dkimHeader)
	}

	if bytes.Contains(rawForwardedEmail, []byte(returnPathHeader)) {
		t.Errorf("Header %s was not removed from the e-mail; it should have been", returnPathHeader)
	}
}

func TestForwardEmailUsecaseGrabsFromSender(t *testing.T) {
	fromHeader := "Sender"
	fromName := "DogCow"
	fromAddress := "moof@apple.com"

	email := fmt.Sprintf(`To: jobs@apple.com
%s: %s <%s>
Subject: lol

Test e-mail.
`, fromHeader, fromName, fromAddress)

	appContext := context.TestApplicationContext{
		ReturnFromReadEmailGateway: []byte(email),
	}

	err := ForwardEmailUsecase("https://email.com", &appContext)

	if err != nil {
		t.Errorf("No error should have been returned from the ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}

	rawForwardedEmail := appContext.ReceivedSendEmailGatewayArguments
	if bytes.Contains(rawForwardedEmail, []byte(fromHeader)) {
		t.Errorf("Header %s was not removed from the e-mail; it should have been", fromHeader)
	}

	if !bytes.Contains(rawForwardedEmail, []byte(fromName)) {
		t.Errorf("The from name %s is missing from the e-mail and it should have been there", fromName)
	}

	if !bytes.Contains(rawForwardedEmail, []byte(fromAddress)) {
		t.Errorf("The from address %s is missing from the e-mail and it should have been there", fromAddress)
	}
}

func TestForwardEmailUsecaseGrabsFromSource(t *testing.T) {
	fromHeader := "Source"
	fromName := "DogCow"
	fromAddress := "moof@apple.com"

	email := fmt.Sprintf(`To: jobs@apple.com
%s: %s <%s>
Subject: lol

Test e-mail.
`, fromHeader, fromName, fromAddress)

	appContext := context.TestApplicationContext{
		ReturnFromReadEmailGateway: []byte(email),
	}

	err := ForwardEmailUsecase("https://email.com", &appContext)

	if err != nil {
		t.Errorf("No error should have been returned from the ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}

	rawForwardedEmail := appContext.ReceivedSendEmailGatewayArguments
	if bytes.Contains(rawForwardedEmail, []byte(fromHeader)) {
		t.Errorf("Header %s was not removed from the e-mail; it should have been", fromHeader)
	}

	if !bytes.Contains(rawForwardedEmail, []byte(fromName)) {
		t.Errorf("The from name %s is missing from the e-mail and it should have been there", fromName)
	}

	if !bytes.Contains(rawForwardedEmail, []byte(fromAddress)) {
		t.Errorf("The from address %s is missing from the e-mail and it should have been there", fromAddress)
	}
}

func TestForwardEmailUsecaseGrabsFromFrom(t *testing.T) {
	fromHeader := "From"
	fromName := "DogCow"
	fromAddress := "moof@apple.com"

	email := fmt.Sprintf(`To: jobs@apple.com
%s: %s <%s>
Subject: lol

Test e-mail.
`, fromHeader, fromName, fromAddress)

	appContext := context.TestApplicationContext{
		ReturnFromReadEmailGateway: []byte(email),
	}

	err := ForwardEmailUsecase("https://email.com", &appContext)

	if err != nil {
		t.Errorf("No error should have been returned from the ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}

	rawForwardedEmail := appContext.ReceivedSendEmailGatewayArguments

	if !bytes.Contains(rawForwardedEmail, []byte(fromName)) {
		t.Errorf("The from name %s is missing from the e-mail and it should have been there", fromName)
	}

	if !bytes.Contains(rawForwardedEmail, []byte(fromAddress)) {
		t.Errorf("The from address %s is missing from the e-mail and it should have been there", fromAddress)
	}
}

func TestForwardEmailUsecaseUsesFromOverSender(t *testing.T) {
	fromEmail := "moof@apple.com"
	senderEmail := "whatever@apple.com"

	email := fmt.Sprintf(`To: jobs@apple.com
Sender: %s
From: %s
Subject: lol

Test e-mail.
`, senderEmail, fromEmail)

	appContext := context.TestApplicationContext{
		ReturnFromReadEmailGateway: []byte(email),
	}

	err := ForwardEmailUsecase("https://email.com", &appContext)

	if err != nil {
		t.Errorf("No error should have been returned from the ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}

	rawForwardedEmail := appContext.ReceivedSendEmailGatewayArguments

	if !bytes.Contains(rawForwardedEmail, []byte(fromEmail)) {
		t.Errorf("The from address %s is missing from the e-mail and it should have been there", fromEmail)
	}

	if bytes.Contains(rawForwardedEmail, []byte(senderEmail)) {
		t.Errorf("The sender address %s is in the e-mail and it shouldn't have been there", senderEmail)
	}
}

func TestForwardEmailUsecaseUsesSenderOverSource(t *testing.T) {
	senderEmail := "moof@apple.com"
	sourceEmail := "whatever@apple.com"

	email := fmt.Sprintf(`To: jobs@apple.com
Source: %s
Sender: %s
Subject: lol

Test e-mail.
`, sourceEmail, senderEmail)

	appContext := context.TestApplicationContext{
		ReturnFromReadEmailGateway: []byte(email),
	}

	err := ForwardEmailUsecase("https://email.com", &appContext)

	if err != nil {
		t.Errorf("No error should have been returned from the ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}

	rawForwardedEmail := appContext.ReceivedSendEmailGatewayArguments

	if !bytes.Contains(rawForwardedEmail, []byte(senderEmail)) {
		t.Errorf("The sender address %s is missing from the e-mail and it should have been there", senderEmail)
	}

	if bytes.Contains(rawForwardedEmail, []byte(sourceEmail)) {
		t.Errorf("The source address %s is in the e-mail and it shouldn't have been there", sourceEmail)
	}
}

func TestForwardEmailUsecaseUsesSourceAddress(t *testing.T) {
	sourceEmail := "moof@apple.com"

	email := fmt.Sprintf(`To: jobs@apple.com
Source: %s
Subject: lol

Test e-mail.
`, sourceEmail)

	appContext := context.TestApplicationContext{
		ReturnFromReadEmailGateway: []byte(email),
	}

	err := ForwardEmailUsecase("https://email.com", &appContext)

	if err != nil {
		t.Errorf("No error should have been returned from the ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}

	rawForwardedEmail := appContext.ReceivedSendEmailGatewayArguments

	if !bytes.Contains(rawForwardedEmail, []byte(sourceEmail)) {
		t.Errorf("The source address %s is missing from the e-mail and it should have been there", sourceEmail)
	}
}

func TestForwardEmailUsecaseWithNoAddress(t *testing.T) {
	forwarderEmail := "moof@apple.com"

	email := `To: jobs@apple.com
Subject: lol

Test e-mail.
`

	appContext := context.TestApplicationContext{
		ReturnFromReadEmailGateway: []byte(email),
		ReturnFromEnvironmentGateway: forwarderEmail,
	}

	err := ForwardEmailUsecase("https://email.com", &appContext)

	if err != nil {
		t.Errorf("No error should have been returned from the ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}

	rawForwardedEmail := appContext.ReceivedSendEmailGatewayArguments

	if !bytes.Contains(rawForwardedEmail, []byte(forwarderEmail)) {
		t.Errorf("The forwarder address %s is missing from the e-mail and it should have been there", forwarderEmail)
	}
	if !bytes.Contains(rawForwardedEmail, []byte("Unknown Sender")) {
		t.Errorf("The forwarder address name %s is missing from the e-mail and it should have been there", "Unknown Sender")
	}
}
