package lib

import (
	"email-conceal/context"
	"github.com/google/uuid"
)

func GenerateGoogleRandomUuid(applicationContext context.ApplicationContext) string {
	randomUuid := uuid.New()
	return randomUuid.String()
}
