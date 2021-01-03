package concealEmail

type ConcealEmailGateway interface {
	AddConcealedEmailToActualEmailMapping(concealPrefix string, actualEmail string, description *string) error
	DeleteConcealedEmailToActualEmailMapping(concealPrefix string) error
	UpdateConcealedEmail(concealPrefix string, description *string) error
}
