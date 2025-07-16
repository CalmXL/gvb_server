package config

import "strconv"

type Mysql struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Config   string `json:"config"` // 高级配置， charset
	DB       string `json:"db"`
	User     string `json:"user"`
	Password string `json:"password"`
	LogLevel string `json:"log_level"`
}

func (m *Mysql) DSN() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?" + m.Config
}
