package middleware

import (
	"ZZK_YUNYING_TASK/model/commen/response"
	"ZZK_YUNYING_TASK/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.Result(response.UnAuthorized, gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}

		j := utils.NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.Result(response.Forbidden, gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.Result(response.Forbidden, gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
