package forwardEmail

import (
	"bytes"
	"fmt"
	"github.com/halprin/email-conceal/src/context"
	"github.com/halprin/email-conceal/src/external/lib/errors"
	"testing"
)

var usecase = ForwardEmailUsecaseImpl{}
var testAppContext = context.ApplicationContext{}

type TestReadEmailGateway struct {
	ReadEmailUri         string
	ReadEmailReturnByte  []byte
	ReadEmailReturnError error
}

func (testGateway *TestReadEmailGateway) ReadEmail(uri string) ([]byte, error) {
	testGateway.ReadEmailUri = uri

	return testGateway.ReadEmailReturnByte, testGateway.ReadEmailReturnError
}

type TestSendEmailGateway struct {
	SendEmailEmail       []byte
	SendEmailRecipients  []string
	SendEmailReturnError error
}

func (testGateway *TestSendEmailGateway) SendEmail(email []byte, recipients []string) error {
	testGateway.SendEmailEmail = email
	testGateway.SendEmailRecipients = recipients

	return testGateway.SendEmailReturnError
}

type TestConfigurationGateway struct {
	GetRealEmailConcealPrefix     string
	GetRealEmailReturn map[string][]*string
	GetRealEmailReturnError       error
}

func (testGateway *TestConfigurationGateway) GetRealEmailAddressForConcealPrefix(concealedRecipientPrefix string) (string, *string, error) {
	testGateway.GetRealEmailConcealPrefix = concealedRecipientPrefix

	returnEmail := ""
	var returnDescription *string
	returnArray := testGateway.GetRealEmailReturn[concealedRecipientPrefix]
	if len(returnArray) == 2 {
		returnEmail = *returnArray[0]
		returnDescription = returnArray[1]
	}

	return returnEmail, returnDescription, testGateway.GetRealEmailReturnError
}

type TestEnvironmentGateway struct {
	GetEnvironmentValueReturn map[string]string
}

func (testGateway *TestEnvironmentGateway) GetEnvironmentValue(key string) string {
	return testGateway.GetEnvironmentValueReturn[key]
}

func resetBaseDependencies() {
	testAppContext.Reset()

	testAppContext.Bind(func() ReadEmailGateway {
		return &TestReadEmailGateway{}
	})

	testAppContext.Bind(func() SendEmailGateway {
		return &TestSendEmailGateway{}
	})

	testAppContext.Bind(func() context.EnvironmentGateway {
		return &TestEnvironmentGateway{}
	})

	testAppContext.Bind(func() ConfigurationGateway {
		return &TestConfigurationGateway{}
	})
}

