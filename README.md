# Custom PubSub implementation example for Golang

I implemented a basic pubsub system example for practise in a day. In this example, I decided a architecture like this:

![Architecture](https://i.ibb.co/TtZ40LZ/Ads-z-2023-06-25-1733.png)

As you can see, there are 3 main components:

- **Publisher**: It is responsible for publishing messages to the topics.
- **Subscriber**: It is responsible for subscribing to the topics and receiving messages from the topics.
- **Broker**: It is responsible for managing topics and messages. It is also responsible for delivering messages to the subscribers.

In my scenario, I have: 2 publishers, 2 subscribers and 1 broker.

- Order Service **(producer)**: It publishes messages to the "order" topic when a new order is created.
- Shipping Service **(producer)**: It publishes messages to the "shipment" topic when status of shipment is updated.
- Shipping Service **(consumer)**: It subscribes to the "order" topic and receives messages from the "order" topic.
- Notification Service **(consumer)**: It subscribes to both "order" and "shipment" topic and receives messages from these topics.

## How to run?

1. Clone the repository:

```bash
git clone https://github.com/anilsenay/go-custom-pubsub.git
```

2. Run the broker:

```bash
go run ./pubsub/main.go
```

3. Run consumers in another tabs (-d flag is for delay in milliseconds):

```bash
go run ./consumers/notification/main.go -d 3000
```

```bash
go run ./consumers/shipping/main.go -d 3000
```

4. Run producers in another tabs (-d flag is for delay in milliseconds):

```bash
go run ./producers/order/main.go -d 5000
```

```bash
go run ./producers/shipping/main.go -d 5000
```

## Demo video:

https://github.com/anilsenay/go-custom-pubsub/assets/1047345/b0edc53a-877a-4823-b1f7-5b396d3877e9

## What more can be done?

- [ ] HTTP -> gRPC or another protocol
- [ ] Processing guarantees (at least once, at most once, exactly once)
- [ ] Message ordering guarantees
- [ ] Message persistence

## Contribution

Any contributions you make are greatly appreciated.
