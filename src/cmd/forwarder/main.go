package main

import (
	"github.com/halprin/email-conceal/src/external/localFileWatch"
	"github.com/halprin/email-conceal/src/external/sqsQueue"
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
	//uncomment below to test locally
	//localFileWatch.LocalFileWatcher()
}
