package forwardEmail

type ReadEmailGateway interface {
	ReadEmail(uri string) ([]byte, error)
}
