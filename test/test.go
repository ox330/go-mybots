package test

import (
	"github/3343780376/go-mybots/api"
)
var bot = api.Bots{Address: "127.0.0.1", Port: 5700,Admin: 1743224847}
func DefaultMessageHandle(event api.Event)  {
	if event.Message=="hello"&&event.UserId==bot.Admin {
		go bot.SendPrivateMsg(bot.Admin,"hello,world",true)
	}
}

func DefaultNoticeHandle(event api.Event)  {

}

func DefaultRequestHandle(event api.Event)  {

}

func DefaultMetaHandle(event api.Event)  {

}
