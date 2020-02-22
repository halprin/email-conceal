package gateways

import (
	"github.com/halprin/email-conceal/forwarder/context"
	"os"
)

func OsEnvEnvironmentGateway(key string, applicationContext context.ApplicationContext) string {
	return os.Getenv(key)
}
