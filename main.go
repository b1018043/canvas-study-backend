package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func main() {
	r := gin.Default()
	m := melody.New()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(400, gin.H{
			"mes": "ok",
		})
	})

	r.GET("/ws", func(ctx *gin.Context) {
		m.HandleRequest(ctx.Writer, ctx.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

	r.Run(":8080")
}
