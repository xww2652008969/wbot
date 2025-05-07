package main

import (
	"github.com/xww2652008969/wbot/client"
)

func main() {
	config := client.Clientconfig{
		Wsurl:      "ws://192.168.10.209",
		Wspost:     "3001",
		Clienthttp: "http://192.168.10.209:4000",
	}
	c, _ := client.New(config)

	c.Run()
}
