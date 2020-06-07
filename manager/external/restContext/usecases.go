package restContext

import (
	"github.com/halprin/email-conceal/manager/context"
	"github.com/halprin/email-conceal/manager/usecases"
)

type RestApplicationContextUsecases struct{
	ParentContext context.ApplicationContext
}

func (appContextUsecases *RestApplicationContextUsecases) AddConcealEmail(email string) (string, error) {
	return usecases.AddConcealEmailUsecase(email, appContextUsecases.ParentContext)
}

func (appContextUsecases *RestApplicationContextUsecases) DeleteConcealEmail(concealPrefix string) error {
	return usecases.DeleteConcealEmailMappingUsecase(concealPrefix, appContextUsecases.ParentContext)
}