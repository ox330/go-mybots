package test

import (
	"github.com/3343780376/go-mybots/api"
)
var Bot = api.Bots{Address: "127.0.0.1", Port: 5700,Admin: 1743224847}
func DefaultMessageHandle(event api.Event)  {
	if event.Message.Message=="hello"&&event.UserId==Bot.Admin {
		go Bot.DeleteMsg(event.MessageId)
	}
}

func MessageTest(event api.Event)  {

}

func DefaultNoticeHandle(event api.Event)  {

}

func DefaultRequestHandle(event api.Event)  {

}

func DefaultMetaHandle(event api.Event)  {

}

func DefaultOnCoCommand(event api.Event, args []string)  {

}
