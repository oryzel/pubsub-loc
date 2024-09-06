package boromir

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/oryzel/pubsub-loc/boromir/model"
	"github.com/oryzel/pubsub-loc/utils"
	"log"
)

const (
	topicID        = "boromir-sync-order"
	subscriptionID = "boromir-sync-order-boromir-process-sync-order"
)

func (s *service) SyncOrder(ctx context.Context) {

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

	payload := generateSyncOrderPayload()
	payloadByte, _ := json.Marshal(payload)
	fmt.Println(payloadByte)

	utils.PubSubPublish(ctx, topic, payloadByte)
	utils.PubSubSubscribe(ctx, s.pubSubClient, topic, subscriptionID)

}

func generateSyncOrderPayload() model.SyncOrderMessage {

	shop := model.Shop{
		Id:                5,
		UserId:            12,
		WmsCustomerId:     "CUSTOMER12",
		MarketplaceShopId: 1072972126,
		PortalShopId:      1708,
		AccessToken:       "54476a656e41684370536d6141714648",
	}

	return model.SyncOrderMessage{
		model.SyncOrderData{
			Shop: shop,
			OrderSNList: []string{
				"240202ES48PM1K",
				"240202FY68FKN9",
			},
		},
	}

}
