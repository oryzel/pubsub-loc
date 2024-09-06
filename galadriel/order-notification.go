package galadriel

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/oryzel/pubsub-loc/galadriel/model"
	"github.com/oryzel/pubsub-loc/utils"
)

const (
	topicOrderNotificationID        = "galadriel-sales-order-status-notification"
	subscriptionOrderNotificationID = "galadriel-sales-order-status-notification-gandalf-process-sales-order"
)

func (s *service) OrderStatusNotification(ctx context.Context) {

	topic := s.pubSubClient.Topic(topicOrderNotificationID)
	exists, err := topic.Exists(ctx)
	if err != nil {
		log.Fatalf("Error checking topic existence: %v", err)
	}
	if !exists {
		_, err := s.pubSubClient.CreateTopic(ctx, topicOrderNotificationID)
		if err != nil {
			log.Fatalf("Failed to create topic: %v", err)
		}
	}

	payload := generateOrderStatusNotificationPayload()
	payloadByte, _ := json.Marshal(payload)
	fmt.Println(payloadByte)

	utils.PubSubPublish(ctx, topic, payloadByte)
	//utils.PubSubSubscribe(ctx, s.pubSubClient, topic, subscriptionOrderNotificationID)
}

func generateOrderStatusNotificationPayload() model.SalesOrderStatusNotificationRequest {
	request := model.SalesOrderStatusNotificationRequest{
		ID:        "id",
		Publisher: "publisher",
		Action:    "action",
		Body: model.SalesOrderStatusNotificationData{
			ShopID: 1957,
			DocNo:  "SO20240829510662635",
			Status: "18",
		},
	}

	return request
}
