package restContext

import (
	"github.com/halprin/email-conceal/manager/context"
	"github.com/halprin/email-conceal/manager/usecases"
)

type RestApplicationContextUsecases struct{
	ParentContext context.ApplicationContext
}

func (appContextUsecases *RestApplicationContextUsecases) AddConcealEmailUsecase(email string) (string, error) {
	return usecases.AddConcealEmailUsecase(email, appContextUsecases.ParentContext)
}

func (appContextUsecases *RestApplicationContextUsecases) DeleteConcealEmailUsecase(concealPrefix string) error {
	return usecases.DeleteConcealEmailMappingUsecase(concealPrefix, appContextUsecases.ParentContext)
}