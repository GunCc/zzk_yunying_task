package config

import "fmt"

type Mysql struct {
	Path     string `mapstructure:"path" json:"path" yaml:"path"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Config   string `mapstructure:"config" json:"config" yaml:"config"`
	DbName   string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.DbName + "?" + m.Config
}

// 数据库不存在时。获取的Dsn
func (i *Mysql) MysqlEmptyDsn() string {
	if i.Path == "" {
		i.Path = "127.0.0.1"
	}

	if i.Port == "" {
		i.Port = "3306"
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", i.Username, i.Password, i.Path, i.Port)
}
