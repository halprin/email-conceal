package main

import (
	"github.com/halprin/email-conceal/src/external/sqsQueue"
)

func main() {
	sqsQueue.SqsQueueListener()
	//uncomment below to test locally
	//localFileWatch.LocalFileWatcher()
}
