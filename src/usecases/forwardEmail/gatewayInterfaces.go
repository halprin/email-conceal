package forwardEmail

type ReadEmailGateway interface {
	ReadEmail(uri string) ([]byte, error)
}

type SendEmailGateway interface {
	SendEmail(email []byte, recipients []string) error
}

type ConfigurationGateway interface {
	GetRealEmailAddressForConcealPrefix(concealedRecipientPrefix string) (string, *string, error)
}
