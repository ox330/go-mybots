//api.go文件包含了可以调用的各个api
package go_mybots

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
	"os"
	"regexp"
	"strconv"
)

type Bots struct {
	Address string
	Port    int
	Admin   int
}

type Message struct {
	Message string `json:"message"`
}

type (
	ConstNoticeType struct {
		GroupUpload   string //群文件上传
		GroupAdmin    string //群管理员变动
		GroupDecrease string //群成员减少
		GroupIncrease string //群成员增加
		GroupBan      string //群禁言
		FriendAdd     string //好友添加
		GroupRecall   string //群消息撤回
		FriendRecall  string //好友消息撤回
		Notify        string //群内戳一戳,群红包运气王,群成员荣誉变更,好友戳一戳
		GroupCard     string //群成员名片更新
		OfflineFile   string //接收到离线文件
	}
	ConstMessageType struct {
		Group   string
		Private string
	}
	//ConstMessageSubType struct {
	//	Friend    string //好友消息
	//	Group     string //临时会话
	//	Other     string //其他消息
	//	Normal    string //正常群消息
	//	Anonymous string //匿名消息
	//	Notice    string //群通知消息
	//}
)

var (

	//MessageSubTypeApi = ConstMessageSubType{
	//	Friend: "friend",
	//	Group: "group",
	//	Other: "other",
	//	Normal: "normal",
	//	Anonymous: "anonymous",
	//	Notice: "notice"}
	NoticeTypeApi = ConstNoticeType{
		GroupUpload:   "group_upload",
		GroupAdmin:    "group_admin",
		GroupDecrease: "group_decrease",
		GroupIncrease: "group_increase",
		GroupBan:      "group_ban",
		FriendAdd:     "friend_add",
		GroupRecall:   "group_recall",
		FriendRecall:  "friend_recall",
		Notify:        "notify",
		OfflineFile:   "offline_file",
		GroupCard:     "group_card"}

	MessageTypeApi = ConstMessageType{
		Group:   "group",
		Private: "private"}
)

type MessageIds struct {
	MessageId int32 `json:"message_id"`
}

type LoginInfo struct {
	UserId   int    `json:"user_id"`
	NickName string `json:"nick_name"`
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
	TargetId      int64     `json:"target_id"` //运气王id
	HonorType     string    `json:"honor_type"`
	MetaEventType string    `json:"meta_event_type"`
	Status        string    `json:"status"`
	Interval      string    `json:"interval"`
	CardNew       string    `json:"card_new"` //新名片
	CardOld       string    `json:"card_old"` //旧名片
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

type responseJson struct {
	Status  string `json:"status"`
	RetCode int    `json:"retcode"`
}

type (
	anonymous struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Flag string `json:"flag"`
	}

	Files struct {
		Id      string `json:"id"`
		Name    string `json:"name"`
		Size    int64  `json:"size"`
		Busid   int64  `json:"busid"`
		FileUrl string `json:"url"`
	}

	GroupMemberInfo struct {
		GroupId         int  `json:"group_id"`
		JoinTime        int  `json:"join_time"`
		LastSentTime    int  `json:"last_sent_time"`
		Unfriendly      bool `json:"unfriendly"`
		TitleExpireTime int  `json:"title_expire_time"`
		CardChangeable  bool `json:"card_changeable"`
		Senders
	}

	getMsgJson struct {
		responseJson
		Data GetMessage `json:"data"`
	}

	GetMessage struct {
		Time      int32   `json:"time"`
		Group     bool    `json:"group"`
		MessageId int32   `json:"message_id"`
		RealId    int32   `json:"real_id"`
		Sender    Senders `json:"sender"`
		Message   string  `json:"message"`
	}

	FriendList struct {
		UserId   int    `json:"user_id"`
		NickName string `json:"nick_name"`
		Remark   string `json:"remark"`
	}

	GroupInfo struct {
		GroupId        int    `json:"group_id"`
		GroupName      string `json:"group_name"`
		MemberCount    int    `json:"member_count"`
		MaxMemberCount int    `json:"max_member_count"`
	}

	GroupHonorInfo struct {
		GroupId          int                  `json:"group_id"`
		CurrentTalkative CurrentTalkativeS    `json:"current_talkative"`
		TalkativeList    []GroupHonorInfoList `json:"talkative_list"`
		PerformerList    []GroupHonorInfoList `json:"performer_list"`
		LegendList       []GroupHonorInfoList `json:"legend_list"`
		StrongNewbieList []GroupHonorInfoList `json:"strong_newbie_list"`
		EmotionList      []GroupHonorInfoList `json:"emotion_list"`
	}

	CurrentTalkativeS struct {
		UserId   int    `json:"user_id"`
		NickName string `json:"nick_name"`
		Avatar   string `json:"avatar"`
		DayCount int    `json:"day_count"`
	}

	GroupHonorInfoList struct {
		UserId      int    `json:"user_id"`
		NickName    string `json:"nick_name"`
		Avatar      string `json:"avatar"`
		Description string `json:"description"`
	}

	Record struct {
		File      string `json:"file"`
		OutFormat string `json:"out_format"`
	}

	Cookie struct {
		Cookies string `json:"cookies"`
	}

	CsrfToken struct {
		Token string `json:"token"`
	}

	OnlineStatus struct {
		Online bool `json:"online"`
		Good   bool `json:"good"`
	}

	Bool struct {
		Yes bool `json:"yes"`
	}

	Image struct {
		File string `json:"file"`
	}

	Credentials struct {
		Cookie
		CsrfToken
	}
	Content struct {
		Type string `json:"type"`
		Data string `json:"data"`
	}
	Node struct {
		Id      int     `json:"id"`
		Name    string  `json:"name"`
		Uin     int     `json:"uin"`
		Content Content `json:"content"`
	}
)

