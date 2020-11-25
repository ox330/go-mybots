package go_mybots

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

var (
	ViewMessage     []ViewMessageApi
	ViewNotice      []ViewOnNotice
	ViewRequest     []func(event Event)
	ViewMeta        []func(event Event)
	ViewOnCoCommand []ViewOnC0CommandApi
	Info            LoginInfo
)

type (
	ViewMessageApi struct {
		OnMessage   func(event Event)
		MessageType string
		SubType     string
	}
	ViewOnC0CommandApi struct {
		CoCommand func(event Event, args []string)
		Command   string
		Allies    string
		OnToMe    bool
	}
	ViewOnNotice struct {
		OnNotice   func(event Event)
		NoticeType string
		SubType    string
	}
)

func init() {

}

func EventMain(body io.Reader) {
	var event Event
	form, _ := ioutil.ReadAll(body)
	_ = json.Unmarshal(form, &event)
	viewsMessage(event)
}

func viewsMessage(event Event) {
	switch event.PostType {
	case "message":
		processMessageHandle(event)
	case "notice":
		processNoticeHandle(event)
	case "request":
		log.Printf("request_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d",
			event.RequestType, event.GroupId, event.UserId)
		for _, v := range ViewRequest {
			go v(event)
		}
	case "meta_event":
		log.Printf("post_type:%s\n\t\t\t\t\tmeta_event_type:%s\n\t\t\t\t\tinterval:%s",
			event.PostType, event.MetaEventType, event.Interval)
		for _, v := range ViewMeta {
			go v(event)
		}
	}
}

func processMessageHandle(event Event) {
	for _, v := range ViewOnCoCommand {
		onlyToMe := strings.Contains(event.Message, fmt.Sprintf("[CQ:at,qq=%d]", Info.UserId))
		content := strings.HasPrefix(event.Message, v.Command)
		allies := strings.HasPrefix(event.Message, v.Allies)
		if v.OnToMe == onlyToMe && (content || allies) {
			go v.CoCommand(event, strings.Fields(event.Message))
			log.Printf("message_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d\n\t\t\t\t\tmessage:%s",
				event.MessageType, event.GroupId, event.UserId, event.Message)
			return
		}
	}
	log.Printf("message_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d\n\t\t\t\t\tmessage:%s",
		event.MessageType, event.GroupId, event.UserId, event.Message)
	for _, v := range ViewMessage {
		if v.SubType == "" {
			v.SubType = event.SubType
		}
		if v.MessageType == "" {
			v.MessageType = event.NoticeType
		}
		if (v.MessageType == event.MessageType) && (v.SubType == event.SubType) {
			go v.OnMessage(event)
		}
	}
}

func processNoticeHandle(event Event) {
	log.Printf("notice_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d",
		event.NoticeType, event.GroupId, event.UserId)

	for _, v := range ViewNotice {
		if v.SubType == "" {
			v.SubType = event.SubType
		}
		if v.NoticeType == "" {
			v.NoticeType = event.NoticeType
		}
		if (v.NoticeType == event.NoticeType) && (v.SubType == event.SubType) {
			go v.OnNotice(event)
		}
	}
}

func OnlyToMe(event Event) bool {
	return strings.Contains(event.Message, fmt.Sprintf("[CQ:at,qq=%d]", Info.UserId))
}

func StartWith(event Event, s string) bool {
	return strings.HasPrefix(event.Message, s)
}

func EndWith(event Event, s string) bool {
	return strings.HasSuffix(event.Message, s)
}
