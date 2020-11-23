# go-mybots

go进行onebot的http简单实现

##基本使用
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

clone本项目，修改main.go的run地址为上面所填写的post_urls，修改test文件夹下test.go的Adress和port，与上面所配置的adress和port
相同，修改Admin为你自己所使用的qq号，此时运行main.go，然后用所设置的admin账号给bot发送私聊信息"hello",此时会收到bot回复的"hello world".

具体实现更多的逻辑，可先参考onebot[文档](https://github.com/howmanybots/onebot)。

##本人为新手初学，不喜勿喷。