type (
	defaultResponseJson struct {
		responseJson
		Data string `json:"data"`
	}

	responseGroupListJson struct {
		responseJson
		Data []GroupInfo `json:"data"`
	}

	responseStrangerInfoJson struct {
		responseJson
		Data Senders `json:"data"`
	}

	responseGroupInfoJson struct {
		responseJson
		Data GroupInfo `json:"data"`
	}

	responseFriendListJson struct {
		responseJson
		Data []FriendList `json:"data"`
	}

	responseMessageJson struct {
		responseJson
		Data MessageIds `json:"data"`
	}

	responseLoginIndoJson struct {
		responseJson
		Data LoginInfo `json:"data"`
	}

	responseGroupMemberInfoJson struct {
		responseJson
		Data GroupMemberInfo `json:"data"`
	}
	responseGroupMemberListJson struct {
		responseJson
		Data []GroupMemberInfo `json:"data"`
	}
	responseGroupHonorInfoJson struct {
		responseJson
		Data GroupHonorInfo `json:"data"`
	}
	responseCookiesJson struct {
		responseJson
		Data Cookie `json:"data"`
	}
	responseCsrfTokenJson struct {
		responseJson
		Data CsrfToken `json:"data"`
	}
	responseCredentialsJson struct {
		responseJson
		Data Credentials `json:"data"`
	}
	responseRecordJson struct {
		responseJson
		Data Record `json:"data"`
	}
	responseImageJson struct {
		responseJson
		Data Image `json:"data"`
	}
	responseCanSendJson struct {
		responseJson
		Data Bool `json:"data"`
	}
	responseOnlineStatus struct {
		responseJson
		Data OnlineStatus `json:"data"`
	}
	responseMsgDataJson struct {
		responseJson
		Data MsgData `json:"data"`
	}
	responseForwardMsgJson struct {
		responseJson
		Data []ForwardMsg `json:"data"`
	}
	responseWordSliceJson struct {
		responseJson
		Data []string `json:"data"`
	}
	responseOcrImageJson struct {
		responseJson
		Data OcrImage `json:"data"`
	}
	responseGroupSystemMsgJson struct {
		responseJson
		Data GroupSystemMsg `json:"data"`
	}
	responseGroupFileSystemInfoJson struct {
		responseJson
		Data GroupFileSystemInfo `json:"data"`
	}
	responseGroupRootFilesJson struct {
		responseJson
		Data GroupRootFiles `json:"data"`
	}
	responseGroupFilesByFolderJson struct {
		responseJson
		Data GroupFilesByFolder `json:"data"`
	}
	responseGroupFileUrlJson struct {
		responseJson
		Data fileUrl `json:"data"`
	}
	responseGroupAtAllRemainJson struct {
		responseJson
		Data GroupAtAllRemain `json:"data"`
	}
	responseDownloadFilePathJson struct {
		responseJson
		Data DownloadFilePath `json:"data"`
	}
	responseMessageHistoryJson struct {
		responseJson
		Data MessageHistory `json:"data"`
	}
	responseOnlineClientsJson struct {
		responseJson
		Data Clients `json:"data"`
	}
	responseVipInfoJson struct {
		responseJson
		Data VipInfo `json:"data"`
	}
)

type (
	fileUrl struct {
		Url string `json:"url"`
	}
	MsgData struct {
		MessageId int     `json:"message_id"`
		RealId    int     `json:"real_id"`
		Sender    Senders `json:"sender"`
		Time      int     `json:"time"`
		Message   string  `json:"message"`
	}
	ForwardMsg struct {
		Content string  `json:"content"`
		Sender  Senders `json:"sender"`
		Time    int     `json:"time"`
	}
	TextDetection struct {
		Text        string      `json:"text"`
		Confidence  int         `json:"confidence"`
		Coordinates interface{} `json:"coordinates"`
	}
	OcrImage struct {
		Texts    []TextDetection `json:"texts"`
		Language string          `json:"language"`
	}
	InvitedRequest struct {
		RequestId   int    `json:"request_id"`   //请求ID
		InvitorUin  int    `json:"invitor_uin"`  //邀请者
		InvitorNick string `json:"invitor_nick"` //邀请者昵称
		GroupId     int    `json:"group_id"`     //群号
		GroupName   string `json:"group_name"`   //群名
		Checked     bool   `json:"checked"`      //是否已被处理
		Actor       int64  `json:"actor"`        //处理者, 未处理为0
	}
	JoinRequest struct {
		RequestId     int    `json:"request_id"`     //请求ID
		RequesterUin  int    `json:"requester_uin"`  //请求者ID
		RequesterNick string `json:"requester_nick"` //请求者昵称
		Message       string `json:"message"`        //验证消息
		GroupId       int    `json:"group_id"`       //群号
		GroupName     string `json:"group_name"`     //群名
		Checked       bool   `json:"checked"`        //是否已被处理
		Actor         int    `json:"actor"`          //处理者, 未处理为0
	}
	GroupSystemMsg struct {
		InvitedRequests []InvitedRequest `json:"invited_requests"` //邀请消息列表
		JoinRequests    []JoinRequest    `json:"join_requests"`    //进群消息列表
	}
	GroupFileSystemInfo struct {
		FileCount  int `json:"file_count"`  //文件总数
		LimitCount int `json:"limit_count"` //文件上限
		UsedSpace  int `json:"used_space"`  //已使用空间
		TotalSpace int `json:"total_space"` //空间上限
	}
	File struct {
		FileId        string `json:"file_id"`        //文件ID
		FileName      string `json:"file_name"`      //文件名
		Busid         int    `json:"busid"`          //文件类型
		FileSize      int64  `json:"file_size"`      //文件大小
		UploadTime    int64  `json:"upload_time"`    //上传时间
		DeadTime      int64  `json:"dead_time"`      //过期时间,永久文件恒为0
		ModifyTime    int64  `json:"modify_time"`    //最后修改时间
		DownloadTimes int32  `json:"download_times"` //下载次数
		Uploader      int64  `json:"uploader"`       //上传者ID
		UploaderName  string `json:"uploader_name"`  //上传者名字
	}
	Folder struct {
		FolderId       string `json:"folder_id"`        //文件夹ID
		FolderName     string `json:"folder_name"`      //文件名
		CreateTime     int    `json:"create_time"`      //创建时间
		Creator        int    `json:"creator"`          //创建者
		CreatorName    string `json:"creator_name"`     //创建者名字
		TotalFileCount int32  `json:"total_file_count"` //子文件数量
	}
	GroupRootFiles struct {
		Files   []File   `json:"files"`
		Folders []Folder `json:"folders"`
	}
	GroupFilesByFolder struct {
		Files   []File   `json:"files"`
		Folders []Folder `json:"folders"`
	}
	GroupAtAllRemain struct {
		CanAtAll                 bool `json:"can_at_all"`                    //是否可以@全体成员
		RemainAtAllCountForGroup int  `json:"remain_at_all_count_for_group"` //群内所有管理当天剩余@全体成员次数
		RemainAtAllCountForUin   int  `json:"remain_at_all_count_for_uin"`   //BOT当天剩余@全体成员次数
	}
	DownloadFilePath struct {
		File string `json:"file"`
	}
	MessageHistory struct {
		Messages []string `json:"messages"`
	}
	Clients struct {
		Clients []Device `json:"clients"` //在线客户端列表
	}
	Device struct {
		AppId      int64  `json:"app_id"`      //客户端ID
		DeviceName string `json:"device_name"` //设备名称
		DeviceKind string `json:"device_kind"` //设备类型
	}
	VipInfo struct {
		UserId         int64   `json:"user_id"`          //QQ 号
		Nickname       string  `json:"nickname"`         //用户昵称
		Level          int64   `json:"level"`            //QQ 等级
		LevelSpeed     float64 `json:"level_speed"`      //等级加速度
		VipLevel       string  `json:"vip_level"`        //会员等级
		VipGrowthSpeed int64   `json:"vip_growth_speed"` //会员成长速度
		VipGrowthTotal int64   `json:"vip_growth_total"` //会员成长总值
	}
)

