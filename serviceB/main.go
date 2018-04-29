package main

// Import Go and NATS packages
import (
	"fmt"
	"log"
	"runtime"

	"github.com/antonholmquist/jason"
	"github.com/nats-io/nats"
)

var (
	c *nats.EncodedConn
)

func main() {
	// Create server connection
	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ = nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()
	log.Println("Connected to " + nats.DefaultURL)

	c.Subscribe("index", index)
	c.Subscribe("login", login)
	runtime.Goexit()
}

func index(subj, reply string, msg string) {
	// Replying
	fmt.Println(msg)
	// 返答する
	c.Publish(reply, "I can help!")
}

func login(subj, reply string, msg string) {
	v, _ := jason.NewObjectFromBytes([]byte(msg))
	email, _ := v.GetString("Email")
	password, _ := v.GetString("Password")
	// 受け取った値をログ出力
	fmt.Println("Email:" + email)
	fmt.Println("Password:" + password)
	// 返答する
	c.Publish(reply, "login success!")
}
