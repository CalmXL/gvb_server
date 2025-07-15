package config

type Logger struct {
	Level        string `json:"level"`
	Port         int    `json:"prefix"`
	Director     string `json:"director"`
	ShowLine     string `json:"show_line"`
	LogInConsole string `json:"log_in_console"`
}
