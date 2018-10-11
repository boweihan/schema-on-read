package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"log"
)

func CreateTopic(name string) {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := "schema-on-read"

	// Creates a client.
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the name for the new topic.
	topicName := name

	// Creates the new topic.
	topic, err := client.CreateTopic(ctx, topicName)
	if err != nil {
		log.Fatalf("Failed to create topic: %v", err)
	}

	fmt.Printf("Topic %v created.\n", topic)
}
