package client

import (
	"github.com/gorilla/websocket"
	"github.com/xww2652008969/wbot/client/chatmessage"
	"net/http"
	"sync"
	"time"
)

const (
	MessageGroup   = 1
	MessagePrivate = 2
	MessageNotice  = 3
	MessageSent    = 4
)

type Client struct {
	Config          Clientconfig
	Ws              *websocket.Conn
	messageChan     chan Client
	Message         Message  //接受到的消息
	Pluginslist     []Plugin //事件总线
	reconnectMutex  sync.Mutex
	currentAttempts int
	stopChan        chan struct{}
}
type Clientconfig struct {
	Wsurl                string
	Wspost               string
	Wstoken              string
	wshead               http.Header
	Clienthttp           string
	Clienthttptoken      string
	MaxReconnect         int           // 最大重连次数，默认5次
	ReconnectInterval    time.Duration // 初始重连间隔，默认1秒
	MaxReconnectInterval time.Duration // 最大重连间隔，默认30秒
}

// v2
type Plugin interface {
	PluginName() string
	PluginVersion() string
	PluginAuthor() string
	GroupHandle(client Client, message Message)
	PrivateHandle(client Client, message Message)
	MessageSendhandle(client Client, message Message)
	NoticeHandle(client Client, message Message)
	Push(client *Client)
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
	Message       []chatmessage.ChatMessageData `json:"message"`
	RawMessage    string                        `json:"raw_message"`
	Font          int                           `json:"font"`
	SubType       string                        `json:"sub_type"`
	MessageFormat string                        `json:"message_format"`
	PostType      string                        `json:"post_type"`
	GroupId       int64                         `json:"group_id"`
	NoticeType    string                        `json:"notice_type,omitempty"`
	TargetId      int64                         `json:"target_id,omitempty"`
	SenderId      int                           `json:"sender_id,omitempty"`
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
