package gateways

import (
	"github.com/halprin/email-conceal/manager/context"
	"os"
)

func OsEnvEnvironmentGateway(key string, applicationContext context.ApplicationContext) string {
	return os.Getenv(key)
}
