package view

type WsResponse struct {
	Type string `json:"type"`
	ID   uint64 `json:"id"`
	X    int64  `json:"x"`
	Y    int64  `json:"y"`
}
