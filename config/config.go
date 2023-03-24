package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type logger struct {
	Path string `json:"path"`
}
type nacos struct {
	Addr      string `yaml:"addr"`
	Namespace string `yaml:"namespace"`
	AppName   string `yaml:"app-name"`
	DataIds   string `yaml:"data-ids"`
}
type BoostrapConfig struct {
	HttpPort string
	Logger   logger `yaml:"logger"`
	Nacos    nacos  `yaml:"nacos"`
}

var BootConfig = &BoostrapConfig{HttpPort: "8080"}

func GetNacosCfg() nacos {
	return BootConfig.Nacos
}

func GetLogCfg() logger {
	return BootConfig.Logger
}

func init() {
	dataBytes, err := os.ReadFile("bootstrap.yml")
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return
	}
	err = yaml.Unmarshal(dataBytes, BootConfig)
	if err != nil {
		fmt.Println("解析 yaml 文件失败：", err)
		return
	}
	if BootConfig.Nacos.Namespace == "" {
		BootConfig.Nacos.Namespace = os.Getenv("NACOS_NAMESPACE")
	}
}

func ReadConfig(configStr string, out interface{}) {
	if configStr == "" {
		return
	}
	err := yaml.Unmarshal([]byte(configStr), out)
	if err != nil {
		return
	}
}
