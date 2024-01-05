package system

import (
	"ZZK_YUNYING_TASK/global"
	"ZZK_YUNYING_TASK/model/commen/request"
	"ZZK_YUNYING_TASK/model/commen/response"
	sysReq "ZZK_YUNYING_TASK/model/system/request"
	"ZZK_YUNYING_TASK/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysVideoApi struct {
}

// @Tags SysVideo
// @Summary 上传文件示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true  "上传文件"
// @Param start_time formData string false  "开始时间"
// @Param end_time formData string false  "结束时间"
// @Success 200 {object} response.Response{data=system.SysVideo,msg=string} "上传文件示例"
// @Router /fileUploadAndDownload/upload [post]
func (v *SysVideoApi) UploadFile(ctx *gin.Context) {
	var videoParams sysReq.UploadVideoParams
	// 接收文件 参数为文件字段
	_, header, err := ctx.Request.FormFile("file")
	videoParams.StartTime = ctx.PostForm("start_time")
	videoParams.EndTime = ctx.PostForm("end_time")

	// 是否保存源文件 0 代表不 1 代表保存
	videoParams.Save = "0"
	if err != nil {
		global.TASK_LOGGER.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", ctx)
		return
	}

	// 获取用户信息
	claims, _ := utils.GetClaims(ctx)
	videoParams.UserId = claims.ID

	file, err := SysVideoService.UploadVideo(header, videoParams)
	if err != nil {
		global.TASK_LOGGER.Error("视频裁剪失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.SuccessWithDetailed(file, "上传成功", ctx)
}

// @Tags    SysVideo
// @Summary 获取视频资源列表（授权）
// @Security ApiKeyAuth
// @Produce application/json
// @Param   data body     request.ListInfo             false "页码, 页面大小"
// @Success 200  {object} response.Response{data=response.ListRes,msg=string} "视频列表信息"
// @Router  /fileUploadAndDownload/getVideoList [post]
func (v *SysVideoApi) GetVideoList(ctx *gin.Context) {
	var info request.ListInfo

	claims, _ := utils.GetClaims(ctx)
	user_id := claims.ID

	err := ctx.ShouldBindJSON(&info)
	if err != nil {
		global.TASK_LOGGER.Error("获取视频列表参数错误", zap.Error(err))
		response.FailWithMessage("获取视频列表参数错误", ctx)
		return
	}

	fmt.Println("user_id", user_id)

	list, total, err := SysVideoService.GetVideoListByUserId(info, user_id)
	if err != nil {
		global.TASK_LOGGER.Error("获取视频列表参数错误", zap.Error(err))
		response.FailWithMessage("获取视频列表参数错误", ctx)
		return
	}
	response.SuccessWithDetailed(response.ListRes{
		List:     list,
		Total:    total,
		Page:     info.Page,
		PageSize: info.PageSize,
	}, "数据获取成功", ctx)
}

// @Tags    SysVideo
// @Summary 下载视频资源
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param  id query string true "视频id"
// @Success 200
// @Router  /fileUploadAndDownload/download [get]
func (v *SysVideoApi) DownloadVideo(ctx *gin.Context) {

	fileId := ctx.Query("id")

	video, err := SysVideoService.DownloadVideo(fileId)

	if err != nil {
		global.TASK_LOGGER.Error("下载失败!", zap.Error(err))
		response.FailWithMessage("下载失败", ctx)
		return
	}

	ctx.Writer.Header().Add("Content-Type", "application/octet-stream")
	ctx.Writer.Header().Add("Content-Disposition", "attachment; filename="+video.OutputFileName)
	ctx.Writer.Header().Add("success", "true")
	ctx.File(video.Url)
}

// @Tags    SysVideo
// @Summary 获取视频裁剪进度
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param  id query string true "视频id"
// @Success 200
// @Router  /fileUploadAndDownload/getVideoProgress [get]
func (v *SysVideoApi) GetVideoProgress(ctx *gin.Context) {
	// w := ctx.Writer
	// r := ctx.Request
	// fmt.Println(r.Method)
	// appId := r.URL.Query()["appId"]
	// page := r.URL.Query()["page"]
	// pageSize := r.URL.Query()["pageSize"]
	// fmt.Println("获取到参数")
	// fmt.Println(appId)
	// fmt.Println(page)
	// fmt.Println(pageSize)

	// w.Header().Set("Content-Type", "text/event-stream")
	// w.Header().Set("Cache-Control", "no-cache")
	// w.Header().Set("Connection", "keep-alive")
	// w.Header().Set("Access-Control-Allow-Origin", "*")

	// flusher, ok := w.(http.Flusher)
	// if !ok {
	// 	log.Panic("server not support")
	// }
	// for {
	// 	if getInput == "" || getInput == "结束" {
	// 		return
	// 	}

	// 	time.Sleep(1 * time.Second)
	// 	//fmt.Fprintf(w, "data: %d%s%s%s\n\n", i, appId[0], page[0], pageSize[0])
	// 	//now := time.Now()
	// 	//timeStr := now.Format("2006-01-02 15:04:05")
	// 	fmt.Fprintf(w, "data: %s\n\n", getInput)
	// 	flusher.Flush()
	// }
	// fmt.Fprintf(w, "event: close\ndata: close\n\n") // 一定要带上data，否则无效
}
