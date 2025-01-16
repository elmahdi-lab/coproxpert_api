package cmd

import (
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

// GCPClients wraps individual clients for different Google Cloud services.
type GCPClients struct {
	StorageClient *storage.Client
	PubSubClient  *pubsub.Client
}

// NewStorageClient initializes and returns a new Google Cloud Storage client.
func NewStorageClient(ctx context.Context) (*storage.Client, error) {
	key := os.Getenv("GOOGLE_JSON_KEY")
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(key))
	if err != nil {
		log.Printf("Failed to create Google Cloud Storage client: %v", err)
		return nil, err
	}
	return client, nil
}

// NewPubSubClient initializes and returns a new Google Cloud Pub/Sub client.
func NewPubSubClient(ctx context.Context) (*pubsub.Client, error) {
	key := os.Getenv("GOOGLE_JSON_KEY")
	projectId := os.Getenv("GOOGLE_CLOUD_PROJECT_ID")

	client, err := pubsub.NewClient(ctx, projectId, option.WithCredentialsFile(key))
	if err != nil {
		log.Printf("Failed to create Google Cloud Pub/Sub client: %v", err)
		return nil, err
	}
	return client, nil
}

// TestStorageConnection tests the connection to Google Cloud Storage.
func TestStorageConnection(ctx context.Context) bool {
	client, _ := NewStorageClient(ctx)
	it := client.Buckets(ctx, os.Getenv("GOOGLE_CLOUD_PROJECT_ID"))
	_, err := it.Next()
	if err != nil {
		log.Printf("Failed to connect to Google Cloud Storage: %v", err)
		return false
	}
	log.Println("Successfully connected to Google Cloud Storage")
	return true
}

// TestPubSubConnection tests the connection to Google Cloud Pub/Sub.
func TestPubSubConnection(ctx context.Context) bool {
	client, _ := NewPubSubClient(ctx)
	topicIterator := client.Topics(ctx)
	_, err := topicIterator.Next()
	if err != nil {
		log.Printf("Failed to connect to Google Cloud Pub/Sub: %v", err)
		return false
	}
	log.Println("Successfully connected to Google Cloud Pub/Sub")
	return true
}
