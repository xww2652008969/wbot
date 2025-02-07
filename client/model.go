package client

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

const (
	MessageGroup   = 1
	MessagePrivate = 2
	MessageNotice  = 3
)

type Client struct {
	Config   Clientconfig
	Ws       *websocket.Conn
	Message  Message     //接受到的消息
	EvebtFun []Eventfunc //事件总线
	pushfunc []Push
	log      *log.Logger //log
	updata   []upd       //处理发送消息的
}
type Clientconfig struct {
	Wsurl      string
	Wspost     string
	Wsheader   http.Header
	Clienthttp string
}

type Clientevent struct {
	Eventtype int
	Message   Message
}

type Message struct {
	SelfId      int64  `json:"self_id"`
	UserId      int64  `json:"user_id"`
	Time        int64  `json:"time"`
	MessageId   int64  `json:"message_id"`
	MessageSeq  int    `json:"message_seq"`
	RealId      int    `json:"real_id"`
	MessageType string `json:"message_type"`
	Sender      struct {
		UserId   int    `json:"user_id"`
		Nickname string `json:"nickname"`
		Card     string `json:"card"`
		Role     string `json:"role"`
	} `json:"sender"`
	RawMessage    string `json:"raw_message"`
	Font          int    `json:"font"`
	SubType       string `json:"sub_type"`
	Message       []upd  `json:"message"`
	MessageFormat string `json:"message_format"`
	PostType      string `json:"post_type"`
	GroupId       int64  `json:"group_id"`
	NoticeType    string `json:"notice_type,omitempty"`
	TargetId      int64  `json:"target_id,omitempty"`
	SenderId      int    `json:"sender_id,omitempty"`
	RawInfo       []struct {
		Col  string `json:"col,omitempty"`
		Nm   string `json:"nm,omitempty"`
		Type string `json:"type"`
		Uid  string `json:"uid,omitempty"`
		Jp   string `json:"jp,omitempty"`
		Src  string `json:"src,omitempty"`
		Txt  string `json:"txt,omitempty"`
		Tp   string `json:"tp,omitempty"`
	} `json:"raw_info,omitempty"`
}
type Eventfunc struct {
	Func  Event
	Event int
}

type Event func(client Client, message Message)
type Push func(client Client)
