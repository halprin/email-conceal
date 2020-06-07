package restContext

import (
	"github.com/halprin/email-conceal/manager/context"
	"github.com/halprin/email-conceal/manager/external/lib"
	"os"
)

type RestApplicationContextLibraries struct{
	ParentContext context.ApplicationContext
}

func (appContextLibraries *RestApplicationContextLibraries) GenerateRandomUuid() string {
	return lib.GenerateGoogleRandomUuid(appContextLibraries.ParentContext)
}

func (appContextLibraries *RestApplicationContextLibraries) Exit(returnCode int) {
	os.Exit(returnCode)
}
