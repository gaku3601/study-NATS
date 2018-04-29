package main

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/nats-io/nats"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/", index)
	e.GET("/simplicity", simplicity)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))

}

// Handler
func index(c echo.Context) error {
	// Connect to server; defer close
	natsConnection, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Nats Server: "+err.Error())
	}
	defer natsConnection.Close()
	log.Println("Connected to " + nats.DefaultURL)

	// Publish message on subject
	subject := "foo"
	natsConnection.Publish(subject, []byte("Hello NATS"))

	// Simple Sync Subscriber
	sub, _ := natsConnection.SubscribeSync("resfoo")
	m, err := sub.NextMsg(10 * time.Millisecond)
	if err != nil {
		return c.String(http.StatusInternalServerError, "ServiceB Connection: "+err.Error())
	}
	return c.String(http.StatusOK, string(m.Data))
}

func simplicity(c echo.Context) error {
	nc, _ := nats.Connect(nats.DefaultURL)
	// Requests
	msg, _ := nc.Request("simplicity", []byte("help me"), 10*time.Millisecond)
	return c.String(http.StatusOK, string(msg.Data))
}
