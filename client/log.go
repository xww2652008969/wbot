package client

func (c *Client) logdebug(v ...interface{}) {
	c.Log.Println(v)
}