//go-cqhttp新增api
type IncreaseApi interface {
	DownloadFile(url string, threadCount int, headers []string) (DownloadFilePath, error)
	GetGroupMsgHistory(messageSeq int64, groupId int) (MessageHistory, error)
	GetOnlineClients(noCache bool) (Clients, error)
	GetVipInfoTest(UserId int) (VipInfo, error)
	SendGroupNotice(groupId int64, content string) error
	ReloadEventFilter() error
	UploadGroupFile(groupId int, file string, name string, folder string) error
}

type SpecialApi interface {
	SetGroupNameSpecial(groupId int, groupName string) error
	SetGroupPortrait(groupId int, file string, cache int) error
	GetMsgSpecial(messageId int) (MsgData, error)
	GetForwardMsg(messageId int) ([]ForwardMsg, error)
	SendGroupForwardMsg(groupId int, messages []Node) error
	GetWordSlices(content string) ([]string, error)
	OcrImage(image string) (OcrImage, error)
	GetGroupSystemMsg() (GroupSystemMsg, error)
	GetGroupFileSystemInfo(groupId int) (GroupFileSystemInfo, error)
	GetGroupRootFiles(groupId int) (GroupRootFiles, error)
	GetGroupFilesByFolder(groupId int, folderId string) (GroupFilesByFolder, error)
	GetGroupFileUrl(groupId int, fileId string, busid int) (fileUrl, error)
	GetGroupAtAllRemain(groupId int) (GroupAtAllRemain, error)
}

type Api interface {
	SendGroupMsg(groupId int, message string, autoEscape bool) (int32, error)
	SendPrivateMsg(userId int, message string, autoEscape bool) (int32, error)
	DeleteMsg(messageId int32) error
	GetMsg(messageId int32) (GetMessage, error)
	SetGroupBan(groupId int, userId int, duration int) error
	SetGroupCard(groupId int, userId int, card string) error
	SendMsg(messageType string, id int, message string, autoEscape bool) (int32, error)
	SendLike(userId int, times int) error
	SetGroupKick(groupId int, UserId int, rejectAddRequest bool) error
	SetGroupAnonymousBan(groupId int, flag string, duration int) error
	SetGroupWholeBan(groupId int, enable bool) error
	SetGroupAdmin(groupId int, UserId int, enable bool) error
	SetGroupAnonymous(groupId int, enable bool) error
	SetGroupName(groupId int, groupName string) error
	SetGroupLeave(groupId int, isDisMiss bool) error
	SetGroupSpecialTitle(groupId int, userId int, specialTitle string, duration int) error
	SetFriendAddRequest(flag string, approve bool, remark string) error
	SetGroupAddRequest(flag string, subType string, approve bool, reason string) error
	GetLoginInfo() (LoginInfo, error)
	GetStrangerInfo() (Senders, error)
	GetFriendList() ([]FriendList, error)
	GetGroupInfo(groupId int, noCache bool) (GroupInfo, error)
	GetGroupList() ([]GroupInfo, error)
	GetGroupMemberInfo(groupId int, UserId int, noCache bool) (GroupMemberInfo, error)
	GetGroupMemberList(groupId int) ([]GroupMemberInfo, error)
	GetGroupHonorInfo(groupId int, honorType string) (GroupHonorInfo, error)
	GetCookies(domain string) (Cookie, error)
	GetCsrfToken() (CsrfToken, error)
	GetCredentials(domain string) (Credentials, error)
	GetRecord(file, outFormat string) (Record, error)
	GetImage(file string) (Image, error)
	CanSendImage() (Bool, error)
	CanSendRecord() (Bool, error)
	GetStatus() (OnlineStatus, error)
	SetRestart(delay int)
	CleanCache()
}

