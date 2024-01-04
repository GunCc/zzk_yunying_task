package utils

import (
	"ZZK_YUNYING_TASK/global"
	"ZZK_YUNYING_TASK/model/system/request"
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")            // 过期
	TokenNotValidYet = errors.New("Token not active yet")        // 未激活
	TokenMalformed   = errors.New("That's not even a token")     // 不是token
	TokenInvalid     = errors.New("Couldn't handle this token:") // 无法处理
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.TASK_CONFIG.JWT.SigningKey),
	}
}

// 创建一个声明
func (j *JWT) CreateClaims(baseClaims request.BaseClaims) request.CustomClaims {
	// 缓冲时间
	bf, _ := ParseDuration(global.TASK_CONFIG.JWT.BufferTime)
	// 有效时间
	ep, _ := ParseDuration(global.TASK_CONFIG.JWT.ExpiresTime)

	clamis := request.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second), //缓冲时间一天
		RegisteredClaims: jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now()),         // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)), // 过期时间 7天  配置文件
			Issuer:    global.TASK_CONFIG.JWT.Issuer,          // 签名的发行者
		},
	}
	return clamis
}

// 创建一个token
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	// 根据声明的数据解析jwt
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	// token解析错误处理
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}
