package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/nats-io/nats"
)

func TestMain(m *testing.M) {
	// ここにテストの初期化処理
	code := m.Run()
	// ここでテストのお片づけ
	os.Exit(code)
}

func TestIndex(t *testing.T) {
	// 受け取り側mock
	nc, _ := nats.Connect(nats.DefaultURL)
	nc.Subscribe("index", func(m *nats.Msg) {
		nc.Publish(m.Reply, []byte(m.Data))
	})

	e := echo.New()
	req := new(http.Request)
	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	c.SetPath("/")

	index(c)
	if rec.Body.String() != "help me" {
		t.Errorf("%s", rec.Body.String())
	}
}

func TestLogin(t *testing.T) {
	// 受け取り側mock
	nc, _ := nats.Connect(nats.DefaultURL)
	nc.Subscribe("login", func(m *nats.Msg) {
		nc.Publish(m.Reply, []byte("rep"))
	})

	e := echo.New()
	req := new(http.Request)
	rec := httptest.NewRecorder()
	c := e.NewContext(standard.NewRequest(req, e.Logger()), standard.NewResponse(rec, e.Logger()))
	c.SetPath("/login")

	login(c)
	if rec.Body.String() != "rep" {
		t.Errorf("%s", rec.Body.String())
	}
}
