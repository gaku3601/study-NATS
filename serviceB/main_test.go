package main

import (
	"os"
	"testing"
	"time"

	"github.com/nats-io/nats"
)

var (
	nc *nats.Conn
	c  *nats.EncodedConn
)

func TestMain(m *testing.M) {
	nc, _ = nats.Connect(nats.DefaultURL)
	c, _ = nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	errChan := make(chan error)
	// 起動
	index(nc, errChan)
	login(c, errChan)

	code := m.Run()
	// ここでテストのお片づけ
	os.Exit(code)
}

func TestIndex(t *testing.T) {
	// Requests
	msg, err := nc.Request("index", []byte("help me"), 10*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	result := string(msg.Data)
	exepect := "I can help!"
	if result != exepect {
		t.Fatalf("\n予測:%v\n結果:%v\n", exepect, result)
	}
}

func TestLogin(t *testing.T) {
	login := &user{Email: "pro.gaku@gmail.com", Password: "test"}
	// Requests
	var res string
	c.Request("login", login, &res, 10*time.Millisecond)
	exepect := "login success!"
	if res != exepect {
		t.Fatalf("\n予測:%v\n結果:%v\n", exepect, res)
	}
}

type user struct {
	Email    string
	Password string
}
