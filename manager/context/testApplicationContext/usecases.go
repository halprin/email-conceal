package testApplicationContext


type TestApplicationContextUsecases struct{
	ReceivedConcealEmailUsecaseEmail   string
	ReturnFromConcealEmailUsecase      string
	ReturnErrorFromConcealEmailUsecase error

	ReceivedDeleteConcealEmailUsecaseConcealPrefixArgument string
	ReturnErrorFromDeleteConcealEmailUsecase               error
}

func (appContextUsecases *TestApplicationContextUsecases) AddConcealEmail(email string) (string, error) {
	appContextUsecases.ReceivedConcealEmailUsecaseEmail = email
	return appContextUsecases.ReturnFromConcealEmailUsecase, appContextUsecases.ReturnErrorFromConcealEmailUsecase
}

func (appContextUsecases *TestApplicationContextUsecases) DeleteConcealEmail(concealPrefix string) error {
	appContextUsecases.ReceivedDeleteConcealEmailUsecaseConcealPrefixArgument = concealPrefix
	return appContextUsecases.ReturnErrorFromDeleteConcealEmailUsecase
}
