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
	// 配置文件名
	const ConfigFile = "settings.yaml"
	// 创建配置结构体指针
	c := &config.Config{}
	// 读取配置文件的内容
	yamlConf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("get yamlConf error: %s", err))
	}

	// 将 yamlConf 字节数据反序列化到变量 c 中
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v", err)
	}

	fmt.Println("config yamlFile load Init success.")
	// 将配置赋值给全局变量
	global.Config = c
}
