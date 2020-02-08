package usecases

import (
	"fmt"
	"github.com/halprin/email-conceal/context"
)

func ConcealEmail(sourceEmail string, applicationContext context.ApplicationContext) string {
	concealedEmailPrefix := applicationContext.GenerateRandomUuid()
	return fmt.Sprintf("%s@asdf.net", concealedEmailPrefix)
}