var (
	responseMsgJson         responseMessageJson
	getMessageJson          getMsgJson
	defaultJson             defaultResponseJson
	LoginInfoJson           responseLoginIndoJson
	StrangerInfo            responseStrangerInfoJson
	FriendListJson          responseFriendListJson
	GroupInfoJson           responseGroupInfoJson
	GroupListJson           responseGroupListJson
	GroupMemberInfoJson     responseGroupMemberInfoJson
	GroupMemberListJson     responseGroupMemberListJson
	GroupHonorInfoJson      responseGroupHonorInfoJson
	CookiesJson             responseCookiesJson
	csrfTokenJson           responseCsrfTokenJson
	credentialsJson         responseCredentialsJson
	recordJson              responseRecordJson
	imageJson               responseImageJson
	canSendJson             responseCanSendJson
	onlineStatusJson        responseOnlineStatus
	msgJson                 responseMsgDataJson
	forwardMsgJson          responseForwardMsgJson
	wordSliceJson           responseWordSliceJson
	ocrImageJson            responseOcrImageJson
	groupSystemMsgJson      responseGroupSystemMsgJson
	groupFileSystemInfoJson responseGroupFileSystemInfoJson
	groupRootFilesJson      responseGroupRootFilesJson
	groupFilesByFolderJson  responseGroupFilesByFolderJson
	groupFileUrlJson        responseGroupFileUrlJson
	groupAtAllRemainJson    responseGroupAtAllRemainJson
	downloadFilePathJson    responseDownloadFilePathJson
	messageHistoryJson      responseMessageHistoryJson
	onlineClientsJson       responseOnlineClientsJson
	vipInfoJson             responseVipInfoJson
)

func (bot Bots) SendGroupMsg(groupId int, message string, autoEscape bool) (int32, error) {
	url := fmt.Sprintf("http://%s:%d/send_group_msg", bot.Address, bot.Port)
	data := fmt.Sprintf("{\"group_id\":%d,\"message\":\"%v\",\"auto_escape\":%v}", groupId, message, autoEscape)
	values := url2.Values{}
	values.Add("group_id", strconv.Itoa(groupId))
	values.Add("message", message)
	values.Add("auto_escape", fmt.Sprintf("%v", autoEscape))
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("newRequest error")
	}
	//req.Header.Set("Command-Type", "application/json")
	//client := http.Client{}
	//response, err := client.Do(req)
	//if err != nil {
	//	log.Panic("client error")
	//}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &responseMsgJson)
	log.Println(url, data, "\n\t\t\t\t\t", defaultJson.RetCode, defaultJson.Status)
	return responseMsgJson.Data.MessageId, err
}

func (bot Bots) SendPrivateMsg(userId int, message string, autoEscape bool) (int32, error) {
	url := fmt.Sprintf("http://%s:%d/send_private_msg", bot.Address, bot.Port)
	data := fmt.Sprintf("{\"user_id\":%d,\"message\":\"%s\",\"auto_escape\":%v}", userId, message, autoEscape)
	values := url2.Values{}
	values.Add("user_id", strconv.Itoa(userId))
	values.Add("message", message)
	values.Add("auto_escape", fmt.Sprintf("%v", autoEscape))
	//log.Println(data)
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("newRequest error")
	}
	//req.Header.Set("Command-Type", "application/json")
	//client := http.Client{}
	//response, err := client.Do(req)
	//if err != nil {
	//	log.Panic("client error")
	//}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &responseMsgJson)
	log.Println(url, data, "\n\t\t\t\t\t", defaultJson.RetCode, defaultJson.Status)
	//log.Println(responseMsgJson.Status, url, values.Encode())
	return responseMsgJson.Data.MessageId, err
}

func (bot Bots) DeleteMsg(messageId int32) error {
	requestUrl := "http://%s:%d/delete_msg?"
	requestUrl += "message_id=%d"
	url := fmt.Sprintf(requestUrl, bot.Address, bot.Port, messageId)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	return err1
}

func (bot Bots) GetMsg(messageId int32) (GetMessage, error) {
	requestUrl := "http://%s:%d/get_msg?"
	requestUrl += "message_id=%d"
	url := fmt.Sprintf(requestUrl, bot.Address, bot.Port, messageId)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	err := json.Unmarshal(responseByte, &getMessageJson)
	log.Println(url, "\n\t\t\t\t\t", defaultJson.RetCode, defaultJson.Status)
	if err != nil {
		panic(err)
	}
	return getMessageJson.Data, err1
}

func (bot Bots) SetGroupBan(groupId int, userId int, duration int) error {
	requestUrl := "http://%s:%d/set_group_ban?"
	requestUrl += "group_id=%d"
	requestUrl += "&user_id=%d"
	requestUrl += "&duration=%d"
	url := fmt.Sprintf(requestUrl, bot.Address, bot.Port, groupId, userId, duration)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	return err1
}

func (bot Bots) SetGroupCard(groupId int, userId int, card string) error {
	requestUrl := "http://%s:%d/set_group_card?"
	requestUrl += "group_id=%d"
	requestUrl += "&user_id=%d"
	requestUrl += "&card=%s"
	url := fmt.Sprintf(requestUrl, bot.Address, bot.Port, groupId, userId, card)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url, "\n\t\t\t\t\t", defaultJson.RetCode, defaultJson.Status)
	return err1
}

func (bot Bots) SendMsg(messageType string, id int, message string, autoEscape bool) (int32, error) {
	var MessageId int32
	var err error
	if messageType == "group" {
		MessageId, err = bot.SendGroupMsg(id, message, autoEscape)
	} else if messageType == "private" {
		MessageId, err = bot.SendPrivateMsg(id, message, autoEscape)
	} else {
		log.Println("请正确指定messageType的值")
	}
	return MessageId, err
}

func (bot Bots) SendLike(userId int, times int) error {
	requestUrl := "http://%s:%d/send_like?"
	requestUrl += "user_id=%d"
	requestUrl += "&times=%d"
	url := fmt.Sprintf(requestUrl, bot.Address, bot.Port, userId, times)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url, "\n\t\t\t\t\t", defaultJson.RetCode, defaultJson.Status)
	return err1
}

