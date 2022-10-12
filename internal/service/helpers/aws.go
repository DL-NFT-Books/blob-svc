package helpers

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"gitlab.com/tokend/nft-books/blob-svc/internal/config"
	"mime/multipart"
)
import "github.com/aws/aws-sdk-go/aws/session"

func NewAWSSession(config *config.AWSConfig) *session.Session {
	return session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(
			config.AccessKeyID,
			config.SecretKeyID,
			""),
		Region:     aws.String(config.Region),
		DisableSSL: aws.Bool(config.SslDisable),
	}))
}

func UploadFile(file multipart.File, key string, config *config.AWSConfig) error {
	awsSession := NewAWSSession(config)
	uploader := s3manager.NewUploader(awsSession)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.Bucket),
		Key:    aws.String(key),
		Body:   file,
	})

	return err
}
