package files

import (
	"bytes"
	"io"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AWSFileManager struct {
	svc *s3.S3
}

func NewAWSFileManager() (*AWSFileManager, error) {
	awsAccessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsRegion := os.Getenv("AWS_REGION")

	// Create an AWS session
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, ""),
	})
	if err != nil {
		return nil, err
	}

	// Create an S3 client
	svc := s3.New(sess)

	return &AWSFileManager{svc: svc}, nil
}

func (f *AWSFileManager) Upload(file io.Reader, filename string) error {
	// Read file content into memory
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// Create a bytes.Reader from the file content
	fileReader := bytes.NewReader(fileContent)

	// Upload file to S3 bucket
	_, err = f.svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Key:    aws.String(filename),
		Body:   fileReader,
	})
	if err != nil {
		return err
	}

	return nil
}

func (f *AWSFileManager) Link(filename string) (string, error) {
	req, _ := f.svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Key:    aws.String(filename),
	})
	url, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (f *AWSFileManager) Delete(filename string) error {
	// Delete file from S3 bucket
	_, err := f.svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(os.Getenv("AWS_BUCKET_NAME")),
		Key:    aws.String(filename),
	})
	if err != nil {
		return err
	}

	return nil
}
