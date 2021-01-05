//event.go为对各种event事件进行处理
package go_mybots

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

type Reflect interface {
	GetNextEvent(int, int) Event
}

type RuleCheck struct {
}

//RuleCheckIn为RuleCheck结构体的实现
type RuleCheckIn interface {
	OnlyToMe(event Event) bool
	StartWith(event Event, s string) bool
	EndWith(event Event, s string) bool
}

var (
	ViewMessage     []ViewMessageApi       //ViewMessage为一个需要上报到各个Message事件的切片
	ViewNotice      []ViewOnNoticeApi      //ViewNotice为一个需要上报到各个Notice事件的切片
	ViewRequest     []func(event Event)    //ViewRequest为一个需要上报到各个Request事件的切片
	ViewMeta        []func(event Event)    //ViewMeta为一个需要上报到各个Meta事件的切片
	ViewOnCoCommand []ViewOnC0CommandApi   //ViewOnCommand为一个需要上报到各个Command事件的切片
	Info            LoginInfo              //Info为当前账号的信息
	c               = make(chan Event, 20) //c是一个event类型的通道
	nextEvent       bool                   //nextEvent是一个决定是否将消息下发的全局变量，当调用GetNextEvent时，该全局变量的值才会发生改变
)

//Rule结构体两个成员分别为RuleCheck类型的函数和需要传递的参数
type Rule struct {
	fun  func(event Event, a ...interface{}) bool
	args []interface{}
}

type (
	//ViewMessageApi结构体为一个Message消息上报位置，包含了具体的方法和需要传递的Message的MessageType和SubType
	ViewMessageApi struct {
		OnMessage   func(event Event)
		MessageType string
		SubType     string
	}
	//ViewOnCommandApi结构体为一个Command上报位置，包含了具体的方法和该命令的具体命令以及别名，RuleChecked为一个Rule切片，
	ViewOnC0CommandApi struct {
		CoCommand   func(event Event, args []string)
		Command     string
		Allies      string
		RuleChecked []Rule
	}
	//ViewOnNoticeApi结构体为一个Notice上报位置，包含了具体的方法和需要传递的Notice的NoticeType和SubType
	ViewOnNoticeApi struct {
		OnNotice   func(event Event)
		NoticeType string
		SubType    string
	}
)

//init函数中对全局变量nextEvent进行了初始化
func init() {
	nextEvent = false
}

//eventMain方法对由http传递过来的body进行json格式化，转化为Event事件
func eventMain(body io.Reader) {
	var event Event
	form, _ := ioutil.ReadAll(body)
	_ = json.Unmarshal(form, &event)
	viewsMessage(event)
}

/*GetNextEvent方法
int n: 需要等待多少个事件之后放弃获取该事件
int UserId: 需要获取的信息对应的qq账号

return: Message类型的Event
*/
func (bot Bots) GetNextEvent(n, UserId int) Event {
	defer func() {
		nextEvent = false
	}()

	nextEvent = true
	for i := 0; i < n; i++ {
		event := <-c
		if event.UserId != UserId {
			c <- event
			go processMessageHandle()
		} else {
			return event
		}
	}

	return Event{}
}

/*
	ViewsMessage方法根据PostType的不同，将不同事件交给对应的处理器处理
*/
func viewsMessage(event Event) {
	switch event.PostType {
	case "message":
		c <- event
		if !nextEvent {
			go processMessageHandle()
		}
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

//执行checkRule切片里所以类型为func(event Event,v []interface)bool 类型的方法
func checkRule(event Event, RuleChecked []Rule) bool {
	for _, rule := range RuleChecked {
		e := rule.fun(event, rule.args)
		if !e {
			return false
		}
	}
	return true
}

/*
	processMessageHandle方法从通道c里面取出event，
	判断是否符合ViewOnCommand切片里面的各个命令，如果
	符合，直接交给命令事件处理。否则，将event下发到各个
	Message事件处理器。
*/
func processMessageHandle() {
	event := <-c
	for _, v := range ViewOnCoCommand {
		content := strings.HasPrefix(event.Message, v.Command)
		allies := strings.HasPrefix(event.Message, v.Allies)
		if checkRule(event, v.RuleChecked) && (content || allies) {
			args := strings.Split(event.Message, " ")
			v.CoCommand(event, args)
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
			v.MessageType = event.MessageType
		}
		if (v.MessageType == event.MessageType) && (v.SubType == event.SubType) {
			go v.OnMessage(event)
		}
	}
}

/*
	processNoticeHandle方法将postType为Notice的event下发到各个处理器
*/
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

/*
	OnlyToMe方法为一个RuleChecked方法
	判断该事件如果在群聊中是否艾特机器人
	如果时私聊信息，直接返回True
*/
func (rule RuleCheck) OnlyToMe(event Event) bool {
	if event.MessageType == "group" {
		return strings.Contains(event.Message, fmt.Sprintf("[CQ:at,qq=%d]", Info.UserId))
	} else if event.MessageType == "private" {
		return true
	} else {
		return true
	}
}

/*
	StartWith方法判断该消息是否以某个字符串开头
	args:
		string s: 需要判断的字符串
*/
func (rule RuleCheck) StartWith(event Event, s string) bool {
	return strings.HasPrefix(event.Message, s)
}

/*
	EndWith方法判断该消息是否以某个字符串结尾
	args:
		string s: 需要判断的字符串
*/
func (rule RuleCheck) EndWith(event Event, s string) bool {
	return strings.HasSuffix(event.Message, s)
}
