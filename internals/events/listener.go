package events

import (
	"encoding/json"
	"fmt"
	"os"

	"cloud.google.com/go/pubsub"
	"golang.org/x/net/context"
	"ithumans.com/coproxpert/cmd"
	"ithumans.com/coproxpert/data/models"
)

type Subscriber interface {
	EntityName()
	HandleMessage(message PubSubMessage)
}

func ListenAndDispatch() error {
	ctx := context.Background()

	// Create Pub/Sub client
	client, err := cmd.NewPubSubClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create Pub/Sub client: %w", err)
	}
	defer func() {
		if err := client.Close(); err != nil {
			fmt.Printf("failed to close Pub/Sub client: %v\n", err)
		}
	}()

	// Get subscription name from environment
	gcpSubscription := os.Getenv("GOOGLE_PUBSUB_SUBSCRIPTION")
	if gcpSubscription == "" {
		return fmt.Errorf("environment variable GOOGLE_PUBSUB_SUBSCRIPTION is not set")
	}

	sub := client.Subscription(gcpSubscription)

	// Define the mapping of PubSub entities to their subscribers
	subscriberMap := map[models.EntityName]Subscriber{
		models.UnitGroupEntity: &UnitGroupSubscriber{},
		models.UnitEntity:      &UnitSubscriber{},
	}

	// Receive messages
	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		var pubSubMessage PubSubMessage

		// Unmarshal message data
		if err := json.Unmarshal(msg.Data, &pubSubMessage); err != nil {
			fmt.Printf("failed to unmarshal message: %v\n", err)
			msg.Nack() // Negative acknowledgement
			return
		}

		// Find the appropriate subscriber for the entity
		subscriber, exists := subscriberMap[pubSubMessage.EntityName]
		if !exists {
			fmt.Printf("no subscriber found for entity: %s\n", pubSubMessage.EntityName)
			fmt.Printf("Raw message received: %s\n", string(msg.Data))
			msg.Nack()
			return
		}

		subscriber.HandleMessage(pubSubMessage)
		msg.Ack()
	})

	if err != nil {
		return fmt.Errorf("failed to receive messages: %w", err)
	}

	return nil
}
