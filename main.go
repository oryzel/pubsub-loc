package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	pubsubBoromir "github.com/oryzel/pubsub-loc/boromir"
	pubsubCeleborn "github.com/oryzel/pubsub-loc/celeborn"
	pubsubDefault "github.com/oryzel/pubsub-loc/default"
	pubsubGaladriel "github.com/oryzel/pubsub-loc/galadriel"
)

const (
	projectID      = "gpn-endor-local"
	topicID        = "boromir-sync-order"
	subscriptionID = "boromir-sync-order-boromir-process-sync-order"
)

func main() {
	ctx := context.Background()
	os.Setenv("PUBSUB_EMULATOR_HOST", "127.0.0.1:8085")

	// Initialize a client without credentials.
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	action(client, "galadriel-inbound-order-status-notification")
}

func action(pubSubClient *pubsub.Client, actionName string) {

	def := pubsubDefault.New(pubsubDefault.Opts{PubSubClient: *pubSubClient})
	boromir := pubsubBoromir.New(pubsubBoromir.Opts{PubSubClient: *pubSubClient})
	celeborn := pubsubCeleborn.New(pubsubCeleborn.Opts{PubSubClient: *pubSubClient})
	galadriel := pubsubGaladriel.New(pubsubGaladriel.Opts{PubSubClient: *pubSubClient})

	switch actionName {
	case "boromir-sync-order":
		boromir.SyncOrder(context.Background())
	case "celeborn-portal-outbound-header":
		celeborn.PortalOutboundOrderHeaderNotification(context.Background())
	case "galadriel-inbound-order-status-notification":
		galadriel.InboundOrderStatusNotification(context.Background())
	case "galadriel-sales-order-status-notification":
		galadriel.OrderStatusNotification(context.Background())
	case "galadriel-sku-stock-updated":
		galadriel.SKUSyncStock(context.Background())
	default:
		def.Publish(context.Background())

	}

}
