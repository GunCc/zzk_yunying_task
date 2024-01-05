package system

import (
	"ZZK_YUNYING_TASK/global"
	"ZZK_YUNYING_TASK/model/commen/response"
	systemRes "ZZK_YUNYING_TASK/model/system/response"

	"ZZK_YUNYING_TASK/model/system"
	"ZZK_YUNYING_TASK/model/system/request"
	"ZZK_YUNYING_TASK/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SysUserApi struct {
}

// @Tags    SysUser
// @Summary 用户注册账号
// @Produce application/json
// @Param   data body     request.Register              true "昵称, 密码"
// @Success 200  {object} response.Response{msg=string} "用户注册账号,返回包括用户信息"
// @Router  /user/register [post]
func (SysUserApi) Register(ctx *gin.Context) {
	var register request.Register

	err := ctx.ShouldBindJSON(&register)
	if err != nil {
		global.TASK_LOGGER.Error("注册信息有误", zap.Error(err))
		response.Fail(ctx)

		return
	}

	// 注册数据校验
	if err := utils.Verify(register, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
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

// @Tags    SysUser
// @Summary 用户登陆账号
// @Produce application/json
// @Param   data body     request.Login                                                true "昵称, 密码"
// @Success 200  {object} response.Response{data=systemRes.UserLoginAfter,msg=string} "返回用户信息和Token"
// @Router  /user/login [post]
func (s *SysUserApi) Login(ctx *gin.Context) {
	var login request.Login

	err := ctx.ShouldBindJSON(&login)

	if err != nil {
		global.TASK_LOGGER.Error("登录参数错误", zap.Error(err))
		response.FailWithMessage("登录失败", ctx)
		return
	}

	// 登录数据校验
	if err := utils.Verify(login, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	user, err := userService.Login(login)
	if err != nil {
		global.TASK_LOGGER.Error("错误:", zap.Error(err))
		response.FailWithMessage(err.Error(), ctx)
		return
	}

	s.TokenNext(ctx, user)
}

// 登录后下发Token
func (s *SysUserApi) TokenNext(c *gin.Context, user system.SysUser) {
	j := &utils.JWT{
		SigningKey: []byte(global.TASK_CONFIG.JWT.SigningKey),
	}

	claims := j.CreateClaims(request.BaseClaims{
		UUID:     user.UUID,
		ID:       user.ID,
		NickName: user.NickName,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.TASK_LOGGER.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}

	response.SuccessWithDetailed(systemRes.UserLoginAfter{
		Token: token,
		User:  user,
	}, "登录成功", c)
}
