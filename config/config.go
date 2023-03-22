package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Logger struct {
	Path string `json:"path"`
}
type Nacos struct {
	Ip        string `yaml:"ip"`
	Port      uint64 `yaml:"port"`
	Namespace string `yaml:"namespace"`
	AppName   string `yaml:"app-name"`
	DataIds   string `yaml:"data-ids"`
}
type BoostrapConfig struct {
	Logger Logger `yaml:"logger"`
	Nacos  Nacos  `yaml:"nacos"`
}

func LoadConfig(boostrapConfig *BoostrapConfig) {
	dataBytes, err := os.ReadFile("bootstrap.yml")
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	fmt.Println("yaml 文件的内容: \n", string(dataBytes))
	err = yaml.Unmarshal(dataBytes, &boostrapConfig)
	if err != nil {
		fmt.Println("解析 yaml 文件失败：", err)
		return
	}
	if boostrapConfig.Nacos.Namespace == "" {
		boostrapConfig.Nacos.Namespace = os.Getenv("NACOS_NAMESPACE")
	}
	fmt.Printf("config → %+v\n", boostrapConfig)
}

func ReadConfig(configStr string, out interface{}) {
	err := yaml.Unmarshal([]byte(configStr), out)
	if err != nil {
		return
	}
}
