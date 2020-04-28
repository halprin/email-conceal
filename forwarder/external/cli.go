package external

import (
	"os"
)

func Cli() {
	applicationContext := &CliApplicationContext{}

	arguments := map[string]interface{}{
		"url": os.Args[1],
	}
	_ = applicationContext.ForwardEmailController(arguments)
}