func (bot Bots) SetGroupKick(groupId int, UserId int, rejectAddRequest bool) error {
	requestUrl := "http://%s:%d/set_group_kick?"
	requestUrl += "group_id=%d"
	requestUrl += "&user_id=%d"
	requestUrl += "&reject_add_request=%s"
	url := fmt.Sprintf(requestUrl, bot.Address, bot.Port, groupId, UserId, strconv.FormatBool(rejectAddRequest))
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url, "\n\t\t\t\t\t", defaultJson.RetCode, defaultJson.Status)
	return err1
}

func (bot Bots) SetGroupAnonymousBan(groupId int, flag string, duration int) error {
	requestUrl := "http://%s:%d/set_group_anonymous_ban?"
	requestUrl += "group_id=%d"
	requestUrl += "&flag=%s"
	requestUrl += "&duration=%d"
	url := fmt.Sprintf(requestUrl, bot.Address, bot.Port, groupId, flag, duration)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url, "\n\t\t\t\t\t", defaultJson.RetCode, defaultJson.Status)
	return err1
}

func (bot Bots) SetGroupWholeBan(groupId int, enable bool) error {
	requestUrl := "http://%s:%d/set_group_whole_ban?"
	requestUrl += "group_id=%d"
	requestUrl += "&enable=%s"
	url := fmt.Sprintf(requestUrl, bot.Address, bot.Port, groupId, strconv.FormatBool(enable))
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url, "\n\t\t\t\t\t", defaultJson.RetCode, defaultJson.Status)
	return err1
}

func (bot Bots) SetGroupAdmin(groupId int, UserId int, enable bool) error {
	requestUrl := "http://%s:%d/set_group_admin?"
	requestUrl += "group_id=%d"
	requestUrl += "&user_id=%d"
	requestUrl += "&enable=%s"
	url := fmt.Sprintf(requestUrl, bot.Address, bot.Port, groupId, UserId, strconv.FormatBool(enable))
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url, "\n\t\t\t\t\t", defaultJson.RetCode, defaultJson.Status)
	return err1
}

func (bot Bots) SetGroupAnonymous(groupId int, enable bool) error {
	requestUrl := "http://%s:%d/set_group_anonymous?"
	requestUrl += "group_id=%d"
	requestUrl += "&enable=%s"
	url := fmt.Sprintf(requestUrl, bot.Address, bot.Port, groupId, strconv.FormatBool(enable))
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url, "\n\t\t\t\t\t", defaultJson.RetCode, defaultJson.Status)
	return err1
}

func (bot Bots) SetGroupName(groupId int, groupName string) error {
	requestUrl := "http://%s:%d/set_group_name?"
	requestUrl += "group_id=%d"
	requestUrl += "&group_name=%s"
	url := fmt.Sprintf(requestUrl, bot.Address, bot.Port, groupId, groupName)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url, "\n\t\t\t\t\t", defaultJson.RetCode, defaultJson.Status)
	return err1
}

func (bot Bots) SetGroupLeave(groupId int, isDisMiss bool) error {
	requestUrl := "http://%s:%d/set_group_leave?"
	requestUrl += "group_id=%d"
	requestUrl += "&is_dis_miss=%s"
	url := fmt.Sprintf(requestUrl, bot.Address, bot.Port, groupId, strconv.FormatBool(isDisMiss))
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url, "\n\t\t\t\t\t", defaultJson.RetCode, defaultJson.Status)
	return err1
}

func (bot Bots) SetGroupSpecialTitle(groupId int, userId int, specialTitle string, duration int) error {
	requestUrl := "http://%s:%d/set_group_special_title?"
	requestUrl += "group_id=%d"
	requestUrl += "&user_id=%d"
	requestUrl += "&special_title=%s"
	requestUrl += "&duration=%d"
	url := fmt.Sprintf(requestUrl, bot.Address, bot.Port, groupId, userId, specialTitle, duration)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url, "\n\t\t\t\t\t", defaultJson.RetCode, defaultJson.Status)
	return err1
}

func (bot Bots) SetFriendAddRequest(flag string, approve bool, remark string) error {
	requestUrl := "http://%s:%d/set_friend_add_request?"
	requestUrl += "flag=%s"
	requestUrl += "&approve=%s"
	requestUrl += "&remark=%s"
	url := fmt.Sprintf(requestUrl, bot.Address, bot.Port, flag, strconv.FormatBool(approve), remark)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url, "\n\t\t\t\t\t", defaultJson.RetCode, defaultJson.Status)
	return err1
}

func (bot Bots) SetGroupAddRequest(flag string, subType string, approve bool, reason string) error {
	requestUrl := "http://%s:%d/set_group_add_request?"
	requestUrl += "flag=%s"
	requestUrl += "&sub_type=%s"
	requestUrl += "&approve=%s"
	requestUrl += "&reason=%s"
	url := fmt.Sprintf(requestUrl, bot.Address, bot.Port, flag, subType, strconv.FormatBool(approve), reason)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url, "\n\t\t\t\t\t", defaultJson.RetCode, defaultJson.Status)
	return err1
}

func (bot Bots) GetLoginInfo() (LoginInfo, error) {
	defer func() {
		err := recover()
		if err != nil {
			log.Println("LoginInfo error")
			os.Exit(3)
		}
	}()
	requestUrl := "http://%s:%d/get_login_info"
	url := fmt.Sprintf(requestUrl, bot.Address, bot.Port)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &LoginInfoJson)
	log.Println(url, "\n\t\t\t\t\t", defaultJson.RetCode, defaultJson.Status)
	return LoginInfoJson.Data, err1
}

func (bot Bots) GetStrangerInfo() (Senders, error) {
	requestUrl := "http://%s:%d/get_stranger_info"
	url := fmt.Sprintf(requestUrl, bot.Address, bot.Port)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &StrangerInfo)
	log.Println(url, "\n\t\t\t\t\t", StrangerInfo.RetCode, StrangerInfo.Status)
	return StrangerInfo.Data, err1
}

