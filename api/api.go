package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type MessageIds struct {
	MessageId int32 `json:"message_id"`
}
type Bots struct {
	Address string
	Port    int
	Admin   int
}
type LoginInfo struct {
	UserId int `json:"user_id"`
	NickName string `json:"nick_name"`
}


type anonymous struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Flag string `json:"flag"`
}

type Files struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Size  int64  `json:"size"`
	Busid int64  `json:"busid"`
}

type Event struct {
	Anonymous     anonymous `json:"anonymous"`
	Font          string    `json:"font"`
	GroupId       int       `json:"group_id"`
	Message       string    `json:"message"`
	MessageType   string    `json:"message_type"`
	PostType      string    `json:"post_type"`
	RawMessage    string    `json:"raw_message"`
	SelfId        int       `json:"self_id"`
	Sender        Senders   `json:"sender"`
	SubType       string    `json:"sub_type"`
	UserId        int       `json:"user_id"`
	Time          int       `json:"time"`
	NoticeType    string    `json:"notice_type"`
	RequestType   string    `json:"request_type"`
	Comment       string    `json:"comment"`
	Flag          string    `json:"flag"`
	OperatorId    int       `json:"operator_id"`
	File          Files     `json:"file"`
	Duration      int64     `json:"duration"`
	TargetId      int64     `json:"target_id"`
	HonorType     string    `json:"honor_type"`
	MetaEventType string    `json:"meta_event_type"`
	Status        string    `json:"status"`
	Interval      string    `json:"interval"`
	MessageIds
}

type Senders struct {
	Age      string `json:"age"`
	Area     string `json:"area"`
	Card     string `json:"card"`
	Level    string `json:"level"`
	NickName string `json:"nickname"`
	Role     string `json:"role"`
	Sex      string `json:"sex"`
	Title    string `json:"title"`
	UserId   int    `json:"user_id"`
}

type GroupMemberInfo struct {
	GroupId int `json:"group_id"`
	JoinTime int `json:"join_time"`
	LastSentTime int `json:"last_sent_time"`
	Unfriendly bool `json:"unfriendly"`
	TitleExpireTime int `json:"title_expire_time"`
	CardChangeable bool `json:"card_changeable"`
	Senders
}

type responseJson struct {
	Status  string `json:"status"`
	RetCode int    `json:"retcode"`
}
type responseLoginIndoJson struct {
	responseJson
	Data LoginInfo `json:"data"`
}

type responseMessageJson struct {
	responseJson
	Data MessageIds `json:"data"`
}

type getMsgJson struct {
	responseJson
	Data GetMessage `json:"data"`
}

type GetMessage struct {
	Time        int32        `json:"time"`
	Group 		bool     	 `json:"group"`
	MessageId   int32        `json:"message_id"`
	RealId      int32        `json:"real_id"`
	Sender      Senders 	 `json:"sender"`
	Message     string		 `json:"message"`
}

type Send interface {
	SendGroupMsg(groupId int, message string,autoEscape bool) int32
	SendPrivateMsg(userId int,message string,autoEscape bool) int32
	DeleteMsg(messageId int32)
	GetMsg(messageId int32) GetMessage
	SetGroupBan(groupId int32,userId int32,duration int)
	SetGroupCard(groupId int32,userId int32,card string)
}
type FriendList struct {
	UserId int `json:"user_id"`
	NickName string `json:"nick_name"`
	Remark string `json:"remark"`
}
type responseFriendListJson struct {
	responseJson
	Data []FriendList `json:"data"`
}


type GroupInfo struct {
	GroupId int `json:"group_id"`
	GroupName string `json:"group_name"`
	MemberCount int `json:"member_count"`
	MaxMemberCount int `json:"max_member_count"`
}
type responseGroupInfoJson struct {
	responseJson
	Data GroupInfo `json:"data"`
}

type GroupHonorInfo struct {
	GroupId int `json:"group_id"`
	CurrentTalkative CurrentTalkativeS `json:"current_talkative"`
	TalkativeList []GroupHonorInfoList `json:"talkative_list"`
	PerformerList []GroupHonorInfoList `json:"performer_list"`
	LegendList    []GroupHonorInfoList `json:"legend_list"`
	StrongNewbieList []GroupHonorInfoList `json:"strong_newbie_list"`
	EmotionList  []GroupHonorInfoList `json:"emotion_list"`
}
type CurrentTalkativeS struct {
	UserId int `json:"user_id"`
	NickName string `json:"nick_name"`
	Avatar string `json:"avatar"`
	DayCount int `json:"day_count"`
}
type GroupHonorInfoList struct {
	UserId int `json:"user_id"`
	NickName string `json:"nick_name"`
	Avatar string `json:"avatar"`
	Description string `json:"description"`
}
type Record struct {
	File string	`json:"file"`
	OutFormat string `json:"out_format"`
}
type Cookie struct {
	Cookies string `json:"cookies"`
}

