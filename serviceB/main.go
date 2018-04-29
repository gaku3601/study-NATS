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
	conn, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer conn.Close()
	defer nc.Close()
	log.Println("Connected to " + nats.DefaultURL)

	errChan := make(chan error)
	// routing
	index(nc, errChan)
	login(conn, errChan)

	fmt.Println(<-errChan)
}

func index(nc *nats.Conn, errChan chan error) {
	// Replies
	_, err := nc.Subscribe("index", func(m *nats.Msg) {
		// Handle the message
		log.Printf("Received message '%s\n", string(m.Data)+"'")
		nc.Publish(m.Reply, []byte("I can help!"))
	})
	if err != nil {
		errChan <- err
	}
}

func login(conn *nats.EncodedConn, errChan chan error) {
	// Replying
	_, err := conn.Subscribe("login", func(subj, reply string, msg string) {
		v, _ := jason.NewObjectFromBytes([]byte(msg))
		email, _ := v.GetString("Email")
		password, _ := v.GetString("Password")
		// 受け取った値をログ出力
		fmt.Println("Email:" + email)
		fmt.Println("Password:" + password)
		// 返答する
		conn.Publish(reply, "login success!")
	})
	if err != nil {
		errChan <- err
	}
}
