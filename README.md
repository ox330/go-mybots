# go-mybots

go进行onebot的http简单实现

## 基本使用
首先需要一个qq无头客户端和一个go运行环境。

[go语言下载官网](https://golang.org/)  !!国内可访问<https://studygolang.com/dl>下载

+ [go-cqhttp](https://github.com/Mrs4s/go-cqhttp)
+ [mirai-native加载cqhttp](https://github.com/iTXTech/mirai-native)

这里以go-cqhttp为例，go-cqhttp配置详情见[原文档](https://github.com/Mrs4s/go-cqhttp/blob/master/docs/config.md)
首先下载go-cqhttp发行版，然后解压后打开exe文件，第一次打开会生成配置文件。然后打开config.hjson文件，旧版为config.json。
然后配置基本的bot所使用的qq号和密码。除此还需要配置。
```
    http_config: {
        // 是否启用正向HTTP服务器
        enabled: true
        // 服务端监听地址
        host: 
        // 服务端监听端口
        port: 
        // 反向HTTP超时时间, 单位秒
        // 最小值为5，小于5将会忽略本项设置
        timeout: 0
        // 反向HTTP POST地址列表
        // 格式: 
        // {
        //    地址: secret
        // }
        post_urls: {}
```

enable设置为true，host和port是go-cqhttp所监听的ip和端口，post_url为go-cqhttp事件上报地址,注意post_urls填写格式。

`go get github.com/3343780376/go-mybots`

创建一个项目，main.go

```
import (
	"github.com/3343780376/go-mybots"
	"log"
)

//实例化一个Bot对象，参数为go-cqhttp监听的地址和端口，以及你的作为测试给bot发信息的账号
var Bot = go_mybots.Bots{Address: "127.0.0.1", Port: 5700,Admin: 1743224847}

//将handle函数加入go_mybots的消息路由
func init(){
    go_mybots.ViewMessage = append(go_mybots.ViewMessage, go_mybots.ViewMessageApi{OnMessage: handle})
}

func main() {
	hand := go_mybots.Hand()
	err := hand.Run("127.0.0.1:8000")  //设置该项目监听地址，及为go-cqhttp的上报地址
	if err != nil {
		log.Println(err.Error())
	}
}

//event参数为上报的消息，详情可参考onebot文档
func handle(event go-mybots.events){
    //判断消息为hello且发送消息的账号为测试账号
    if event.Message=="hello"&&event.UserId== Bot.Admin {
            //异步调用SendGroupMsg这个api,三个参数分别为，发送的对象qq号，消息内容和是否解析cq码
    		go Bot.SendPrivateMsg(event.UserId,"hello,world",false)
    	}
}
```

此时使用测试的qq号给bot发送hello，即可看到

<img src="https://github.com/3343780376/go-mybots/blob/master/test1.png" />

具体实现更多的逻辑，可先参考onebot[文档](https://github.com/howmanybots/onebot)。

##有兴趣可加qq3343780376