type CsrfToken struct {
	Token string `json:"token"`
}

type Credentials struct {
	Cookie
	CsrfToken
}
type Image struct {
	File string `json:"file"`
}

type Bool struct {
	Yes bool `json:"yes"`
}

type OnlineStatus struct {
	Online bool `json:"online"`
	Good bool `json:"good"`
}
type defaultResponseJson struct {
	responseJson
	Data string `json:"data"`
}

type responseStrangerInfoJson struct {
	responseJson
	Data Senders `json:"data"`
}

type Api interface {
	SendMsg(messageType string,id int,message string,autoEscape bool)int32
	SendLike(userId int,times int)
	SetGroupKick(groupId int,UserId int,rejectAddRequest bool)
	SetGroupAnonymousBan(groupId int,flag string,duration int)
	SetGroupWholeBan(groupId int,enable bool)
	SetGroupAdmin(groupId int,UserId int,enable bool)
	SetGroupAnonymous(groupId int,enable bool)
	SetGroupName(groupId int,groupName string)
	SetGroupLeave(groupId int,isDisMiss bool)
	SetGroupSpecialTitle(groupId int,userId int,specialTitle string,duration int)
	SetFriendAddRequest(flag string,approve bool,remark string)
	SetGroupAddRequest(flag string,subType string,approve bool,reason string)
	GetLoginInfo() LoginInfo
	GetStrangerInfo()Senders
	GetFriendList()[]FriendList
	GetGroupInfo(groupId int,noCache bool)GroupInfo
	GetGroupList()[]GroupInfo
	GetGroupMemberInfo(groupId int,UserId int,noCache bool)GroupMemberInfo
	GetGroupMemberList(groupId int)[]GroupMemberInfo
	GetGroupHonorInfo(groupId int,honorType string)GroupHonorInfo
	GetCookies(domain string)Cookie
	GetCsrfToken()CsrfToken
	GetCredentials(domain string)Credentials
	GetRecord()Record
	GetImage(file string)
	CanSendImage()Bool
	CanSendRecord()Bool
	GetStatus()OnlineStatus
}

var (
	responseMsgJson responseMessageJson
	getMessageJson getMsgJson
	defaultJson defaultResponseJson
	LoginInfoJson responseLoginIndoJson
	StrangerInfo responseStrangerInfoJson
	FriendListJson responseFriendListJson
	GroupInfoJson responseGroupInfoJson
)

func (bot Bots) SendGroupMsg(groupId int,message string,autoEscape bool) int32  {
	requestUrl := "http://%s:%d/send_group_msg?"
	requestUrl += "group_id=%d"
	requestUrl += "&message=%s"
	requestUrl += "&auto_escape=%s"
	url := fmt.Sprintf(requestUrl,bot.Address, bot.Port, groupId,message,strconv.FormatBool(autoEscape))
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &responseMsgJson)
	return responseMsgJson.Data.MessageId
}



func (bot Bots) SendPrivateMsg(userId int, message string,autoEscape bool) int32 {
	requestUrl := "http://%s:%d/send_private_msg?"
	requestUrl += "user_id=%d"
	requestUrl += "&message=%s"
	requestUrl += "&auto_escape=%s"
	url := fmt.Sprintf(requestUrl, bot.Address,bot.Port, userId,message,strconv.FormatBool(autoEscape))
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &responseMsgJson)
	return responseMsgJson.Data.MessageId
}

func (bot Bots) DeleteMsg(messageId int32) {
	requestUrl := "http://%s:%d/delete_msg?"
	requestUrl += "message_id=%d"
	url := fmt.Sprintf(requestUrl, bot.Address,bot.Port,messageId)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
}

func (bot Bots) GetMsg(messageId int32) GetMessage {
	requestUrl := "http://%s:%d/get_msg?"
	requestUrl += "message_id=%d"
	url := fmt.Sprintf(requestUrl, bot.Address,bot.Port,messageId)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	err := json.Unmarshal(responseByte, &getMessageJson)
	if err != nil {
		panic(err)
	}
	return getMessageJson.Data
}

func (bot Bots) SetGroupBan(groupId int32, userId int32, duration int) {
	requestUrl := "http://%s:%d/set_group_ban?"
	requestUrl += "group_id=%d"
	requestUrl += "&user_id=%d"
	requestUrl += "&duration=%d"
	url := fmt.Sprintf(requestUrl, bot.Address,bot.Port,groupId,userId,duration)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
}

