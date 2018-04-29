package main

import (
	"os"
	"testing"
	"time"

	"github.com/nats-io/nats"
)

var (
	c *nats.EncodedConn
)

func TestMain(m *testing.M) {
	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ = nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()
	// 起動
	index(c)
	login(c)

	code := m.Run()
	// ここでテストのお片づけ
	os.Exit(code)
}

func TestIndex(t *testing.T) {
	// Requests
	var res string
	c.Request("index", []byte("help me"), &res, 10*time.Millisecond)
	exepect := "I can help!"
	if res != exepect {
		t.Fatalf("\n予測:%v\n結果:%v\n", exepect, res)
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