func (bot Bots) GetFriendList() ([]FriendList, error) {
	requestUrl := "http://%s:%d/get_friend_list"
	url := fmt.Sprintf(requestUrl, bot.Address, bot.Port)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &FriendListJson)
	log.Println(url, "\n\t\t\t\t\t", FriendListJson.RetCode, FriendListJson.Status)
	return FriendListJson.Data, err1
}

func (bot Bots) GetGroupInfo(groupId int, noCache bool) (GroupInfo, error) {
	requestUrl := "http://%s:%d/get_group_info?"
	requestUrl += "group_id=%d"
	requestUrl += "no_cache=%s"
	url := fmt.Sprintf(requestUrl, bot.Address, bot.Port, groupId, strconv.FormatBool(noCache))
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &GroupInfoJson)
	log.Println(url, "\n\t\t\t\t\t", GroupInfoJson.RetCode, GroupInfoJson.Status)
	return GroupInfoJson.Data, err1
}

func (bot Bots) GetGroupList() ([]GroupInfo, error) {
	requestUrl := "http://%s:%d/get_group_list"
	url := fmt.Sprintf(requestUrl, bot.Address, bot.Port)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &GroupListJson)
	log.Println(url, "\n\t\t\t\t\t", GroupListJson.RetCode, GroupListJson.Status)
	return GroupListJson.Data, err1
}

func (bot Bots) GetGroupMemberInfo(groupId int, UserId int, noCache bool) (GroupMemberInfo, error) {
	requestUrl := "http://%s:%d/get_group_member_info?"
	requestUrl += "group_id=%d"
	requestUrl += "&user_id=%d"
	requestUrl += "&no_cache=%v"
	url := fmt.Sprintf(requestUrl, bot.Address, bot.Port, groupId, UserId, noCache)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &GroupMemberInfoJson)
	log.Println(url, "\n\t\t\t\t\t", GroupMemberInfoJson.RetCode, GroupMemberInfoJson.Status)
	return GroupMemberInfoJson.Data, err1
}

func (bot Bots) GetGroupMemberList(groupId int) ([]GroupMemberInfo, error) {
	requestUrl := "http://%s:%d/get_group_member_list?"
	requestUrl += "group_id=%d"
	url := fmt.Sprintf(requestUrl, bot.Address, bot.Port, groupId)
	response, err1 := http.Get(url)
	if err1 != nil {
		log.Panicf("上报到%s:%d失败\n", bot.Address, bot.Port)
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &GroupMemberListJson)
	log.Println(url, "\n\t\t\t\t\t", GroupMemberListJson.RetCode, GroupMemberListJson.Status)
	return GroupMemberListJson.Data, err1
}

func (bot Bots) GetGroupHonorInfo(groupId int, honorType string) (GroupHonorInfo, error) {
	url := fmt.Sprintf("http://%s:%d/get_group_honor_info", bot.Address, bot.Port)
	values := url2.Values{}
	values.Add("group_id", strconv.Itoa(groupId))
	values.Add("honor_type", honorType)
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &GroupHonorInfoJson)
	log.Println(url, values.Encode(), "\n\t\t\t\t\t", GroupHonorInfoJson.RetCode, GroupHonorInfoJson.Status)
	return GroupHonorInfoJson.Data, err
}

func (bot Bots) GetCookies(domain string) (Cookie, error) {
	url := fmt.Sprintf("http://%s:%d/get_cookies", bot.Address, bot.Port)
	values := url2.Values{}
	values.Add("domain", domain)
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &CookiesJson)
	log.Println(url, values.Encode(), "\n\t\t\t\t\t", CookiesJson.RetCode, CookiesJson.Status)
	return CookiesJson.Data, err
}

func (bot Bots) GetCsrfToken() (CsrfToken, error) {
	url := fmt.Sprintf("http://%s:%d/get_csrf_token", bot.Address, bot.Port)
	response, err := http.Get(url)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &csrfTokenJson)
	log.Println(url, "\n\t\t\t\t\t", csrfTokenJson.RetCode, csrfTokenJson.Status)
	return csrfTokenJson.Data, err
}

func (bot Bots) GetCredentials(domain string) (Credentials, error) {
	url := fmt.Sprintf("http://%s:%d/get_credentials", bot.Address, bot.Port)
	values := url2.Values{}
	values.Add("domain", domain)
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &credentialsJson)
	log.Println(url, values.Encode(), "\n\t\t\t\t\t", credentialsJson.RetCode, credentialsJson.Status)
	return credentialsJson.Data, err
}

func (bot Bots) GetRecord(file, outFormat string) (Record, error) {
	url := fmt.Sprintf("http://%s:%d/get_record", bot.Address, bot.Port)
	values := url2.Values{}
	values.Add("file", file)
	values.Add("out_format", outFormat)
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &recordJson)
	log.Println(url, values.Encode(), "\n\t\t\t\t\t", recordJson.RetCode, recordJson.Status)
	return recordJson.Data, err
}

func (bot Bots) GetImage(file string) (Image, error) {
	url := fmt.Sprintf("http://%s:%d/get_image", bot.Address, bot.Port)
	values := url2.Values{}
	values.Add("file", file)
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &imageJson)
	log.Println(url, values.Encode(), "\n\t\t\t\t\t", imageJson.RetCode, imageJson.Status)
	return imageJson.Data, err
}

func (bot Bots) CanSendImage() (Bool, error) {
	url := fmt.Sprintf("http://%s:%d/can_send_image", bot.Address, bot.Port)

	response, err := http.Get(url)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &canSendJson)
	log.Println(url, "\n\t\t\t\t\t", canSendJson.RetCode, canSendJson.Status)
	return canSendJson.Data, err
}

func (bot Bots) CanSendRecord() (Bool, error) {
	url := fmt.Sprintf("http://%s:%d/can_send_record", bot.Address, bot.Port)
	response, err := http.Get(url)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &canSendJson)
	log.Println(url, "\n\t\t\t\t\t", canSendJson.RetCode, canSendJson.Status)
	return canSendJson.Data, err
}

