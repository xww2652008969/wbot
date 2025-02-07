package client

func (c *Client) logdebug(v ...interface{}) {
	c.log.Println(v)
}
