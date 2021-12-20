package main

import (
	"log"

	"example.com/rcs/pkg/server"
)

func main() {
	if err := server.Server.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
