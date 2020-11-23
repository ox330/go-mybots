package events


import (
	"encoding/json"
	"github.com/3343780376/go-mybots/api"
	"github.com/3343780376/go-mybots/test"
	"io"
	"io/ioutil"
	"log"
)
var (
	ViewMessage []func(event api.Event)
	ViewNotice  []func(event api.Event)
	ViewRequest []func(event api.Event)
	ViewMeta    []func(event api.Event)
)


func init() {
	ViewMessage = append(ViewMessage, test.DefaultMessageHandle)
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
		log.Printf("message_type:%s\n\t\t\t\t\tgroup_id:%d\n\t\t\t\t\tuser_id:%d\n\t\t\t\t\tmessage:%s",
			event.MessageType,event.GroupId,event.UserId,event.Message)
		for _,v := range ViewMessage {
			go v(event)
		}
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