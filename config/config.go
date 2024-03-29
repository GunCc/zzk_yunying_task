package config

type Config struct {
	Zap    Zap    `yaml:"zap"  json:"zap"`
	Server Server `yaml:"server"  json:"server"`
	Mysql  Mysql  `yaml:"mysql"  json:"mysql"`
	JWT    JWT    `yaml:"jwt"  json:"jwt"`
	Redis  Redis  `yaml:"redis"  json:"redis"`
}
