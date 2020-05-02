package gateways

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/halprin/email-conceal/forwarder/context"
	"log"
	"regexp"
)

var s3Downloader = s3manager.NewDownloader(awsSession)

var s3UrlParseRegex = regexp.MustCompile(`^s3://([^/]+)(/.+)$`)

func S3ReadEmailGateway(url string, applicationContext context.ApplicationContext) ([]byte, error) {
	if sessionErr != nil {
		return nil, sessionErr
	}

	bucket, object := parseBucketAndObjectFromS3Url(url)

	var internalBuffer []byte
	inMemoryBuffer := aws.NewWriteAtBuffer(internalBuffer)

	log.Printf("Downloading S3 file at %s", url)
	_, err := s3Downloader.Download(inMemoryBuffer, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(object),
	})

	if err != nil {
		return nil, err
	}
	log.Printf("Download S3 file complete at %s", url)

	return inMemoryBuffer.Bytes(), nil
}

func parseBucketAndObjectFromS3Url(s3Url string) (string, string) {
	matchingStrings := s3UrlParseRegex.FindStringSubmatch(s3Url)
	return matchingStrings[1], matchingStrings[2]
}
