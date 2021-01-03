package osEnvironmentVariable

import (
	"os"
)


type OsEnvironmentGateway struct {}

func (receiver OsEnvironmentGateway) GetEnvironmentValue(key string) string {
	return os.Getenv(key)
}
