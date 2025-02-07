package main

import (
	"github.com/xww2652008969/wbot/client"
)

func main() {
	config := client.Clientconfig{
		Wsurl:      "ws://192.168.10.209",
		Wspost:     "3001",
		Wsheader:   nil,
		Clienthttp: "http://192.168.10.209:4000",
	}
	c, err := client.New(config)
	if err != nil {
		panic(err)
	}
	c.RegisterPush(te())
	c.Run()
	select {}

}
func te() client.Push {
	return func(client client.Client) {
		client.AddText("主动执行")
		client.SendPrivateMsg(1271701079)
	}
}
