package main

import (
	"github.com/halprin/email-conceal/src/external/manager/localRest"
	"github.com/halprin/email-conceal/src/external/manager/rest"
	"os"
)

func main() {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "local" {
		localRest.Init()
		localRest.Rest()
		return
	}

	rest.Init()
	rest.Rest()
}