func (bot Bots) SetGroupCard(groupId int32, userId int32, card string) {
	requestUrl := "http://%s:%d/set_group_card?"
	requestUrl += "group_id=%d"
	requestUrl += "&user_id=%d"
	requestUrl += "&card=%s"
	url := fmt.Sprintf(requestUrl, bot.Address,bot.Port,groupId,userId,card)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url,"\n\t\t\t\t\t",defaultJson.RetCode,defaultJson.Status)

}

func (bot Bots) SendMsg(messageType string, id int, message string, autoEscape bool) int32 {
	var MessageId int32
	if messageType == "group" {
		MessageId = bot.SendGroupMsg(id,message,autoEscape)
	}else if messageType == "private"{
		MessageId = bot.SendPrivateMsg(id,message,autoEscape)
	}else {
		log.Println("请正确指定messageType的值")
	}
	return MessageId
}

func (bot Bots) SendLike(userId int, times int) {
	requestUrl := "http://%s:%d/send_like?"
	requestUrl += "user_id=%d"
	requestUrl += "&times=%d"
	url := fmt.Sprintf(requestUrl,bot.Address,bot.Port,userId,times)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url,"\n\t\t\t\t\t",defaultJson.RetCode,defaultJson.Status)
}

func (bot Bots) SetGroupKick(groupId int, UserId int, rejectAddRequest bool) {
	requestUrl := "http://%s:%d/set_group_kick?"
	requestUrl += "group_id=%d"
	requestUrl += "&user_id=%d"
	requestUrl += "&reject_add_request=%s"
	url := fmt.Sprintf(requestUrl,bot.Address,bot.Port,groupId,UserId,strconv.FormatBool(rejectAddRequest))
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url,"\n\t\t\t\t\t",defaultJson.RetCode,defaultJson.Status)
}

func (bot Bots) SetGroupAnonymousBan(groupId int, flag string, duration int) {
	requestUrl := "http://%s:%d/set_group_anonymous_ban?"
	requestUrl += "group_id=%d"
	requestUrl += "&flag=%s"
	requestUrl += "&duration=%d"
	url := fmt.Sprintf(requestUrl,bot.Address,bot.Port,groupId,flag,duration)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url,"\n\t\t\t\t\t",defaultJson.RetCode,defaultJson.Status)
}

func (bot Bots) SetGroupWholeBan(groupId int, enable bool) {
	requestUrl := "http://%s:%d/set_group_whole_ban?"
	requestUrl += "group_id=%d"
	requestUrl += "&enable=%s"
	url := fmt.Sprintf(requestUrl,bot.Address,bot.Port,groupId,strconv.FormatBool(enable))
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url,"\n\t\t\t\t\t",defaultJson.RetCode,defaultJson.Status)
}

func (bot Bots) SetGroupAdmin(groupId int, UserId int, enable bool) {
	requestUrl := "http://%s:%d/set_group_admin?"
	requestUrl += "group_id=%d"
	requestUrl += "&user_id=%d"
	requestUrl += "&enable=%s"
	url := fmt.Sprintf(requestUrl,bot.Address,bot.Port,groupId,UserId,strconv.FormatBool(enable))
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url,"\n\t\t\t\t\t",defaultJson.RetCode,defaultJson.Status)
}

func (bot Bots) SetGroupAnonymous(groupId int, enable bool) {
	requestUrl := "http://%s:%d/set_group_anonymous?"
	requestUrl += "group_id=%d"
	requestUrl += "&enable=%s"
	url := fmt.Sprintf(requestUrl,bot.Address,bot.Port,groupId,strconv.FormatBool(enable))
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url,"\n\t\t\t\t\t",defaultJson.RetCode,defaultJson.Status)
}

func (bot Bots) SetGroupName(groupId int, groupName string) {
	requestUrl := "http://%s:%d/set_group_name?"
	requestUrl += "group_id=%d"
	requestUrl += "&group_name=%s"
	url := fmt.Sprintf(requestUrl,bot.Address,bot.Port,groupId,groupName)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url,"\n\t\t\t\t\t",defaultJson.RetCode,defaultJson.Status)
}

func (bot Bots) SetGroupLeave(groupId int, isDisMiss bool) {
	requestUrl := "http://%s:%d/set_group_leave?"
	requestUrl += "group_id=%d"
	requestUrl += "&is_dis_miss=%s"
	url := fmt.Sprintf(requestUrl,bot.Address,bot.Port,groupId,strconv.FormatBool(isDisMiss))
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url,"\n\t\t\t\t\t",defaultJson.RetCode,defaultJson.Status)
}

