package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/anilsenay/go-custom-pubsub/producers/order/models"
	"github.com/anilsenay/go-custom-pubsub/producers/order/services"
)

var pubsub_url = flag.String("h", "http://localhost:8080/produce", "URL to send messages to")
var topic = flag.String("t", "ORDER", "Topic to send messages to")
var producer_count = flag.Int("c", 5, "Number of producers to spawn")
var delay = flag.Int("d", 1000, "Delay between messages in milliseconds")

func main() {
	flag.Parse()

	producers := make([]*services.OrderService, *producer_count)
	for i := 0; i < *producer_count; i++ {
		producers[i] = services.NewOrderService(*pubsub_url, *topic)
	}

	for {
		for _, producer := range producers {
			go func(c *services.OrderService) {
				randomOrderId := rand.Intn(100000)
				c.CreateOrder(models.Order{
					OrderID:    randomOrderId,
					CustomerID: 1234,
					ItemID:     1,
					Quantity:   1,
					Price:      1.0,
					Total:      1.0,
				})
			}(producer)
			time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)
		}
		time.Sleep(time.Duration(rand.Intn(*delay)) * time.Millisecond)
	}
}
