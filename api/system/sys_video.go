package system

import (
	"ZZK_YUNYING_TASK/global"
	"ZZK_YUNYING_TASK/model/commen/response"
	"ZZK_YUNYING_TASK/model/system"
	"ZZK_YUNYING_TASK/utils"

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
// @Param start_time formData string true  "开始时间"
// @Param end_time formData string true  "结束时间"
// @Success 200 {object} response.Response{data=system.SysVideo,msg=string} "上传文件示例,返回包括文件详情"
// @Router /fileUploadAndDownload/upload [post]
func (v *SysVideoApi) UploadFile(ctx *gin.Context) {
	var file system.SysVideo
	// 接收文件 参数为文件字段
	_, header, err := ctx.Request.FormFile("file")
	file.StartTime = ctx.PostForm("start_time")
	file.EndTime = ctx.PostForm("end_time")
	if err != nil {
		global.TASK_LOGGER.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", ctx)
		return
	}

	// 获取用户信息
	claims, _ := utils.GetClaims(ctx)
	file.UserId = claims.ID

	file, err = SysVideoService.UploadVideo(header, file)
	if err != nil {
		global.TASK_LOGGER.Error("修改数据库链接失败!", zap.Error(err))
		response.FailWithMessage("修改数据库链接失败", ctx)
		return
	}
	response.SuccessWithDetailed(file, "上传成功", ctx)
}
