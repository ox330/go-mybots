package main

import (
	"gononebot/urls"
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