package request

import (
	jwt "github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
)

// 自定义标准
type CustomClaims struct {
	BaseClaims
	BufferTime           int64
	jwt.RegisteredClaims //jwt自带标准要求
}

// 基本要求
type BaseClaims struct {
	UUID     uuid.UUID
	ID       uint
	NickName string
}
