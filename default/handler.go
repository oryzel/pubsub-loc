package _default

import (
	"context"
	"github.com/oryzel/pubsub-loc/utils"
	"log"
)

const (
	topicID        = "topic-default"
	subscriptionID = "topic-default-subscription"
)

func (s *service) Publish(ctx context.Context) {

	// Create a topic if it doesn't exist.
	topic := s.pubSubClient.Topic(topicID)
	exists, err := topic.Exists(ctx)
	if err != nil {
		log.Fatalf("Error checking topic existence: %v", err)
	}
	if !exists {
		_, err := s.pubSubClient.CreateTopic(ctx, topicID)
		if err != nil {
			log.Fatalf("Failed to create topic: %v", err)
		}
	}

	utils.PubSubPublish(ctx, topic, []byte("Hello, Pub/Sub!"))
	utils.PubSubSubscribe(ctx, s.pubSubClient, topic, subscriptionID)

}
