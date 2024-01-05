package config

type JWT struct {
	SigningKey  string `json:"signingkey" yaml:"signingkey"` // jwt签名
	ExpiresTime string `json:"expires" yaml:"expires"`       // 过期时间
	BufferTime  string `json:"buffer" yaml:"buffer"`         // 缓冲时间
	Issuer      string `json:"issuer" yaml:"issuer"`         // 签发者
}
