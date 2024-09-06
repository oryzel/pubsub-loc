package utils

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
)

func PubSubPublish(ctx context.Context, topic *pubsub.Topic, payload []byte) string {

	result := topic.Publish(ctx, &pubsub.Message{
		Data: payload,
	})
	id, err := result.Get(ctx)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}
	fmt.Printf("Published a message with ID: %s\n", id)

	return id
}

func PubSubSubscribe(ctx context.Context, client pubsub.Client, topic *pubsub.Topic, subscriptionID string) {

	// Create a subscription if it doesn't exist.
	subscription := client.Subscription(subscriptionID)
	exists, err := subscription.Exists(ctx)
	if err != nil {
		log.Fatalf("Error checking subscription existence: %v", err)
	}
	if !exists {
		_, err := client.CreateSubscription(ctx, subscriptionID, pubsub.SubscriptionConfig{
			Topic:       topic,
			AckDeadline: 10 * time.Second,
			DeadLetterPolicy: &pubsub.DeadLetterPolicy{
				DeadLetterTopic:     "",
				MaxDeliveryAttempts: 1,
			},
		})
		if err != nil {
			log.Fatalf("Failed to create subscription: %v", err)
		}
	}

	// Pull and process messages from the subscription.
	err = subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Printf("Received message: %s\n", string(msg.Data))
		msg.Ack()
	})
	if err != nil {
		log.Fatalf("Error receiving messages: %v", err)
	}
}
