package main

import (
	"github/3343780376/go-mybots/urls"
	"log"
)

func init() {
}

func main() {
	rout := urls.Hand()
	err := rout.Run("127.0.0.1:8000")

	if err != nil {
		log.Fatal(err)
	}
}