package config

type JWT struct {
	SigningKey  string `json:"signing-key" yaml:"signing-key"`   // jwt签名
	ExpiresTime string `json:"expires-time" yaml:"expires-time"` // 过期时间
	BufferTime  string `json:"buffer-time" yaml:"buffer-time"`   // 缓冲时间
	Issuer      string `json:"issuer" yaml:"issuer"`             // 签发者
}
