package test

import (
	"github.com/3343780376/go-mybots"
	"log"
)

var Bot = go_mybots.Bots{Address: "127.0.0.1", Port: 5700, Admin: 1743224847}

func init() {
	go_mybots.Info, _ = Bot.GetLoginInfo()
	go_mybots.ViewOnCoCommand = append(go_mybots.ViewOnCoCommand, go_mybots.ViewOnC0CommandApi{
		CoCommand: DefaultOnCoCommand, Command: "weather", Allies: "天气", OnToMe: false})
	go_mybots.ViewMessage = append(go_mybots.ViewMessage, go_mybots.ViewMessageApi{OnMessage: DefaultMessageHandle,
		MessageType: go_mybots.MessageTypeApi.Private, SubType: ""})
	go_mybots.ViewMessage = append(go_mybots.ViewMessage, go_mybots.ViewMessageApi{OnMessage: MessageTest,
		MessageType: go_mybots.MessageTypeApi.Group, SubType: ""})
	go_mybots.ViewNotice = append(go_mybots.ViewNotice, go_mybots.ViewOnNotice{OnNotice: DefaultNoticeHandle,
		NoticeType: go_mybots.NoticeTypeApi.GroupIncrease,
		SubType:    "approve"})
	go_mybots.ViewRequest = append(go_mybots.ViewRequest, DefaultRequestHandle)
	go_mybots.ViewMeta = append(go_mybots.ViewMeta, DefaultMetaHandle)
}

func DefaultMessageHandle(event go_mybots.Event) {
	log.Println("收到了私聊信息")
	go Bot.SendPrivateMsg(event.UserId, "hello   world", false)

}

func MessageTest(event go_mybots.Event) {
	log.Println("收到了")
	if event.GroupId == 972264701 {
		go Bot.SendGroupMsg(event.GroupId, event.Message, false)
	}
}

func DefaultNoticeHandle(event go_mybots.Event) {

}

func DefaultRequestHandle(event go_mybots.Event) {

}

func DefaultMetaHandle(event go_mybots.Event) {

}

func DefaultOnCoCommand(event go_mybots.Event, args []string) {

}
