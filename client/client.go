package client

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/xww2652008969/wbot/MessageType"
	"github.com/xww2652008969/wbot/client/Message"
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
			continue
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
				Eventtype: MessageType.Group,
				Message:   c.Message,
			}
			c.sendevent(a)
		case "private":
			a := Clientevent{
				Eventtype: MessageType.Private,
				Message:   c.Message,
			}
			c.sendevent(a)
		}
	case "notice":
		a := Clientevent{
			Eventtype: MessageType.Notice,
			Message:   c.Message,
		}
		c.sendevent(a)
	}
}
func (c *Client) sendevent(clientevent Clientevent) {
	for _, v := range c.EvebtFun {
		if clientevent.Eventtype == v.Event {
			go v.Func(clientevent.Message)
		}
	}
}
func (c *Client) Register(event int, f Message.Event) {
	var d Eventfunc
	d.Event = event
	d.Func = f
	c.EvebtFun = append(c.EvebtFun, d)
}
