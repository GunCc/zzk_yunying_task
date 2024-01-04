package system

import (
	"ZZK_YUNYING_TASK/global"
	"ZZK_YUNYING_TASK/model/system"
	"ZZK_YUNYING_TASK/model/system/request"
	"ZZK_YUNYING_TASK/utils"
	"errors"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type SysUserService struct {
}

// @function: Register
// @description: 注册用户
// @param: register request.Register
// @return: user *system.SysUser, err error
func (SysUserService) Register(register request.Register) (user *system.SysUser, err error) {
	if !errors.Is(global.TASK_DB.Where("nick_name = ?", register.NickName).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("昵称重复")
	}

	user = &system.SysUser{
		NickName: register.NickName,
	}

	// 给注册用户创建一个uuid
	user.UUID = uuid.NewV4()

	// 加密密码(数据库中存明文密码不安全)
	user.Password = utils.BcryptHash(register.Password)

	err = global.TASK_DB.Create(user).Error
	return user, err
}

func (SysUserService) Login(login request.Login) (sys_user *system.SysUser, err error) {
	err = global.TASK_DB.Where("nickname = ?", login.NickName).First(&sys_user).Error

	if err == nil {
		if ok := utils.BcryptCheck(sys_user.Password, login.Password); !ok {
			return nil, errors.New("密码错误")
		}
	}

	return sys_user, err
}
