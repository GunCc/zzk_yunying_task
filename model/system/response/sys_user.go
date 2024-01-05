package response

import "ZZK_YUNYING_TASK/model/system"

type UserLoginAfter struct {
	User  system.SysUser `json:"user"`
	Token string         `json:"token"`
}
