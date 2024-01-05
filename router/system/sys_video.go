package system

import (
	"ZZK_YUNYING_TASK/api"

	"github.com/gin-gonic/gin"
)

type VideoRouter struct {
}

func (VideoRouter) InitVideoRouter(router *gin.RouterGroup) {
	videoRouter := router.Group("fileUploadAndDownload")
	videoApi := api.ApiGroupApp.SystemApiGroup.SysVideoApi
	{
		// 上传视频文件路由
		videoRouter.POST("upload", videoApi.UploadFile)
		// 获取资源列表
		videoRouter.POST("getVideoList", videoApi.GetVideoList)
	}
}
