package s3FileReader

import "github.com/aws/aws-sdk-go/aws/session"

var awsSession, sessionErr = session.NewSession()
