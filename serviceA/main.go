package main

import (
	"log"
	"time"

	"github.com/nats-io/nats"
)

func main() {
	// Connect to server; defer close
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	defer natsConnection.Close()
	log.Println("Connected to " + nats.DefaultURL)
	// Publish message on subject
	subject := "foo"
	natsConnection.Publish(subject, []byte("Hello NATS"))
	log.Println("Published message on subject " + subject)

	// Simple Sync Subscriber
	sub, _ := natsConnection.SubscribeSync("resfoo")
	m, _ := sub.NextMsg(10 * time.Millisecond)
	log.Printf("Received message '%s\n", string(m.Data)+"'")
}
