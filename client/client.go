package client

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

func New(config Clientconfig) (Client, error) {
	var client Client
	con, _, err := websocket.DefaultDialer.Dial(config.Wsurl+":"+config.Wspost, config.Wsheader)
	if err != nil {
		fmt.Print(err) //后面加log
		return client, errors.New("创建ws失败")
	}
	client.Ws = con
	client.Config = config
	client.log = log.Default()
	return client, nil
}
func (c *Client) Run() {
	go c.cron()
	for {
		err := c.Ws.ReadJSON(&c.Message)
		if err != nil {
			continue
		}
		c.postevent()
	}
	select {}
}

// cron 用于定时执行某些程序用于推送
func (c *Client) cron() {
	for _, v := range c.pushfunc {
		go v(*c)
	}
}

// postevent 内部函数 向框架推送event执行方法
func (c *Client) postevent() {
	switch c.Message.PostType {
	case "message":
		switch c.Message.MessageType {
		case "group":
			a := Clientevent{
				Eventtype: MessageGroup,
				Message:   c.Message,
			}
			sendevent(*c, a)
		case "private":
			a := Clientevent{
				Eventtype: MessagePrivate,
				Message:   c.Message,
			}
			sendevent(*c, a)
		}
	case "notice":
		a := Clientevent{
			Eventtype: MessageNotice,
			Message:   c.Message,
		}
		sendevent(*c, a)
	}
}
func sendevent(client Client, clientevent Clientevent) {
	for _, v := range client.EvebtFun {
		if clientevent.Eventtype == v.Event {
			go v.Func(client, clientevent.Message)
		}
	}
}

// too  修改为传入event
func (c *Client) RegisterGroupHandle(f Event) {
	var d Eventfunc
	d.Event = MessageGroup
	d.Func = f
	c.EvebtFun = append(c.EvebtFun, d)
}
func (c *Client) RegisterPrivateHandle(f Event) {
	var d Eventfunc
	d.Event = MessagePrivate
	d.Func = f
	c.EvebtFun = append(c.EvebtFun, d)
}
func (c *Client) RegisterNoticeHandle(f Event) {
	var d Eventfunc
	d.Event = MessageNotice
	d.Func = f
	c.EvebtFun = append(c.EvebtFun, d)
}

func (c *Client) RegisterPush(f Push) {
	c.pushfunc = append(c.pushfunc, f)
}