func (bot Bots) GetStatus() (OnlineStatus, error) {
	url := fmt.Sprintf("http://%s:%d/can_send_image", bot.Address, bot.Port)
	response, err := http.Get(url)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &onlineStatusJson)
	log.Println(url, "\n\t\t\t\t\t", onlineStatusJson.RetCode, onlineStatusJson.Status)
	return onlineStatusJson.Data, err
}

func (bot Bots) SetRestart(delay int) {
	url := fmt.Sprintf("http://%s:%d/set_restart", bot.Address, bot.Port)
	values := url2.Values{}
	values.Add("delay", strconv.Itoa(delay))
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url, values.Encode(), "\n\t\t\t\t\t", defaultJson.RetCode, defaultJson.Status)
}

func (bot Bots) CleanCache() {
	url := fmt.Sprintf("http://%s:%d/can_send_image", bot.Address, bot.Port)
	_, err := http.Get(url)
	if err != nil {
		log.Panic("client error")
	}
	log.Println("已发送清理缓存请求")
}

//go-cqhttp  APi

func (bot Bots) SetGroupNameSpecial(groupId int, groupName string) error {
	url := fmt.Sprintf("http://%s:%d/send_group_name", bot.Address, bot.Port)
	values := url2.Values{}
	values.Add("group_id", strconv.Itoa(groupId))
	values.Add("group_name", groupName)
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &responseMsgJson)
	log.Println(url, values.Encode(), "\n\t\t\t\t\t", defaultJson.RetCode, defaultJson.Status)
	return err
}

func (bot Bots) SetGroupPortrait(groupId int, file string, cache int) error {
	url := fmt.Sprintf("http://%s:%d/send_group_portrait", bot.Address, bot.Port)
	values := url2.Values{}
	values.Add("group_id", strconv.Itoa(groupId))
	values.Add("file", file)
	values.Add("cache", strconv.Itoa(cache))
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &responseMsgJson)
	log.Println(url, values.Encode(), "\n\t\t\t\t\t", defaultJson.RetCode, defaultJson.Status)
	return err
}

func (bot Bots) GetMsgSpecial(messageId int) (MsgData, error) {
	url := fmt.Sprintf("http://%s:%d/get_msg", bot.Address, bot.Port)
	values := url2.Values{}
	values.Add("message_id", strconv.Itoa(messageId))
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &msgJson)
	log.Println(url, values.Encode(), "\n\t\t\t\t\t", msgJson.RetCode, msgJson.Status)
	return msgJson.Data, err
}

func (bot Bots) GetForwardMsg(messageId int) ([]ForwardMsg, error) {
	url := fmt.Sprintf("http://%s:%d/get_forward_msg", bot.Address, bot.Port)
	values := url2.Values{}
	values.Add("message_id", strconv.Itoa(messageId))
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &forwardMsgJson)
	log.Println(url, values.Encode(), "\n\t\t\t\t\t", forwardMsgJson.RetCode, forwardMsgJson.Status)
	return forwardMsgJson.Data, err
}

func (bot Bots) SendGroupForwardMsg(groupId int, messages []Node) error {
	url := fmt.Sprintf("http://%s:%d/send_group_forward_msg", bot.Address, bot.Port)
	values := url2.Values{}
	values.Add("group_id", strconv.Itoa(groupId))
	values.Add("messages", fmt.Sprintf("%v", messages))
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url, values.Encode(), "\n\t\t\t\t\t", defaultJson.RetCode, defaultJson.Status)
	return err
}

func (bot Bots) GetWordSlices(content string) ([]string, error) {
	url := fmt.Sprintf("http://%s:%d/.get_word_slices", bot.Address, bot.Port)
	values := url2.Values{}
	values.Add("content", content)
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &wordSliceJson)
	log.Println(url, values.Encode(), "\n\t\t\t\t\t", wordSliceJson.RetCode, wordSliceJson.Status)
	return wordSliceJson.Data, err
}

func (bot Bots) OcrImage(image string) (OcrImage, error) {
	url := fmt.Sprintf("http://%s:%d/.ocr_image", bot.Address, bot.Port)
	values := url2.Values{}
	values.Add("image", image)
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &ocrImageJson)
	log.Println(url, values.Encode(), "\n\t\t\t\t\t", ocrImageJson.RetCode, ocrImageJson.Status)
	return ocrImageJson.Data, err
}

func (bot Bots) GetGroupSystemMsg() (GroupSystemMsg, error) {
	url := fmt.Sprintf("http://%s:%d/get_group_system_msg", bot.Address, bot.Port)
	values := url2.Values{}
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &groupSystemMsgJson)
	log.Println(url, values.Encode(), "\n\t\t\t\t\t", groupSystemMsgJson.RetCode, groupSystemMsgJson.Status)
	return groupSystemMsgJson.Data, err
}

func (bot Bots) GetGroupFileSystemInfo(groupId int) (GroupFileSystemInfo, error) {
	url := fmt.Sprintf("http://%s:%d/get_group_file_system_info", bot.Address, bot.Port)
	values := url2.Values{}
	values.Add("group_id", strconv.Itoa(groupId))
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &groupFileSystemInfoJson)
	log.Println(url, values.Encode(), "\n\t\t\t\t\t", groupFileSystemInfoJson.RetCode, groupFileSystemInfoJson.Status)
	return groupFileSystemInfoJson.Data, err
}

func (bot Bots) GetGroupRootFiles(groupId int) (GroupRootFiles, error) {
	url := fmt.Sprintf("http://%s:%d/get_group_root_files", bot.Address, bot.Port)
	values := url2.Values{}
	values.Add("group_id", strconv.Itoa(groupId))
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &groupRootFilesJson)
	log.Println(url, values.Encode(), "\n\t\t\t\t\t", groupRootFilesJson.RetCode, groupRootFilesJson.Status)
	return groupRootFilesJson.Data, err
}

