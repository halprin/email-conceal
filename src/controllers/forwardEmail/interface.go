package forwardEmail

type ForwardEmail interface {
	ForwardEmail(arguments map[string]interface{}) error
}
