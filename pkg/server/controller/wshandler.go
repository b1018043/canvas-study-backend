package controller

import (
	"encoding/json"
	"log"

	"github.com/b1018043/canvas-study-backend/pkg/server/view"
	"gopkg.in/olahol/melody.v1"
)

const (
	JOIN  string = "join"
	LEAVE string = "leave"
	MOVE  string = "move"
)

var id uint64 = 0

var clientid map[*melody.Session]uint64 = make(map[*melody.Session]uint64)

func ConnectHandlerGenerator(m *melody.Melody) func(s *melody.Session) {
	return func(s *melody.Session) {
		clientid[s] = id
		id++
		wr := &view.WsResponse{
			Type: JOIN,
			ID:   clientid[s],
		}
		if bytes, err := json.Marshal(wr); err != nil {
			// Nothing to do
		} else {
			m.BroadcastOthers(bytes, s)
		}
	}
}

func DisconnectHandlerGenerator(m *melody.Melody) func(s *melody.Session) {
	return func(s *melody.Session) {
		wr := &view.WsResponse{
			Type: LEAVE,
			ID:   clientid[s],
		}
		if bytes, err := json.Marshal(wr); err != nil {
			// Nothing to do
		} else {
			m.BroadcastOthers(bytes, s)
			delete(clientid, s)
		}
	}
}

func MessageHandlerGenerator(m *melody.Melody) func(s *melody.Session, b []byte) {
	return func(s *melody.Session, msg []byte) {
		var mvreq view.MoveRequest
		if err := json.Unmarshal(msg, &mvreq); err != nil {
			log.Fatalln("failed unmarshal json")
		}
		wr := &view.WsResponse{
			Type: MOVE,
			ID:   clientid[s],
			X:    mvreq.X,
			Y:    mvreq.Y,
		}
		if bytes, err := json.Marshal(wr); err != nil {
			log.Fatalln("failed marshal")
		} else {
			m.BroadcastOthers(bytes, s)
		}
	}
}
