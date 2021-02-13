package message

import (
	"strings"
)

type Message struct {
	Message string `json:"message"`

	Id int `json:"id"`

	Magic int `json:"magic"` //发送语音时可选，表示是否变声

	QQ int `json:"qq"`

	Fle     string `json:"fle"`      //文件路径或下载链接
	Type    string `json:"type"`     //图片是否为闪照 闪照为flash，秀图为show  发送音乐时 qq表示qq音乐，163表示网易云音乐，custom表示自定义分享
	Url     string `json:"url"`      //网络url
	Cache   int    `json:"cache"`    //只在通过网络 URL 发送时有效，表示是否使用已缓存的文件，默认 1
	Proxy   int    `json:"proxy"`    //只在通过网络 URL 发送时有效，表示是否通过代理下载文件（需通过环境变量或配置文件配置代理），默认 1
	TimeOut int    `json:"time_out"` //只在通过网络 URL 发送时有效，单位秒，表示下载网络文件的超时时间，默认不超时
	C       int    `json:"c"`        //下载时使用的线程数

	Audio string `json:"audio"` //分享音乐时音乐Url

	Lat float32 `json:"lat"` //纬度
	Lon float32 `json:"lon"` //经度

	Text string `json:"text"` //回复的消息
	Time int64  `json:"time"` //自定义回复消息时的时间戳

	Cover string `json:"cover"` //发送短视频时的封面消息

	Resid int `json:"resid"` //0为走小程序通道，其他为富文本通道

	Title   string `json:"title"`   //标题
	Content string `json:"content"` //内容
	Image   string `json:"image"`   //图片链接，可选
}

//json转义
func JsonEscape(s string) string {
	s = strings.Replace(s, ",", "&#44;", -1)
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "[", "&#91;")
	s = strings.ReplaceAll(s, "]", "&#93;")
	return s
}
