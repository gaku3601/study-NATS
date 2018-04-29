package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/nats-io/nats"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/", index)
	e.GET("/login", login)

	// Start server
	e.Run(standard.New(":8080"))
}

func index(c echo.Context) error {
	nc, _ := nats.Connect(nats.DefaultURL)
	// Requests
	msg, _ := nc.Request("index", []byte("help me"), 10*time.Millisecond)
	return c.String(http.StatusOK, string(msg.Data))
}

func login(c echo.Context) error {
	// ログイン情報の作成
	loginUser := &user{Email: "pro.gaku@gmail.com", Password: "test"}
	nc, _ := nats.Connect(nats.DefaultURL)
	conn, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	// Requests
	var response string
	conn.Request("login", loginUser, &response, 10*time.Millisecond)
	return c.String(http.StatusOK, response)
}

type user struct {
	Email    string
	Password string
}
