package external

import (
	"os"
)

func Cli() {
	applicationContext := &CliApplicationContext{}

	applicationContext.ForwardEmailGateway(os.Args)
}
