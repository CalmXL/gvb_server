package config

import "fmt"

type System struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Env  string `json:"env"`
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
