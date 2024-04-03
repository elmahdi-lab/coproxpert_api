package gcp

import (
	"cloud.google.com/go/storage"
	"context"
	"google.golang.org/api/option"
	"io"
	"os"
	"time"
)

type FileManager struct {
	client *storage.Client
}

func NewFileManager() (*FileManager, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(os.Getenv("GCP_CREDENTIALS_PATH")))
	if err != nil {
		return nil, err
	}
	return &FileManager{client: client}, nil
}

func (f *FileManager) Upload(file io.Reader, filename string) error {
	ctx := context.Background()
	bucketName := os.Getenv("GCP_BUCKET_NAME")

	wc := f.client.Bucket(bucketName).Object(filename).NewWriter(ctx)

	if _, err := io.Copy(wc, file); err != nil {
		return err
	}

	if err := wc.Close(); err != nil {
		return err
	}
	return nil
}

func (f *FileManager) Link(filename string) (string, error) {
	bucketName := os.Getenv("GCP_BUCKET_NAME")
	expirationTime := time.Now().Add(15 * time.Minute)
	url, err := storage.SignedURL(bucketName, filename, &storage.SignedURLOptions{
		GoogleAccessID: os.Getenv("GCP_SERVICE_ACCOUNT_EMAIL"),
		PrivateKey:     []byte(os.Getenv("GCP_PRIVATE_KEY")),
		Method:         "GET",
		Expires:        expirationTime,
	})
	if err != nil {
		return "", err
	}
	return url, nil
}

func (f *FileManager) Delete(filename string) error {
	ctx := context.Background()
	bucketName := os.Getenv("GCP_BUCKET_NAME")

	err := f.client.Bucket(bucketName).Object(filename).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}
