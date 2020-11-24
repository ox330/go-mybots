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
	ViewMessage     []func(event Event)
	ViewNotice      []func(event Event)
	ViewRequest     []func(event Event)
	ViewMeta        []func(event Event)
	ViewOnCoCommand []ViewOnC0CommandApi
	Info            LoginInfo
)


type ViewOnC0CommandApi struct {
	CoCommand func(event Event,args []string)
	Content   string
	Allies    string
	OnlyToMe  bool
}

func init() {

}

func EventMain(body io.Reader)  {
	var event Event
	form, _ := ioutil.ReadAll(body)
	_ = json.Unmarshal(form, &event)
	viewsMessage(event)
}


func viewsMessage(event Event)  {
	switch event.PostType {
	case "message":
		processMessageHandle(event)
	case "notice":
		log.Printf("notice_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d",
			event.NoticeType,event.GroupId,event.UserId)
		for _,v := range ViewNotice {
			go v(event)
		}
	case "request" :
		log.Printf("request_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d",
			event.RequestType,event.GroupId,event.UserId)
		for _,v := range ViewRequest {
			go v(event)
		}
	case "meta_event":
		log.Printf("post_type:%s\n\t\t\t\t\tmeta_event_type:%s\n\t\t\t\t\tinterval:%s",
			event.PostType,event.MetaEventType,event.Interval)
		for _,v := range ViewMeta {
			go v(event)
		}
	}
}

func processMessageHandle(event Event)  {
	for _,v := range ViewOnCoCommand {
		onlyToMe := strings.Contains(event.Message,fmt.Sprintf("[CQ:at,qq=%d]", Info.UserId))
		content := strings.HasPrefix(event.Message,v.Content)
		allies := strings.HasPrefix(event.Message,v.Allies)
		log.Println(onlyToMe,content,allies)
		if onlyToMe == v.OnlyToMe && (content||allies){
			go v.CoCommand(event,strings.Fields(event.Message))
			log.Printf("message_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d\n\t\t\t\t\tmessage:%s",
				event.MessageType,event.GroupId,event.UserId,event.Message)
			return
		}
	}
	log.Printf("message_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d\n\t\t\t\t\tmessage:%s",
		event.MessageType,event.GroupId,event.UserId,event.Message)
	for _,v := range ViewMessage {
		go v(event)
	}
}