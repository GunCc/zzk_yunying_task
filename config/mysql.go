package config

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
