package external

import (
	"os"
)

func Cli() {
	applicationContext := &CliApplicationContext{}

	_ = applicationContext.ForwardEmailGateway(os.Args)
}
