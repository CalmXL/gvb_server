package core

import (
	"fmt"
	"gvb_server/config"
	"gvb_server/global"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// 读取配置
func InitConf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}

	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}

	fmt.Println("config yamlFile load Init success.")
	// fmt.Println(c)
	global.Config = c
}