func TestForwardEmailUsecaseWithFailingToReadEmail(t *testing.T) {
	resetBaseDependencies()

	errorFromGateway := errors.New("something bad happened")
	testReadEmailGateway := TestReadEmailGateway{
		ReadEmailReturnError: errorFromGateway,
	}
	testAppContext.Bind(func() ReadEmailGateway {
		return &testReadEmailGateway
	})

	testUrl := "https://email.com"
	err := usecase.ForwardEmail(testUrl)

	if !errors.Is(err, NewUnableToReadEmailError(testUrl, errorFromGateway)) {
		t.Errorf("An UnableToReadEmailError should have been returned from ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}
}

func TestForwardEmailUsecaseWithFailingToParseEmail(t *testing.T) {
	resetBaseDependencies()

	badEmail := ` To: jobs@apple.com
From: moof@apple.com
Subject: bad T header

There is an initial space and that is bad.
`
	testReadEmailGateway := TestReadEmailGateway{
		ReadEmailReturnByte: []byte(badEmail),
	}
	testAppContext.Bind(func() ReadEmailGateway {
		return &testReadEmailGateway
	})

	err := usecase.ForwardEmail("https://email.com")

	if !errors.Is(err, NewUnableToParseEmailError(nil)) {
		t.Errorf("An UnableToParseEmailError should have been returned from ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}
}

func TestForwardEmailUsecaseWillRemoveCertainHeaders(t *testing.T) {
	resetBaseDependencies()

	dkimHeader := "Dkim-Signature"
	returnPathHeader := "Return-Path"

	email := `To: jobs@apple.com
From: moof@apple.com
Dkim-Signature: asdf
Return-Path: asdf
Subject: lol

Test e-mail.
`
	testReadEmailGateway := TestReadEmailGateway{
		ReadEmailReturnByte: []byte(email),
	}
	testAppContext.Bind(func() ReadEmailGateway {
		return &testReadEmailGateway
	})

	testSendEmailGateway := TestSendEmailGateway{}
	testAppContext.Bind(func() SendEmailGateway {
		return &testSendEmailGateway
	})

	err := usecase.ForwardEmail("https://email.com")

	if err != nil {
		t.Errorf("No error should have been returned from the ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}

	rawForwardedEmail := testSendEmailGateway.SendEmailEmail
	if bytes.Contains(rawForwardedEmail, []byte(dkimHeader)) {
		t.Errorf("Header %s was not removed from the e-mail; it should have been", dkimHeader)
	}

	if bytes.Contains(rawForwardedEmail, []byte(returnPathHeader)) {
		t.Errorf("Header %s was not removed from the e-mail; it should have been", returnPathHeader)
	}
}

func TestForwardEmailUsecaseGrabsFromSender(t *testing.T) {
	resetBaseDependencies()

	fromHeader := "Sender"
	fromName := "DogCow"
	fromAddress := "moof@apple.com"

	email := fmt.Sprintf(`To: jobs@apple.com
%s: %s <%s>
Subject: lol

Test e-mail.
`, fromHeader, fromName, fromAddress)

	testReadEmailGateway := TestReadEmailGateway{
		ReadEmailReturnByte: []byte(email),
	}
	testAppContext.Bind(func() ReadEmailGateway {
		return &testReadEmailGateway
	})

	testSendEmailGateway := TestSendEmailGateway{}
	testAppContext.Bind(func() SendEmailGateway {
		return &testSendEmailGateway
	})

	err := usecase.ForwardEmail("https://email.com")

	if err != nil {
		t.Errorf("No error should have been returned from the ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}

	rawForwardedEmail := testSendEmailGateway.SendEmailEmail
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
	resetBaseDependencies()

	fromHeader := "Source"
	fromName := "DogCow"
	fromAddress := "moof@apple.com"

	email := fmt.Sprintf(`To: jobs@apple.com
%s: %s <%s>
Subject: lol

Test e-mail.
`, fromHeader, fromName, fromAddress)

	testReadEmailGateway := TestReadEmailGateway{
		ReadEmailReturnByte: []byte(email),
	}
	testAppContext.Bind(func() ReadEmailGateway {
		return &testReadEmailGateway
	})

	testSendEmailGateway := TestSendEmailGateway{}
	testAppContext.Bind(func() SendEmailGateway {
		return &testSendEmailGateway
	})

	err := usecase.ForwardEmail("https://email.com")

	if err != nil {
		t.Errorf("No error should have been returned from the ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}

	rawForwardedEmail := testSendEmailGateway.SendEmailEmail
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
	resetBaseDependencies()

	fromHeader := "From"
	fromName := "DogCow"
	fromAddress := "moof@apple.com"

	email := fmt.Sprintf(`To: jobs@apple.com
%s: %s <%s>
Subject: lol

Test e-mail.
`, fromHeader, fromName, fromAddress)

	testReadEmailGateway := TestReadEmailGateway{
		ReadEmailReturnByte: []byte(email),
	}
	testAppContext.Bind(func() ReadEmailGateway {
		return &testReadEmailGateway
	})

	testSendEmailGateway := TestSendEmailGateway{}
	testAppContext.Bind(func() SendEmailGateway {
		return &testSendEmailGateway
	})

	err := usecase.ForwardEmail("https://email.com")

	if err != nil {
		t.Errorf("No error should have been returned from the ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}

	rawForwardedEmail := testSendEmailGateway.SendEmailEmail

	if !bytes.Contains(rawForwardedEmail, []byte(fromName)) {
		t.Errorf("The from name %s is missing from the e-mail and it should have been there", fromName)
	}

	if !bytes.Contains(rawForwardedEmail, []byte(fromAddress)) {
		t.Errorf("The from address %s is missing from the e-mail and it should have been there", fromAddress)
	}
}

func TestForwardEmailUsecaseUsesFromOverSender(t *testing.T) {
	resetBaseDependencies()

	fromEmail := "moof@apple.com"
	senderEmail := "whatever@apple.com"

	email := fmt.Sprintf(`To: jobs@apple.com
Sender: %s
From: %s
Subject: lol

Test e-mail.
`, senderEmail, fromEmail)

	testReadEmailGateway := TestReadEmailGateway{
		ReadEmailReturnByte: []byte(email),
	}
	testAppContext.Bind(func() ReadEmailGateway {
		return &testReadEmailGateway
	})

	testSendEmailGateway := TestSendEmailGateway{}
	testAppContext.Bind(func() SendEmailGateway {
		return &testSendEmailGateway
	})

	err := usecase.ForwardEmail("https://email.com")

	if err != nil {
		t.Errorf("No error should have been returned from the ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}

	rawForwardedEmail := testSendEmailGateway.SendEmailEmail

	if !bytes.Contains(rawForwardedEmail, []byte(fromEmail)) {
		t.Errorf("The from address %s is missing from the e-mail and it should have been there", fromEmail)
	}

	if bytes.Contains(rawForwardedEmail, []byte(senderEmail)) {
		t.Errorf("The sender address %s is in the e-mail and it shouldn't have been there", senderEmail)
	}
}

func TestForwardEmailUsecaseUsesSenderOverSource(t *testing.T) {
	resetBaseDependencies()

	senderEmail := "moof@apple.com"
	sourceEmail := "whatever@apple.com"

	email := fmt.Sprintf(`To: jobs@apple.com
Source: %s
Sender: %s
Subject: lol

Test e-mail.
`, sourceEmail, senderEmail)

	testReadEmailGateway := TestReadEmailGateway{
		ReadEmailReturnByte: []byte(email),
	}
	testAppContext.Bind(func() ReadEmailGateway {
		return &testReadEmailGateway
	})

	testSendEmailGateway := TestSendEmailGateway{}
	testAppContext.Bind(func() SendEmailGateway {
		return &testSendEmailGateway
	})

	err := usecase.ForwardEmail("https://email.com")

	if err != nil {
		t.Errorf("No error should have been returned from the ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}

	rawForwardedEmail := testSendEmailGateway.SendEmailEmail

	if !bytes.Contains(rawForwardedEmail, []byte(senderEmail)) {
		t.Errorf("The sender address %s is missing from the e-mail and it should have been there", senderEmail)
	}

	if bytes.Contains(rawForwardedEmail, []byte(sourceEmail)) {
		t.Errorf("The source address %s is in the e-mail and it shouldn't have been there", sourceEmail)
	}
}

func TestForwardEmailUsecaseUsesSourceAddress(t *testing.T) {
	resetBaseDependencies()

	sourceEmail := "moof@apple.com"

	email := fmt.Sprintf(`To: jobs@apple.com
Source: %s
Subject: lol

Test e-mail.
`, sourceEmail)

	testReadEmailGateway := TestReadEmailGateway{
		ReadEmailReturnByte: []byte(email),
	}
	testAppContext.Bind(func() ReadEmailGateway {
		return &testReadEmailGateway
	})

	testSendEmailGateway := TestSendEmailGateway{}
	testAppContext.Bind(func() SendEmailGateway {
		return &testSendEmailGateway
	})

	err := usecase.ForwardEmail("https://email.com")

	if err != nil {
		t.Errorf("No error should have been returned from the ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}

	rawForwardedEmail := testSendEmailGateway.SendEmailEmail

	if !bytes.Contains(rawForwardedEmail, []byte(sourceEmail)) {
		t.Errorf("The source address %s is missing from the e-mail and it should have been there", sourceEmail)
	}
}

func TestForwardEmailUsecaseWithNoAddress(t *testing.T) {
	resetBaseDependencies()

	forwarderPrefix := "moof"
	domain := "apple.com"
	forwarderEmail := fmt.Sprintf("%s@%s", forwarderPrefix, domain)

	email := `To: jobs@apple.com
Subject: lol

Test e-mail.
`

	testReadEmailGateway := TestReadEmailGateway{
		ReadEmailReturnByte: []byte(email),
	}
	testAppContext.Bind(func() ReadEmailGateway {
		return &testReadEmailGateway
	})

	testSendEmailGateway := TestSendEmailGateway{}
	testAppContext.Bind(func() SendEmailGateway {
		return &testSendEmailGateway
	})

	testEnvironmentGateway := TestEnvironmentGateway{
		GetEnvironmentValueReturn: map[string]string{
			"FORWARDER_EMAIL_PREFIX": forwarderPrefix,
			"DOMAIN": domain,
		},
	}
	testAppContext.Bind(func() context.EnvironmentGateway {
		return &testEnvironmentGateway
	})

	err := usecase.ForwardEmail("https://email.com")

	if err != nil {
		t.Errorf("No error should have been returned from the ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}

	rawForwardedEmail := testSendEmailGateway.SendEmailEmail

	if !bytes.Contains(rawForwardedEmail, []byte(forwarderEmail)) {
		t.Errorf("The forwarder address %s is missing from the e-mail and it should have been there", forwarderEmail)
	}
	if !bytes.Contains(rawForwardedEmail, []byte("Unknown Sender")) {
		t.Errorf("The forwarder address name %s is missing from the e-mail and it should have been there", "Unknown Sender")
	}
}

func TestForwardEmailUsecaseWithFailingToSendEmail(t *testing.T) {
	resetBaseDependencies()

	email := `To: jobs@apple.com
From: moof@apple.com
Subject: lol

Test e-mail
`
	testReadEmailGateway := TestReadEmailGateway{
		ReadEmailReturnByte: []byte(email),
	}
	testAppContext.Bind(func() ReadEmailGateway {
		return &testReadEmailGateway
	})

	testSendEmailGateway := TestSendEmailGateway{
		SendEmailReturnError: errors.New("sending failed"),
	}
	testAppContext.Bind(func() SendEmailGateway {
		return &testSendEmailGateway
	})

	err := usecase.ForwardEmail("https://email.com")

	if !errors.Is(err, NewUnableToSendEmailError(nil)) {
		t.Errorf("An NewUnableToSendEmailError should have been returned from ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}
}

func TestForwardEmailUsecaseEverythingWorks(t *testing.T) {
	resetBaseDependencies()

	body := "This is the coolest e-mail ever"

	email := fmt.Sprintf(`To: jobs@apple.com
From: moof@apple.com
Subject: lol

%s
`, body)

	testReadEmailGateway := TestReadEmailGateway{
		ReadEmailReturnByte: []byte(email),
	}
	testAppContext.Bind(func() ReadEmailGateway {
		return &testReadEmailGateway
	})

	testSendEmailGateway := TestSendEmailGateway{}
	testAppContext.Bind(func() SendEmailGateway {
		return &testSendEmailGateway
	})

	err := usecase.ForwardEmail("https://email.com")

	if err != nil {
		t.Errorf("No error should have been returned from the ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}

	rawForwardedEmail := testSendEmailGateway.SendEmailEmail

	if !bytes.Contains(rawForwardedEmail, []byte(body)) {
		t.Errorf("The e-mail body %s is missing from the e-mail and it should have been there", body)
	}
}

func TestForwardEmailUsecaseThatConvertsKnownConcealAddressesToActualAddresses(t *testing.T) {
	resetBaseDependencies()

	knownConcealedEmail := "known@apple.com"
	knownConcealedEmail2 := "known2@apple.com"
	actualEmail := "moof@dogcow.com"

	email := fmt.Sprintf(`To: %s, %s
From: moof@apple.com
Subject: lol

This is the coolest e-mail ever
`, knownConcealedEmail, knownConcealedEmail2)

	testReadEmailGateway := TestReadEmailGateway{
		ReadEmailReturnByte: []byte(email),
	}
	testAppContext.Bind(func() ReadEmailGateway {
		return &testReadEmailGateway
	})

	testSendEmailGateway := TestSendEmailGateway{}
	testAppContext.Bind(func() SendEmailGateway {
		return &testSendEmailGateway
	})

	testConfigurationGateway := TestConfigurationGateway{
		GetRealEmailReturn: map[string][]*string{
			"known": {&actualEmail, nil},
			"known2": {&actualEmail, nil},
		},
	}
	testAppContext.Bind(func() ConfigurationGateway {
		return &testConfigurationGateway
	})

	testEnvironmentGateway := TestEnvironmentGateway{
		GetEnvironmentValueReturn: map[string]string{
			"DOMAIN": "apple.com",
		},
	}
	testAppContext.Bind(func() context.EnvironmentGateway {
		return &testEnvironmentGateway
	})

	err := usecase.ForwardEmail("https://email.com")

	if err != nil {
		t.Errorf("No error should have been returned from the ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}

	forwardedEmailRecipients := testSendEmailGateway.SendEmailRecipients
	if !contains(forwardedEmailRecipients, actualEmail) {
		t.Errorf("Should have converted the concealed e-mail to the actual e-mails")
	}
}

func TestForwardEmailUsecaseThatDoesNotConvertsUnknownConcealAddresses(t *testing.T) {
	resetBaseDependencies()

	unknownConcealedEmail := "known@apple.com"
	unknownConcealedEmail2 := "known2@apple.com"

	email := fmt.Sprintf(`To: %s, %s
From: moof@apple.com
Subject: lol

This is the coolest e-mail ever
`, unknownConcealedEmail, unknownConcealedEmail2)

	testReadEmailGateway := TestReadEmailGateway{
		ReadEmailReturnByte: []byte(email),
	}
	testAppContext.Bind(func() ReadEmailGateway {
		return &testReadEmailGateway
	})

	testSendEmailGateway := TestSendEmailGateway{}
	testAppContext.Bind(func() SendEmailGateway {
		return &testSendEmailGateway
	})

	testConfigurationGateway := TestConfigurationGateway{
		GetRealEmailReturnError: errors.New("can't find the actual e-mail"),
	}
	testAppContext.Bind(func() ConfigurationGateway {
		return &testConfigurationGateway
	})

	err := usecase.ForwardEmail("https://email.com")

	if err != nil {
		t.Errorf("No error should have been returned from the ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}

	if len(testSendEmailGateway.SendEmailEmail) !=0 && len(testSendEmailGateway.SendEmailRecipients) != 0 {
		t.Errorf("The passed in recipients had some items, instead there shouldn't have been any; passed in recipients = %s", testSendEmailGateway.SendEmailRecipients)
	}
}

func TestToHeaderCorrectlyPutsDescriptionToSpecificEmail(t *testing.T) {
	resetBaseDependencies()

	knownConcealedEmail := "known@apple.com"
	knownConcealedEmail2 := "known2@apple.com"

	actualEmail1 := "moof@dogcow.com"
	actualDescription1 := "The coolest description"

	actualEmail2 := "halprin@dogcow.com"
	actualDescription2 := "Kaboom"

	email := fmt.Sprintf(`To: %s, %s
From: moof@apple.com
Subject: lol

This is the coolest e-mail ever
`, knownConcealedEmail, knownConcealedEmail2)

	testReadEmailGateway := TestReadEmailGateway{
		ReadEmailReturnByte: []byte(email),
	}
	testAppContext.Bind(func() ReadEmailGateway {
		return &testReadEmailGateway
	})

	testSendEmailGateway := TestSendEmailGateway{}
	testAppContext.Bind(func() SendEmailGateway {
		return &testSendEmailGateway
	})

	testConfigurationGateway := TestConfigurationGateway{
		GetRealEmailReturn: map[string][]*string{
			"known": {&actualEmail1, &actualDescription1},
			"known2": {&actualEmail2, &actualDescription2},
		},
	}
	testAppContext.Bind(func() ConfigurationGateway {
		return &testConfigurationGateway
	})

	testEnvironmentGateway := TestEnvironmentGateway{
		GetEnvironmentValueReturn: map[string]string{
			"DOMAIN": "apple.com",
		},
	}
	testAppContext.Bind(func() context.EnvironmentGateway {
		return &testEnvironmentGateway
	})

	err := usecase.ForwardEmail("https://email.com")

	if err != nil {
		t.Errorf("No error should have been returned from the ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}

	rawForwardedEmail := testSendEmailGateway.SendEmailEmail

	if !bytes.Contains(rawForwardedEmail, []byte(actualDescription1)) {
		t.Errorf("The actual e-mail recipient's description %s is missing from the e-mail and it should have been there", actualDescription1)
	}

	if !bytes.Contains(rawForwardedEmail, []byte(actualDescription2)) {
		t.Errorf("The actual e-mail recipient's description %s is missing from the e-mail and it should have been there", actualDescription2)
	}

	if bytes.Contains(rawForwardedEmail, []byte(", \r\n")) {
		t.Errorf("Incorrectly formatted the To header, there's a trailing comma")
	}
}

func TestToHeaderDoesNotChangeUnknownEmailDescription(t *testing.T) {
	resetBaseDependencies()

	unknownEmail := "sugar@other.com"
	unknownDescription := "Not Hear"

	knownConcealedEmail := "known@apple.com"  //need at least one known e-mail for the logic to get to the header munging logic
	actualEmail := "moof@dogcow.com"
	actualDescription := "The coolest description"

	email := fmt.Sprintf(`To: %s <%s>, %s
From: moof@apple.com
Subject: lol

This is the coolest e-mail ever
`, unknownDescription, unknownEmail, knownConcealedEmail)

	testReadEmailGateway := TestReadEmailGateway{
		ReadEmailReturnByte: []byte(email),
	}
	testAppContext.Bind(func() ReadEmailGateway {
		return &testReadEmailGateway
	})

	testSendEmailGateway := TestSendEmailGateway{}
	testAppContext.Bind(func() SendEmailGateway {
		return &testSendEmailGateway
	})

	testConfigurationGateway := TestConfigurationGateway{
		GetRealEmailReturn: map[string][]*string{
			"known": {&actualEmail, &actualDescription},
		},
	}
	testAppContext.Bind(func() ConfigurationGateway {
		return &testConfigurationGateway
	})

	testEnvironmentGateway := TestEnvironmentGateway{
		GetEnvironmentValueReturn: map[string]string{
			"DOMAIN": "apple.com",
		},
	}
	testAppContext.Bind(func() context.EnvironmentGateway {
		return &testEnvironmentGateway
	})

	err := usecase.ForwardEmail("https://email.com")

	if err != nil {
		t.Errorf("No error should have been returned from the ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}

	rawForwardedEmail := testSendEmailGateway.SendEmailEmail
	rawForwardedEmailString := string(rawForwardedEmail)
	fmt.Println(rawForwardedEmailString)

	if !bytes.Contains(rawForwardedEmail, []byte(unknownEmail)) {
		t.Errorf("The unknown e-mail %s is missing from the e-mail and it should have been there", unknownEmail)
	}
	if !bytes.Contains(rawForwardedEmail, []byte(unknownDescription)) {
		t.Errorf("The unknown description %s is missing from the e-mail and it should have been there", unknownDescription)
	}
}

func TestToHeaderDoesChangeKnownEmailDescriptionWithDescriptionAlready(t *testing.T) {
	resetBaseDependencies()

	knownConcealedEmail := "known@apple.com"
	knownConcealedDescription := "Nope"

	actualEmail := "moof@dogcow.com"
	actualDescription := "The coolest description"

	email := fmt.Sprintf(`To: %s <%s>
From: moof@apple.com
Subject: lol

This is the coolest e-mail ever
`, knownConcealedDescription, knownConcealedEmail)

	testReadEmailGateway := TestReadEmailGateway{
		ReadEmailReturnByte: []byte(email),
	}
	testAppContext.Bind(func() ReadEmailGateway {
		return &testReadEmailGateway
	})

	testSendEmailGateway := TestSendEmailGateway{}
	testAppContext.Bind(func() SendEmailGateway {
		return &testSendEmailGateway
	})

	testConfigurationGateway := TestConfigurationGateway{
		GetRealEmailReturn: map[string][]*string{
			"known": {&actualEmail, &actualDescription},
		},
	}
	testAppContext.Bind(func() ConfigurationGateway {
		return &testConfigurationGateway
	})

	testEnvironmentGateway := TestEnvironmentGateway{
		GetEnvironmentValueReturn: map[string]string{
			"DOMAIN": "apple.com",
		},
	}
	testAppContext.Bind(func() context.EnvironmentGateway {
		return &testEnvironmentGateway
	})

	err := usecase.ForwardEmail("https://email.com")

	if err != nil {
		t.Errorf("No error should have been returned from the ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}

	rawForwardedEmail := testSendEmailGateway.SendEmailEmail

	if !bytes.Contains(rawForwardedEmail, []byte(actualDescription)) {
		t.Errorf("The actual e-mail recipient's description %s is missing from the e-mail and it should have been there", actualDescription)
	}

	if bytes.Contains(rawForwardedEmail, []byte(knownConcealedDescription)) {
		t.Errorf("The original e-mail recipient's description %s is still in the e-mail but it shouldn't have been there", knownConcealedDescription)
	}
}

func TestToHeaderDoesNotChangeKnownEmailDescriptionWithNoDescriptionReturned(t *testing.T) {
	resetBaseDependencies()

	knownConcealedEmail := "known@apple.com"
	knownConcealedDescription := "Nope"

	actualEmail := "moof@dogcow.com"

	email := fmt.Sprintf(`To: %s <%s>
From: moof@apple.com
Subject: lol

This is the coolest e-mail ever
`, knownConcealedDescription, knownConcealedEmail)

	testReadEmailGateway := TestReadEmailGateway{
		ReadEmailReturnByte: []byte(email),
	}
	testAppContext.Bind(func() ReadEmailGateway {
		return &testReadEmailGateway
	})

	testSendEmailGateway := TestSendEmailGateway{}
	testAppContext.Bind(func() SendEmailGateway {
		return &testSendEmailGateway
	})

	testConfigurationGateway := TestConfigurationGateway{
		GetRealEmailReturn: map[string][]*string{
			"known": {&actualEmail, nil},  //notice the nil here, that means the description has not been set
		},
	}
	testAppContext.Bind(func() ConfigurationGateway {
		return &testConfigurationGateway
	})

	testEnvironmentGateway := TestEnvironmentGateway{
		GetEnvironmentValueReturn: map[string]string{
			"DOMAIN": "apple.com",
		},
	}
	testAppContext.Bind(func() context.EnvironmentGateway {
		return &testEnvironmentGateway
	})

	err := usecase.ForwardEmail("https://email.com")

	if err != nil {
		t.Errorf("No error should have been returned from the ForwardEmailUsecase")
		t.Errorf("Instead this was returned %+v", err)
	}

	rawForwardedEmail := testSendEmailGateway.SendEmailEmail

	if !bytes.Contains(rawForwardedEmail, []byte(knownConcealedDescription)) {
		t.Errorf("The original e-mail recipient's description %s is missing from the e-mail and it should have been there", knownConcealedDescription)
	}
}

func contains(slice []string, item string) bool {
	for _, currentItem := range slice {
		if currentItem == item {
			return true
		}
	}

	return false
}
