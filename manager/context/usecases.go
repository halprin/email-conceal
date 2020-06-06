package context

type ApplicationContextUsecases interface {
	AddConcealEmailUsecase(email string) (string, error)
	DeleteConcealEmailUsecase(concealPrefix string) error
}
