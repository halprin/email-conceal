package lib

import (
	"github.com/google/uuid"
	"github.com/halprin/email-conceal/context"
)

func GenerateGoogleRandomUuid(applicationContext context.ApplicationContext) string {
	randomUuid := uuid.New()
	return randomUuid.String()
}
