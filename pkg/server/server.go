package server

import (
	"github.com/b1018043/canvas-study-backend/pkg/server/controller"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

const (
	JOIN  string = "join"
	LEAVE string = "leave"
	MOVE  string = "move"
)

var Server *gin.Engine

var m *melody.Melody

func init() {
	Server = gin.Default()

	m = melody.New()

	Server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(400, gin.H{
			"mes": "ok",
		})
	})

	Server.GET("/ws", func(ctx *gin.Context) {
		m.HandleRequest(ctx.Writer, ctx.Request)
	})

	m.HandleConnect(controller.ConnectHandlerGenerator(m))

	m.HandleDisconnect(controller.DisconnectHandlerGenerator(m))

	m.HandleMessage(controller.MessageHandlerGenerator(m))

}