func (bot Bots) GetGroupFilesByFolder(groupId int, folderId string) (GroupFilesByFolder, error) {
	url := fmt.Sprintf("http://%s:%d/get_group_files_by_folder", bot.Address, bot.Port)
	values := url2.Values{}
	values.Add("group_id", strconv.Itoa(groupId))
	values.Add("folder_id", folderId)
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &groupFilesByFolderJson)
	log.Println(url, values.Encode(), "\n\t\t\t\t\t", groupFilesByFolderJson.RetCode, groupFilesByFolderJson.Status)
	return groupFilesByFolderJson.Data, err
}

func (bot Bots) GetGroupFileUrl(groupId int, fileId string, busid int) (fileUrl, error) {
	url := fmt.Sprintf("http://%s:%d/get_group_file_url", bot.Address, bot.Port)
	values := url2.Values{}
	values.Add("group_id", strconv.Itoa(groupId))
	values.Add("file_id", fileId)
	values.Add("busid", strconv.Itoa(busid))
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &groupFileUrlJson)
	log.Println(url, values.Encode(), "\n\t\t\t\t\t", groupFileUrlJson.RetCode, groupFileUrlJson.Status)
	return groupFileUrlJson.Data, err
}

func (bot Bots) GetGroupAtAllRemain(groupId int) (GroupAtAllRemain, error) {
	url := fmt.Sprintf("http://%s:%d/get_group_at_all_remain", bot.Address, bot.Port)
	values := url2.Values{}
	values.Add("group_id", strconv.Itoa(groupId))
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &groupAtAllRemainJson)
	log.Println(url, values.Encode(), "\n\t\t\t\t\t", groupAtAllRemainJson.RetCode, groupAtAllRemainJson.Status)
	return groupAtAllRemainJson.Data, err
}

func (bot Bots) DownloadFile(url string, threadCount int, headers []string) (DownloadFilePath, error) {
	urls := fmt.Sprintf("http://%s:%d/download_file", bot.Address, bot.Port)
	values := url2.Values{}
	values.Add("url", url)
	values.Add("thread_count", strconv.Itoa(threadCount))
	values.Add("headers", fmt.Sprintf("%v", headers))
	response, err := http.PostForm(urls, values)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &downloadFilePathJson)
	log.Println(urls, values.Encode(), "\n\t\t\t\t\t", downloadFilePathJson.RetCode, downloadFilePathJson.Status)
	return downloadFilePathJson.Data, err
}

func (bot Bots) GetGroupMsgHistory(messageSeq int64, groupId int) (MessageHistory, error) {
	url := fmt.Sprintf("http://%s:%d/get_group_message_history", bot.Address, bot.Port)
	values := url2.Values{}
	values.Add("message_seq", strconv.FormatInt(messageSeq, 10))
	values.Add("group_id", strconv.Itoa(groupId))
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &messageHistoryJson)
	log.Println(url, values.Encode(), "\n\t\t\t\t\t", messageHistoryJson.RetCode, messageHistoryJson.Status)
	return messageHistoryJson.Data, err
}

func (bot Bots) GetOnlineClients(noCache bool) (Clients, error) {
	url := fmt.Sprintf("http://%s:%d/get_online_clients", bot.Address, bot.Port)
	values := url2.Values{}
	values.Add("no_cache", strconv.FormatBool(noCache))
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &onlineClientsJson)
	log.Println(url, values.Encode(), "\n\t\t\t\t\t", onlineClientsJson.RetCode, onlineClientsJson.Status)
	return onlineClientsJson.Data, err
}

func (bot Bots) GetVipInfoTest(UserId int) (VipInfo, error) {
	url := fmt.Sprintf("http://%s:%d/_get_vip_info", bot.Address, bot.Port)
	values := url2.Values{}
	values.Add("user_id", strconv.Itoa(UserId))
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &vipInfoJson)
	log.Println(url, values.Encode(), "\n\t\t\t\t\t", vipInfoJson.RetCode, vipInfoJson.Status)
	return vipInfoJson.Data, err
}

func (bot Bots) SendGroupNotice(groupId int64, content string) error {
	url := fmt.Sprintf("http://%s:%d/send_group_notice", bot.Address, bot.Port)
	values := url2.Values{}
	values.Add("group_id", strconv.FormatInt(groupId, 10))
	values.Add("content", content)
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url, values.Encode(), "\n\t\t\t\t\t", defaultJson.RetCode, defaultJson.Status)
	return err
}

func (bot Bots) ReloadEventFilter() error {
	url := fmt.Sprintf("http://%s:%d/get_group_at_all_remain", bot.Address, bot.Port)
	response, err := http.PostForm(url, nil)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url, "\n\t\t\t\t\t", defaultJson.RetCode, defaultJson.Status)
	return err
}

func (bot Bots) UploadGroupFile(groupId int, file string, name string, folder string) error {
	url := fmt.Sprintf("http://%s:%d/upload_group_file", bot.Address, bot.Port)
	values := url2.Values{}
	values.Add("group_id", strconv.Itoa(groupId))
	values.Add("file", url2.PathEscape(file))
	values.Add("name", name)
	if folder != "" {
		values.Add("folder", folder)
	}
	response, err := http.PostForm(url, values)
	if err != nil {
		log.Panic("client error")
	}
	defer response.Body.Close()
	responseByte, _ := ioutil.ReadAll(response.Body)
	_ = json.Unmarshal(responseByte, &defaultJson)
	log.Println(url, values.Encode(), "\n\t\t\t\t\t", defaultJson.RetCode, defaultJson.Status)
	return err
}

//MessageImage
func MessageImage(path string) Message {
	return Message{fmt.Sprintf("[CQ:image.file=file:/%s]", path)}
}

func MessageAt(UserId int) Message {
	return Message{fmt.Sprintf("[CQ:at,qq=%d]", UserId)}
}

//MatchImage
func MatchImage(m Message) []string {
	reg := regexp.MustCompile(`[CQ:image,file=.]`)
	if reg == nil {
		log.Panic("regexp error")
	}
	match := reg.FindAllString(m.Message, -1)
	return match
}
