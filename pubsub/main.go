package main

import (
	"flag"

	"github.com/anilsenay/go-basic-pubsub/pubsub/handlers"
	"github.com/anilsenay/go-basic-pubsub/pubsub/pubsub"
)

var topics = []string{"ORDER", "SHIPMENT", "INVOICE"}

var buffer_size = flag.Int("b", 10000, "Buffer size for each channel")
var port = flag.String("p", "8080", "Port to listen on")

func main() {
	flag.Parse()

	pubsub_manager := pubsub.NewPubSubManager(*buffer_size)

	// create channels
	for _, topic := range topics {
		pubsub_manager.CreateTopic(topic)
	}

	// handler
	handler := handlers.NewPubSubHandler(*port, pubsub_manager)
	handler.Listen()
}
