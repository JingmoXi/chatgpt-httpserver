package httphandler

import "github.com/gin-gonic/gin"
import "go_demo/pkg/chatgpt-api"

func Chat(c *gin.Context) {
	msg := c.Query("msg")
	//msg, _ := c.Params.Get("msg")
	if msg == "" {
		c.String(400, "msg is empty")
	}

	resp, _ := chatgpt_api.Call(msg)
	c.String(200, resp)
}
