package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/anilsenay/go-basic-pubsub/producers/shipping/models"
	"github.com/anilsenay/go-basic-pubsub/producers/shipping/services"
)

var pubsub_url = flag.String("h", "http://localhost:8080/produce", "URL to send messages to")
var topic = flag.String("t", "SHIPMENT", "Topic to send messages to")
var producer_count = flag.Int("c", 5, "Number of producers to spawn")
var delay = flag.Int("d", 1000, "Delay between messages in milliseconds")

func main() {
	flag.Parse()

	producers := make([]*services.ShippingService, *producer_count)
	for i := 0; i < *producer_count; i++ {
		producers[i] = services.NewShippingService(*pubsub_url, *topic)
	}

	for {
		for _, producer := range producers {
			go func(c *services.ShippingService) {
				c.UpdateStatus(models.ShippingStatus{
					OrderID:    rand.Intn(100000),
					CustomerID: 1234,
					Status:     randomShippingStatus(),
				})
			}(producer)
			time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		}
		time.Sleep(time.Duration(rand.Intn(*delay)) * time.Millisecond)
	}
}

func randomShippingStatus() string {
	statuses := []string{"SHIPPED", "DELIVERED", "RETURNED"}
	return statuses[rand.Intn(len(statuses))]
}