func (bot Bots) SetGroupSpecialTitle(groupId int, userId int, specialTitle string, duration int) {
	requestUrl := "http://%s:%d/set_group_special_title?"
	requestUrl += "group_id=%d"
	requestUrl += "&user_id=%d"
	requestUrl += "&special_title=%s"
	requestUrl += "&duration=%d"
	url := fmt.Sprintf(requestUrl,bot.Address,bot.Port,groupId,userId,specialTitle,duration)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url,"\n\t\t\t\t\t",defaultJson.RetCode,defaultJson.Status)
}

func (bot Bots) SetFriendAddRequest(flag string, approve bool, remark string) {
	requestUrl := "http://%s:%d/set_friend_add_request?"
	requestUrl += "flag=%s"
	requestUrl += "&approve=%s"
	requestUrl += "&remark=%s"
	url := fmt.Sprintf(requestUrl,bot.Address,bot.Port,flag,strconv.FormatBool(approve),remark)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url,"\n\t\t\t\t\t",defaultJson.RetCode,defaultJson.Status)
}

func (bot Bots) SetGroupAddRequest(flag string, subType string, approve bool, reason string) {
	requestUrl := "http://%s:%d/set_group_add_request?"
	requestUrl += "flag=%s"
	requestUrl += "&sub_type=%s"
	requestUrl += "&approve=%s"
	requestUrl += "&reason=%s"
	url := fmt.Sprintf(requestUrl,bot.Address,bot.Port,flag,subType,strconv.FormatBool(approve),reason)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url,"\n\t\t\t\t\t",defaultJson.RetCode,defaultJson.Status)
}

func (bot Bots) GetLoginInfo() LoginInfo {
	requestUrl := "http://%s:%d/set_group_add_request"
	url := fmt.Sprintf(requestUrl,bot.Address,bot.Port)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &LoginInfoJson)
	log.Println(url,"\n\t\t\t\t\t",defaultJson.RetCode,defaultJson.Status)
	return LoginInfoJson.Data
}

func (bot Bots) GetStrangerInfo() Senders {
	requestUrl := "http://%s:%d/get_stranger_info"
	url := fmt.Sprintf(requestUrl,bot.Address,bot.Port)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &StrangerInfo)
	log.Println(url,"\n\t\t\t\t\t",StrangerInfo.RetCode,StrangerInfo.Status)
	return StrangerInfo.Data
}

func (bot Bots) GetFriendList() []FriendList {
	requestUrl := "http://%s:%d/get_stranger_info"
	url := fmt.Sprintf(requestUrl,bot.Address,bot.Port)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &FriendListJson)
	log.Println(url,"\n\t\t\t\t\t",FriendListJson.RetCode,FriendListJson.Status)
	return FriendListJson.Data
}

func (bot Bots) GetGroupInfo(groupId int, noCache bool) GroupInfo {
	requestUrl := "http://%s:%d/get_group_info?"
	requestUrl += "group_id=%d"
	requestUrl += "no_cache=%s"
	url := fmt.Sprintf(requestUrl,bot.Address,bot.Port,groupId,strconv.FormatBool(noCache))
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &GroupInfoJson)
	log.Println(url,"\n\t\t\t\t\t",GroupInfoJson.RetCode,GroupInfoJson.Status)
	return  GroupInfoJson.Data
}

func (bot Bots) GetGroupList() []GroupInfo {
	panic("implement me")
}

func (bot Bots) GetGroupMemberInfo(groupId int, UserId int, noCache bool) GroupMemberInfo {
	panic("implement me")
}

func (bot Bots) GetGroupMemberList(groupId int) []GroupMemberInfo {
	panic("implement me")
}

func (bot Bots) GetGroupHonorInfo(groupId int, honorType string) GroupHonorInfo {
	panic("implement me")
}

func (bot Bots) GetCookies(domain string) Cookie {
	panic("implement me")
}

func (bot Bots) GetCsrfToken() CsrfToken {
	panic("implement me")
}

func (bot Bots) GetCredentials(domain string) Credentials {
	panic("implement me")
}

func (bot Bots) GetRecord() Record {
	panic("implement me")
}

func (bot Bots) GetImage(file string) {
	panic("implement me")
}

func (bot Bots) CanSendImage() Bool {
	panic("implement me")
}

func (bot Bots) CanSendRecord() Bool {
	panic("implement me")
}

func (bot Bots) GetStatus() OnlineStatus {
	panic("implement me")
}