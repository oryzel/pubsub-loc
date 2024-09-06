package _default

import "cloud.google.com/go/pubsub"

type Opts struct {
	PubSubClient pubsub.Client
}

type service struct {
	pubSubClient pubsub.Client
}

func New(o Opts) service {

	return service{
		pubSubClient: o.PubSubClient,
	}

}
