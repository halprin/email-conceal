package testApplicationContext


type TestApplicationContextUsecases struct{
	ReceivedConcealEmailUsecaseEmail   string
	ReturnFromConcealEmailUsecase      string
	ReturnErrorFromConcealEmailUsecase error

	ReceivedDeleteConcealEmailUsecaseConcealPrefixArgument string
	ReturnErrorFromDeleteConcealEmailUsecase               error

	ReceivedAddDescriptionUsecaseConcealPrefixArgument string
	ReceivedAddDescriptionUsecaseDescriptionArgument string
	ReturnErrorFromAddDescriptionUsecase               error
}

func (appContextUsecases *TestApplicationContextUsecases) AddConcealEmail(email string, description *string) (string, error) {
	appContextUsecases.ReceivedConcealEmailUsecaseEmail = email
	return appContextUsecases.ReturnFromConcealEmailUsecase, appContextUsecases.ReturnErrorFromConcealEmailUsecase
}

func (appContextUsecases *TestApplicationContextUsecases) DeleteConcealEmail(concealPrefix string) error {
	appContextUsecases.ReceivedDeleteConcealEmailUsecaseConcealPrefixArgument = concealPrefix
	return appContextUsecases.ReturnErrorFromDeleteConcealEmailUsecase
}

func (appContextUsecases *TestApplicationContextUsecases) AddDescriptionToExistingEmail(concealPrefix string, description string) error {
	appContextUsecases.ReceivedAddDescriptionUsecaseConcealPrefixArgument = concealPrefix
	appContextUsecases.ReceivedAddDescriptionUsecaseDescriptionArgument = description
	return appContextUsecases.ReturnErrorFromAddDescriptionUsecase
}
