package main

import (
	"github.com/gin-gonic/gin"
	"go_demo/pkg/httphandler"
)

func main() {

	r := gin.Default()
	r.GET("/healthy", func(c *gin.Context) {
		c.String(200, "Hello, Geektutu")
	})
	r.GET("/chat", httphandler.Chat)
	r.Run()
}

func initConfig() {

}
