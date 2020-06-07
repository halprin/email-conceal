package context

type ApplicationContextUsecases interface {
	AddConcealEmail(email string, description *string) (string, error)
	DeleteConcealEmail(concealPrefix string) error
}
