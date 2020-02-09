package external

import (
	"os"
)

func Cli() {
	applicationContext := &CliApplicationContext{}

	applicationContext.ConcealEmailGateway(os.Args)
}
