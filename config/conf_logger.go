package config

type Logger struct {
	Level        string `json:"level" yaml:"level"`
	Prefix       string `json:"prefix" yaml:"prefix"`
	Director     string `json:"director" yaml:"director"`
	ShowLine     bool   `json:"show_line" yaml:"show_line"`
	LogInConsole bool   `json:"log_in_console" yaml:"log_in_console"`
}
