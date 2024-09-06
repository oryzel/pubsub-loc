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
	topicInboundNotificationID        = "galadriel-inbound-order-status-notification"
	subscriptionInboundNotificationID = "portal-inbound-order-status-notification-faramir-process-inbound-order"
)

func (s *service) InboundOrderStatusNotification(ctx context.Context) {

	topic := s.pubSubClient.Topic(topicInboundNotificationID)
	exists, err := topic.Exists(ctx)
	if err != nil {
		log.Fatalf("Error checking topic existence: %v", err)
	}
	if !exists {
		_, err := s.pubSubClient.CreateTopic(ctx, topicInboundNotificationID)
		if err != nil {
			log.Fatalf("Failed to create topic: %v", err)
		}
	}

	payload := generateInboundOrderStatusNotificationPayload()
	payloadByte, _ := json.Marshal(payload)
	fmt.Println(payloadByte)

	utils.PubSubPublish(ctx, topic, payloadByte)
	//utils.PubSubSubscribe(ctx, s.pubSubClient, topic, subscriptionInboundNotificationID)

}

func generateInboundOrderStatusNotificationPayload() model.InboundOrderStatusNotificationRequest {
	skus := []model.InboundStatusNotificationSKU{
		{
			Code:   "IK-12",
			Qty:    7,
			Status: "Good",
		},
		{
			Code:   "IK-12",
			Qty:    3,
			Status: "Reject",
		},
	}
	request := model.InboundOrderStatusNotificationRequest{
		ID:        "id",
		Publisher: "publisher",
		Action:    "action",
		Data: model.InboundOrderStatusNotificationData{
			UserID: 12,
			DocNo:  "IN20240906088964888",
			Status: "10",
			SKUS:   skus,
		},
	}

	return request
}
