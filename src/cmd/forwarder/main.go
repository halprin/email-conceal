package main

import (
	"github.com/halprin/email-conceal/src/external/forwarder/localFileWatch"
	"github.com/halprin/email-conceal/src/external/forwarder/sqsQueue"
	"os"
)

func main() {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "local" {
		localFileWatch.Init()
		localFileWatch.LocalFileWatcher()
		return
	}

	sqsQueue.Init()
	sqsQueue.SqsQueueListener()
}
