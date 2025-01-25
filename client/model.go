package client

import (
	"bot/client/Message"
	"github.com/gorilla/websocket"
	"net/http"
)

type Client struct {
	Config   Clientconfig
	Ws       *websocket.Conn
	Message  Message.Message
	EvebtFun []Eventfunc //
}
type Clientconfig struct {
	Wsurl      string
	Wspost     string
	Wsheader   http.Header
	Clienthttp string
}

type Clientevent struct {
	Eventtype int
	Message   Message.Message
}

type Eventfunc struct {
	Func  []Message.MessageEvent
	Event []int
}

const (
	Group   = 1
	Private = 2
)
