package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type WsResponse struct {
	Type string `json:"type"`
	ID   uint64 `json:"id"`
	X    int64  `json:"x"`
	Y    int64  `json:"y"`
}

func main() {
	r := gin.Default()
	m := melody.New()

	var id uint64 = 0

	clientid := make(map[*melody.Session]uint64)

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(400, gin.H{
			"mes": "ok",
		})
	})

	r.GET("/ws", func(ctx *gin.Context) {
		m.HandleRequest(ctx.Writer, ctx.Request)
	})

	m.HandleConnect(func(s *melody.Session) {
		clientid[s] = id
		id++
		wr := &WsResponse{
			Type: "join",
			ID:   clientid[s],
		}
		if bytes, err := json.Marshal(wr); err != nil {
			// Nothing to do
		} else {
			m.BroadcastOthers(bytes, s)
		}
	})

	m.HandleDisconnect(func(s *melody.Session) {
		wr := &WsResponse{
			Type: "leave",
			ID:   clientid[s],
		}
		if bytes, err := json.Marshal(wr); err != nil {
			// Nothing to do
		} else {
			m.BroadcastOthers(bytes, s)
			delete(clientid, s)
		}
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.BroadcastOthers(msg, s)
	})

	r.Run(":8080")
}
