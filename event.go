package go_mybots

import (
	"encoding/json"
	"fmt"
	"github.com/3343780376/go-mybots/api"
	"github.com/3343780376/go-mybots/test"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

var (
	ViewMessage     []func(event api.Event)
	ViewNotice      []func(event api.Event)
	ViewRequest     []func(event api.Event)
	ViewMeta        []func(event api.Event)
	ViewOnCoCommand []ViewOnC0CommandApi
	info            api.LoginInfo
)


type ViewOnC0CommandApi struct {
	CoCommand func(event api.Event,args []string)
	content   string
	Allies    string
	OnlyToMe  bool
}

func init() {
	info = test.Bot.GetLoginInfo()
	ViewOnCoCommand = append(ViewOnCoCommand, ViewOnC0CommandApi{test.DefaultOnCoCommand, "weather","天气", true})
	ViewMessage = append(ViewMessage, test.DefaultMessageHandle)
	ViewMessage = append(ViewMessage, test.MessageTest)
	ViewNotice = append(ViewNotice,test.DefaultNoticeHandle)
	ViewRequest = append(ViewRequest,test.DefaultRequestHandle)
	ViewMeta = append(ViewMeta,test.DefaultMetaHandle)
}

func EventMain(body io.Reader)  {
	var event api.Event
	form, _ := ioutil.ReadAll(body)
	_ = json.Unmarshal(form, &event)
	viewsMessage(event)
}


func viewsMessage(event api.Event)  {
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

func processMessageHandle(event api.Event)  {
	for _,v := range ViewOnCoCommand {
		onlyToMe := strings.Contains(event.Message.Message,fmt.Sprintf("[CQ:at,qq=%d]", info.UserId))
		content := strings.HasPrefix(event.Message.Message,v.content)
		allies := strings.HasPrefix(event.Message.Message,v.Allies)
		log.Println(onlyToMe,content,allies)
		if onlyToMe == v.OnlyToMe && (content||allies){
			go v.CoCommand(event,strings.Fields(event.Message.Message))
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