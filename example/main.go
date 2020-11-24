package main

import (
	"github.com/3343780376/go-mybots"
	_ "github.com/3343780376/go-mybots/example/test"
	"log"
)

func main() {
	hand := go_mybots.Hand()
	err := hand.Run("127.0.0.1:8000")
	if err != nil {
		log.Println(err.Error())
	}
}