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
	topicSKUStockUpdatedID        = "galadriel-sku-stock-updated"
	subscriptionSKUStockUpdatedID = "galadriel-sku-stock-updated-faramir-sync-stock"
)

func (s *service) SKUSyncStock(ctx context.Context) {

	topic := s.pubSubClient.Topic(topicSKUStockUpdatedID)
	exists, err := topic.Exists(ctx)
	if err != nil {
		log.Fatalf("Error checking topic existence: %v", err)
	}
	if !exists {
		_, err := s.pubSubClient.CreateTopic(ctx, topicSKUStockUpdatedID)
		if err != nil {
			log.Fatalf("Failed to create topic: %v", err)
		}
	}

	payload := generateSkuSyncStockPayload()
	payloadByte, _ := json.Marshal(payload)
	fmt.Println(payloadByte)

	utils.PubSubPublish(ctx, topic, payloadByte)
	//utils.PubSubSubscribe(ctx, s.pubSubClient, topic, subscriptionSKUStockUpdatedID)
}

func generateSkuSyncStockPayload() model.StockUpdateRequest {
	stockUpdateRequest := make([]model.StockUpdateRequestData, 0)
	stockUpdateRequest = append(stockUpdateRequest, model.StockUpdateRequestData{
		UserID:      "12",
		SkuID:       309666,
		SKU:         "BUKU-12",
		WarehouseID: "ID_IDCOMGW013",
		Stock:       9,
		UpdatedAt:   "2024-08-21 07:03:51",
	})

	request := model.StockUpdateRequest{
		ID:        "id",
		Publisher: "publisher",
		Action:    "action",
		Data:      stockUpdateRequest,
	}

	return request
}
