package pubsub

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"log"
)

func CreateMessage(message string, topicName string) {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := "schema-on-read"

	// Creates a client.
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// get topic iterator and find the topic that we need
	it := client.Topics(ctx)
	var topic *pubsub.Topic

	for {
		t, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		if t.String() == topicName {
			topic = t
		}
	}

	topic.Publish(ctx, &pubsub.Message{Data: []byte(message)})

	fmt.Printf("Message sent: %v.\n", message)
}
