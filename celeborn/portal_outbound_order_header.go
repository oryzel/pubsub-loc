package celeborn

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/oryzel/pubsub-loc/celeborn/model"
	"github.com/oryzel/pubsub-loc/utils"
)

const (
	topicPortalOutboundOrderHeaderID        = "portal-outbound-order-header"
	subscriptionPortalOutboundOrderHeaderID = "portal-outbound-order-header-celeborn-process-sales-order"
)

func (s *service) PortalOutboundOrderHeaderNotification(ctx context.Context) {
	topic := s.pubSubClient.Topic(topicPortalOutboundOrderHeaderID)
	exists, err := topic.Exists(ctx)
	if err != nil {
		log.Fatalf("Error checking topic existence: %v", err)
	}
	if !exists {
		_, err := s.pubSubClient.CreateTopic(ctx, topicPortalOutboundOrderHeaderID)
		if err != nil {
			log.Fatalf("Failed to create topic: %v", err)
		}
	}

	payload := generatePortalOutboundOrderHeaderNotificationPayload()
	payloadByte, _ := json.Marshal(payload)
	fmt.Println(payloadByte)

	utils.PubSubPublish(ctx, topic, payloadByte)
	//utils.PubSubSubscribe(ctx, s.pubSubClient, topic, subscriptionPortalOutboundOrderHeaderID)

}

func generatePortalOutboundOrderHeaderNotificationPayload() model.PortalOutboundOrderHeaderRequest {
	request := model.PortalOutboundOrderHeaderRequest{
		ID:        "id",
		Publisher: "publisher",
		Action:    "action",
		Body: model.PortalOutboundOrderHeaderData{
			DocNo:       "SO20240704460378279",
			OldStatus:   "",
			Status:      "12",
			OrderNumber: "ABI2024070402",
			OrderSource: "NPLT",
		},
	}

	return request
}
