package events

import (
	"encoding/json"
	"fmt"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"ithumans.com/coproxpert/cmd"
	"ithumans.com/coproxpert/models"
)

type PubSubEntity string
type EventType string

const (
	Created EventType = "created"
	Updated EventType = "updated"
	Deleted EventType = "deleted"
)

type PubSubMessage struct {
	UserID     uuid.UUID         `json:"UserID"`
	EntityID   uuid.UUID         `json:"EntityID"`
	EntityName models.EntityName `json:"EntityName"`
	EventType  EventType         `json:"EventType"`
}

func PublishMessage(userID uuid.UUID, entityID uuid.UUID, entityName models.EntityName, eventType EventType) error {
	client, err := cmd.NewPubSubClient(context.Background()) // Use a global context

	gcpTopic := os.Getenv("GOOGLE_PUBSUB_TOPIC")

	topic := client.Topic(gcpTopic)

	msg := PubSubMessage{
		UserID:     userID,
		EntityName: entityName,
		EntityID:   entityID,
		EventType:  eventType,
	}

	msgBytes, err := json.Marshal(msg)

	// Publish the message
	result := topic.Publish(context.Background(), &pubsub.Message{
		Data: msgBytes,
	})

	// Wait for the result
	_, err = result.Get(context.Background())
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	return nil
}
