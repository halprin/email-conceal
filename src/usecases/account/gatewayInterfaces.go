package account

type AccountConfigurationGateway interface {
	AddAccount(emailUsername string, passwordHash string) error
}
