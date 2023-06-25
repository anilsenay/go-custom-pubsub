package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anilsenay/go-basic-pubsub/pubsub/models"
)

type PubSub interface {
	Enqueue(topic, msg string) error
	Dequeue(topic, subscriber string) (string, error)
	Subscribe(topic, subscriber string) error
}

type PubSubHandler struct {
	port        string
	httpHandler *http.Server
	pubSub      PubSub
}

func NewPubSubHandler(port string, pubsub PubSub) *PubSubHandler {
	mux := http.NewServeMux()
	h := PubSubHandler{
		port:        port,
		httpHandler: &http.Server{Addr: ":" + port, Handler: mux},
		pubSub:      pubsub,
	}
	mux.HandleFunc("/produce", h.HandleProduce())
	mux.HandleFunc("/consume", h.HandleConsume())
	mux.HandleFunc("/subscribe", h.HandleSubscribe())
	return &h
}

func (h *PubSubHandler) HandleSubscribe() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		topic := r.URL.Query().Get("topic")
		subscriber := r.URL.Query().Get("subscriber")
		err := h.pubSub.Subscribe(topic, subscriber)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Subscribed to %s", topic)))
	}
}

func (h *PubSubHandler) HandleProduce() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := models.Message{}
		err := json.NewDecoder(r.Body).Decode(&msg)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		b, err := json.Marshal(msg.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		h.pubSub.Enqueue(msg.Topic, string(b))
		w.WriteHeader(http.StatusOK)
	}
}

func (h *PubSubHandler) HandleConsume() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		topic := r.URL.Query().Get("topic")
		subscriber := r.URL.Query().Get("subscriber")
		message, err := h.pubSub.Dequeue(topic, subscriber)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	}
}

func (h *PubSubHandler) Listen() {
	fmt.Printf("Listening on port %s\n", h.port)
	err := h.httpHandler.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func (h *PubSubHandler) Close() {
	h.httpHandler.Shutdown(context.Background())
}
