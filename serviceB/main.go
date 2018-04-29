package main

// Import Go and NATS packages
import (
	"fmt"
	"log"

	"github.com/antonholmquist/jason"
	"github.com/nats-io/nats"
)

func main() {
	// Create server connection
	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()
	log.Println("Connected to " + nats.DefaultURL)

	// routing
	index(c)
	login(c)
}

func index(c *nats.EncodedConn) {
	// Replying
	c.Subscribe("index", func(subj, reply string, msg string) {
		fmt.Println(msg)
		// 返答する
		c.Publish(reply, "I can help!")
	})
}

func login(c *nats.EncodedConn) {
	// Replying
	c.Subscribe("login", func(subj, reply string, msg string) {
		v, _ := jason.NewObjectFromBytes([]byte(msg))
		email, _ := v.GetString("Email")
		password, _ := v.GetString("Password")
		// 受け取った値をログ出力
		fmt.Println("Email:" + email)
		fmt.Println("Password:" + password)
		// 返答する
		c.Publish(reply, "login success!")
	})
}
