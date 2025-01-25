package client

import "fmt"

func (c *Client) sendevent(clientevent Clientevent) {
	for _, v := range c.EvebtFun {
		for i := 0; i < len(v.Event); i++ {
			if clientevent.Eventtype == v.Event[i] {
				go v.Func[i](clientevent.Message)
				fmt.Print("执行了某函数")
			}
		}
	}
}
