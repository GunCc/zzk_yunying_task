package system

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type SysNotifyService struct {
}

var SysNotifyServiceApp = new(SysNotifyService)

var ifChannelsMapInit = false

var channelsMap = map[string]chan string{}

func initChannelsMap() {
	channelsMap = make(map[string]chan string)
}

func AddChannel(uid string, traceId string) {
	if !ifChannelsMapInit {
		initChannelsMap()
		ifChannelsMapInit = true
	}
	var newChannel = make(chan string)
	channelsMap[uid+traceId] = newChannel
	fmt.Println("与用户id为" + uid + "，通道ID" + traceId)
}
func (v *SysNotifyService) BuildNotificationChannel(uid string, traceId string, c *gin.Context) {
	AddChannel(uid, traceId)
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	fmt.Println("用户ID", uid)
	w := c.Writer
	flusher, _ := w.(http.Flusher)
	closeNotify := c.Request.Context().Done()
	go func() {
		<-closeNotify
		delete(channelsMap, uid+traceId)
		fmt.Println("关闭与用户id为：" + uid + "的链接，通道ID" + traceId)
		return
	}()
	for msg := range channelsMap[uid+traceId] {
		fmt.Println("数据循环了，" + msg)
		fmt.Fprintf(w, "data: %s\n\n", msg)
		flusher.Flush()
	}
}

func (v *SysNotifyService) SendNotification(uid string) {
	fmt.Println("发送消息给用户id：" + uid)

	for key := range channelsMap {
		if strings.Contains(key, uid) {
			channel := channelsMap[key]
			channel <- "你有视频完成啦！"
		}
	}
}
