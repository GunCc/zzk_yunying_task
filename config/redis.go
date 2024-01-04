package config

type Redis struct {
	Db       int    `json:"db" yaml:"db"`
	Password string `json:"password" yaml:"password"`
	Addr     string `json:"addr" yaml:"addr"`
}
