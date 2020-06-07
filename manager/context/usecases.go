package context

type ApplicationContextUsecases interface {
	AddConcealEmail(email string) (string, error)
	DeleteConcealEmail(concealPrefix string) error
}
