package server

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type WsResponse struct {
	Type string `json:"type"`
	ID   uint64 `json:"id"`
	X    int64  `json:"x"`
	Y    int64  `json:"y"`
}

type MoveRequest struct {
	X int64 `json:"x"`
	Y int64 `json:"y"`
}

const (
	JOIN  string = "join"
	LEAVE string = "leave"
	MOVE  string = "move"
)

var Server *gin.Engine

var Melody *melody.Melody

func init() {
	Server = gin.Default()

	Melody = melody.New()

	var id uint64 = 0

	clientid := make(map[*melody.Session]uint64)

	Server.GET("/", func(ctx *gin.Context) {
		ctx.JSON(400, gin.H{
			"mes": "ok",
		})
	})

	Server.GET("/ws", func(ctx *gin.Context) {
		Melody.HandleRequest(ctx.Writer, ctx.Request)
	})

	Melody.HandleConnect(func(s *melody.Session) {
		clientid[s] = id
		id++
		wr := &WsResponse{
			Type: JOIN,
			ID:   clientid[s],
		}
		if bytes, err := json.Marshal(wr); err != nil {
			// Nothing to do
		} else {
			Melody.BroadcastOthers(bytes, s)
		}
	})

	Melody.HandleMessage(func(s *melody.Session, b []byte) {
		Melody.BroadcastOthers(b, s)
	})

	Melody.HandleDisconnect(func(s *melody.Session) {
		wr := &WsResponse{
			Type: LEAVE,
			ID:   clientid[s],
		}
		if bytes, err := json.Marshal(wr); err != nil {
			// Nothing to do
		} else {
			Melody.BroadcastOthers(bytes, s)
			delete(clientid, s)
		}
	})

	Melody.HandleMessage(func(s *melody.Session, msg []byte) {
		var mvreq MoveRequest
		if err := json.Unmarshal(msg, &mvreq); err != nil {
			log.Fatalln("failed unmarshal json")
		}
		wr := &WsResponse{
			Type: MOVE,
			ID:   clientid[s],
			X:    mvreq.X,
			Y:    mvreq.Y,
		}
		if bytes, err := json.Marshal(wr); err != nil {
			log.Fatalln("failed marshal")
		} else {
			Melody.BroadcastOthers(bytes, s)
		}
	})

}
