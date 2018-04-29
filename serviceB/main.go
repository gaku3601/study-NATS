package main

// Import Go and NATS packages
import (
	"fmt"
	"log"

	"github.com/nats-io/nats"
)

func main() {
	// Create server connection
	natsConnection, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connected to " + nats.DefaultURL)
	// Subscribe to subject
	log.Printf("Subscribing to subject 'foo'\n")
	// Channel Subscriber
	errChan := make(chan error)
	go func() {
		_, err := natsConnection.Subscribe("foo", func(msg *nats.Msg) {
			// Handle the message
			log.Printf("Received message '%s\n", string(msg.Data)+"'")

			//返信
			natsConnection.Publish("resfoo", []byte("response message!!!!!!!!!!!!!!"))
		})
		if err != nil {
			errChan <- err
		}
	}()
	// simplicity
	go func() {
		// Replies
		_, err := natsConnection.Subscribe("simplicity", func(m *nats.Msg) {
			// Handle the message
			log.Printf("Received message '%s\n", string(m.Data)+"'")
			natsConnection.Publish(m.Reply, []byte("I can help!"))
		})
		if err != nil {
			errChan <- err
		}
	}()
	fmt.Println(<-errChan)
}
