package actualEmail

import "github.com/halprin/email-conceal/src/context"


var controller = ActualEmailController{}
var testAppContext = context.ApplicationContext{}

type TestActualEmailUsecase struct {
	AddReceiveActualEmail string
	AddReturnError        error
}

func (testUsecase *TestActualEmailUsecase) Add(actualEmail string) error {
	testUsecase.AddReceiveActualEmail = actualEmail
	return testUsecase.AddReturnError
}
