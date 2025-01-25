package client

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
)

func Create(config Clientconfig) (Client, error) {
	var client Client
	con, _, err := websocket.DefaultDialer.Dial(config.Wsurl+":"+config.Wspost, config.Wsheader)
	if err != nil {
		fmt.Print(err) //后面加log
		return client, errors.New("创建ws失败")
	}
	client.Ws = con
	client.Config = config
	return client, nil
}
func (c *Client) Run() {
	for {
		err := c.Ws.ReadJSON(&c.Message)
		if err != nil {
			fmt.Print(err)
			return
		}
		c.Message.Httpclient = c.Config.Clienthttp
		c.postevent()
	}
}

// postevent 内部函数 向框架推送event执行方法
func (c *Client) postevent() {
	switch c.Message.PostType {
	case "message":
		switch c.Message.MessageType {
		case "group":
			a := Clientevent{
				Eventtype: Group,
				Message:   c.Message,
			}
			c.sendevent(a)
		case "private":
			a := Clientevent{
				Eventtype: Private,
				Message:   c.Message,
			}
			c.sendevent(a)
		}

	}
}
func (c Client) Te() {
	fmt.Print("嘻嘻")
}
