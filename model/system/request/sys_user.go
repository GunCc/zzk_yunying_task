package request

// 一般来说,注册和登陆的模型是不一样的,注册所要接受的数据是很多的,而登录为了防止爆破请求,一般会增加验证码验证

// 注册
type Register struct {
	NickName string `json:"nickname"`
	Password string `json:"password"`
}

// 登录
type Login struct {
	NickName string `json:"nickname"`
	Password string `json:"password"`
}
