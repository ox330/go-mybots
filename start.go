package go_mybots

import (
	"github.com/3343780376/go-mybots/urls"
	"log"
)

func Run(address string)  {
	hand := urls.Hand()
	err := hand.Run(address)
	if err != nil {
		log.Fatalf("%s is err",address)
	}
}
