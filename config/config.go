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
	Addr                string `yaml:"addr"`
	Namespace           string `yaml:"namespace"`
	AppName             string `yaml:"app-name"`
	DataIds             string `yaml:"data-ids"`
	Cluster             string `yaml:"cluster"`
	AvailabilityCluster string `yaml:"availability-cluster"`
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
		panic(err)
	}
	err = yaml.Unmarshal(dataBytes, BootConfig)
	if err != nil {
		fmt.Println("解析 yaml 文件失败：", err)
		panic(err)
	}
	if BootConfig.Nacos.Namespace == "" {
		BootConfig.Nacos.Namespace = os.Getenv("NACOS_NAMESPACE")
	}
	if BootConfig.Nacos.Cluster == "" {
		BootConfig.Nacos.Cluster = "DEFAULT"
	}
	if BootConfig.Nacos.AvailabilityCluster == "" {
		BootConfig.Nacos.AvailabilityCluster = "DEFAULT"
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
