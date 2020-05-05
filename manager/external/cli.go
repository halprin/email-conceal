package external

import (
	"os"
)

func Cli() {
	applicationContext := &CliApplicationContext{}

	applicationContext.ConcealEmailController(os.Args)
}
