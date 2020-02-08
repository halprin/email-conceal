package usecases

import (
	"email-conceal/context"
	"fmt"
)

func ConcealEmail(sourceEmail string, applicationContext context.ApplicationContext) string {
	concealedEmailPrefix := applicationContext.GenerateRandomUuid()
	return fmt.Sprintf("%s@asdf.net", concealedEmailPrefix)
}
