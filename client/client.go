package client

import (
	"github.com/gorilla/websocket"
	"github.com/xww2652008969/wbot/client/chatmessage"
	"github.com/xww2652008969/wbot/client/sendapi"
	"net/http"
	"sync"
	"time"
)

// New 构造函数
func New(config Clientconfig) (*Client, error) {
	// 设置默认值
	if config.MaxReconnect == 0 {
		config.MaxReconnect = 5
	}
	if config.ReconnectInterval == 0 {
		config.ReconnectInterval = 1 * time.Second
	}
	if config.MaxReconnectInterval == 0 {
		config.MaxReconnectInterval = 30 * time.Second
	}

	client := &Client{
		Config:      config,
		messageChan: make(chan Client, 10),
		stopChan:    make(chan struct{}),
	}
	if config.Wstoken != "" {
		client.Config.wshead = make(http.Header)
		client.Config.wshead["Authorization"] = []string{"Bearer " + config.Wstoken}
	}
	// 初始化连接
	if err := client.connect(); err != nil {
		return nil, err
	}
	return client, nil
}

// connect 建立连接
func (c *Client) connect() error {
	c.reconnectMutex.Lock()
	defer c.reconnectMutex.Unlock()

	// 关闭旧连接
	if c.Ws != nil {
		c.Ws.Close()
	}
	// 建立新连接
	conn, _, err := websocket.DefaultDialer.Dial(
		c.Config.Wsurl+":"+c.Config.Wspost,
		c.Config.wshead,
	)
	if err != nil {
		return err
	}

	c.Ws = conn
	c.currentAttempts = 0 // 重置重连计数器
	return nil
}
func (c *Client) Run() {
	go c.cron()
	go c.postevent()
	for {
		select {
		case <-c.stopChan:
			return
		default:
			err := c.Ws.ReadJSON(&c.Message)
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					c.handleError(err)
				}
				continue
			}
			c.messageChan <- *c
		}
	}
}

// handleError 错误处理
func (c *Client) handleError(err error) {
	c.reconnectMutex.Lock()
	defer c.reconnectMutex.Unlock()
	c.currentAttempts++
	if c.currentAttempts > c.Config.MaxReconnect {
		panic("超过最大重连次数")
	}

	// 计算指数退避重连间隔
	delay := c.Config.ReconnectInterval * time.Duration(1<<uint(c.currentAttempts))
	if delay > c.Config.MaxReconnectInterval {
		delay = c.Config.MaxReconnectInterval
	}

	time.AfterFunc(delay, func() {
		if err := c.connect(); err == nil {
			c.currentAttempts = 0 // 重置计数器
		}
	})
}

// cron 用于定时执行某些程序用于推送
func (c *Client) cron() {
	if len(c.Pluginslist) < 1 {
		return
	}
	for _, v := range c.Pluginslist {
		go v.Push(c) // 传递指向 Client 的指针，而不是复制
	}
}

// postevent 内部函数 向框架推送event执行方法
func (c *Client) postevent() {
	for m := range c.messageChan {
		switch m.Message.PostType {
		case "message":
			switch m.Message.MessageType {
			case "group":
				sendevent(&m, MessageGroup)
			case "private":
				sendevent(&m, MessagePrivate)
			}
		case "notice":
			sendevent(&m, MessageNotice)
		case "message_sent":
			sendevent(&m, MessageSent)
		}
	}
}
func (c *Client) Close() {
	close(c.stopChan)
	c.Ws.Close()
}

func sendevent(client *Client, Event int) {
	var wg sync.WaitGroup                // 创建 WaitGroup
	timeoutDuration := 100 * time.Second // 设置超时为 100 秒
	if len(client.Pluginslist) < 1 {
		return
	}
	for _, v := range client.Pluginslist {
		wg.Add(1)
		go func(plugins Plugin) {
			switch Event {
			case MessageGroup:
				plugins.GroupHandle(*client, client.Message)
			case MessagePrivate:
				plugins.PrivateHandle(*client, client.Message)
			case MessageNotice:
				plugins.NoticeHandle(*client, client.Message)
			case MessageSent:
				plugins.MessageSendhandle(*client, client.Message)
			}
			defer wg.Done()
		}(v)
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

func (c *Client) AddPlugin(Ppugin Plugin) {
	c.Pluginslist = append(c.Pluginslist, Ppugin)
}
func (c *Client) Newsenapi() *sendapi.SendAPI {
	a := sendapi.SendAPI{}
	a.SetHttpUrl(c.Config.Clienthttp)
	a.SetHttpClient(http.DefaultClient)
	if c.Config.Clienthttptoken != "" {
		a.SetHttpAuthorization("Bearer " + c.Config.Clienthttptoken)
	}

	a.SetChatmessage(&chatmessage.ChatMessage{
		Group_id: 0,
		UserId:   0,
		Message:  make([]chatmessage.ChatMessageData, 0),
	})
	return &a
}
