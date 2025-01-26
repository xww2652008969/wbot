package client

import (
	"github.com/gorilla/websocket"
	"github.com/xww2652008969/wbot/client/Message"
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
	Func  Message.Event
	Event int
}
