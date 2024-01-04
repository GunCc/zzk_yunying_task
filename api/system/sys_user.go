package system

import (
	"ZZK_YUNYING_TASK/global"
	"ZZK_YUNYING_TASK/model/commen/response"
	"ZZK_YUNYING_TASK/model/system/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysUserApi struct {
}

// @Tags SysUser
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body request.Register true "昵称, 密码"
// @Success 200 {object} response.Response{msg=string} "用户注册账号,返回包括用户信息"
// @Router /user/register [post]
func (SysUserApi) Register(ctx *gin.Context) {
	var register request.Register

	err := ctx.ShouldBindJSON(&register)
	if err != nil {
		global.TASK_LOGGER.Error("注册信息有误", zap.Error(err))
		response.Fail(ctx)

		return
	}

	_, err = userService.Register(register)
	if err != nil {
		global.TASK_LOGGER.Error("注册失败", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}
	response.SuccessWithMessage("注册成功", ctx)
}
