package config

type Config struct {
	Zap    Zap    `yaml:"zap"  json:"zap"`
	Server Server `yaml:"server"  json:"server"`
}
