package main

import (
	"log"

	"github.com/b1018043/canvas-study-backend/pkg/server"
)

func main() {
	if err := server.Server.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
