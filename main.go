package main

import (
	"github.com/xww2652008969/wbot/client"
)

func main() {
	config := client.Clientconfig{
		Wsurl:      "ws://127.0.0.1",
		Wspost:     "3001",
		Wsheader:   nil,
		Clienthttp: "http://127.0.0.1:4000",
	}
	c, err := client.New(config)
	if err != nil {
		panic(err)
	}
	c.RegisterPush(te())
	c.Run()
}

func te() client.Push {
	return func(client client.Client) {
		api := client.Newsenapi()
		c := api.Getchatmessage()
		c.AddText("xixi")
		c.UserId = 1271701079
		api.SendPrivateMsg()
	}
}
