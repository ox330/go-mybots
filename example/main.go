package main

import (
	"github.com/3343780376/go-mybots"
	_ "github.com/3343780376/go-mybots/example/test"
	"log"
)

func main() {
	hand := go_mybots.Hand()
	go_mybots.LoadFilter("E:\\projects\\gononebot\\go-mybots\\example\\config.json")
	err := hand.Run("127.0.0.1:8000")
	if err != nil {
		log.Println(err.Error())
	}
}
