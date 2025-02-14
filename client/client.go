package client

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
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
	client.messageChan = make(chan Client, 10)
	return client, nil
}
func (c *Client) Run() {
	go c.cron()
	go c.postevent()
	for {
		err := c.Ws.ReadJSON(&c.Message)
		if err != nil {
			continue
		}
		c.messageChan <- *c
	}
}

// cron 用于定时执行某些程序用于推送
func (c *Client) cron() {
	for _, v := range c.pushfunc {
		go v(*c)
	}
}

// postevent 内部函数 向框架推送event执行方法
func (c *Client) postevent() {
	for m := range c.messageChan {
		fmt.Println("开始处理")
		if m.Interceptor != nil {
			fmt.Println("开始处理")
			if !c.Interceptor(m, m.Message) {
				continue
			}
		}
		switch m.Message.PostType {
		case "message":
			switch m.Message.MessageType {
			case "group":
				a := Clientevent{
					Eventtype: MessageGroup,
					Message:   m.Message,
				}
				sendevent(m, a)
			case "private":
				a := Clientevent{
					Eventtype: MessagePrivate,
					Message:   m.Message,
				}
				sendevent(m, a)
			}
		case "notice":
			a := Clientevent{
				Eventtype: MessageNotice,
				Message:   m.Message,
			}
			sendevent(m, a)
		case "message_sent":
			a := Clientevent{
				Eventtype: MessageSent,
				Message:   m.Message,
			}
			sendevent(m, a)
		}
	}
}

func sendevent(client Client, clientevent Clientevent) {
	var wg sync.WaitGroup               // 创建 WaitGroup
	timeoutDuration := 10 * time.Second // 设置超时为 10 秒
	for _, v := range client.EvebtFun {
		if clientevent.Eventtype == v.Event {
			wg.Add(1)              // 增加 WaitGroup 计数
			go func(f Eventfunc) { // 使用闭包捕获 EventFunc
				defer wg.Done()                     // 确保在函数结束时调用 Done
				f.Func(client, clientevent.Message) // 调用事件处理函数
			}(v) // 将当前 v 传递给闭包
		}
	}
	done := make(chan struct{})
	go func() {
		wg.Wait()   // 等待所有 goroutine 完成
		close(done) // 关闭 done 通道
	}()

	select {
	case <-done:
	case <-time.After(timeoutDuration):
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
func (c *Client) RegisterMessageSenthandle(f Event) {
	var d Eventfunc
	d.Event = MessageSent
	d.Func = f
	c.EvebtFun = append(c.EvebtFun, d)
}

func (c *Client) RegisterPush(f Push) {
	c.pushfunc = append(c.pushfunc, f)
}
func (c *Client) RegisterInterceptor(f Interceptorfunc) {
	c.Interceptor = f
}
