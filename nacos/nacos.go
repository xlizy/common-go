package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/xlizy/common-go/config"
	"github.com/xlizy/common-go/utils"
	"reflect"
	"strings"
)

var namingClient naming_client.INamingClient
var configClient config_client.IConfigClient
var configVal map[string]string

func InitNacos(cf config.Nacos, webSerPort *int) {
	BaseWebConfigVal = &BaseWebConfig{}

	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(cf.Namespace), //当namespace是public时，此处填空字符串。
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("error"),
	)
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(cf.Ip, cf.Port, constant.WithContextPath("/nacos")),
	}

	//创建配置文件客户端
	_cc, err3 := clients.NewConfigClient(vo.NacosClientParam{
		ClientConfig:  &cc,
		ServerConfigs: sc,
	})
	if err3 != nil {
		panic(err3)
	}
	configClient = _cc
	loadConfig(cf.DataIds)
	if !reflect.DeepEqual(BaseWebConfigVal.ServerConfig, WebServerConfig{}) && BaseWebConfigVal.ServerConfig.Port != 0 {
		*webSerPort = BaseWebConfigVal.ServerConfig.Port
	}

	// 创建服务发现客户端的另一种方式 (推荐)
	_nc, err1 := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err1 != nil {
		panic(err1)
	}
	//注册服务
	_, err2 := _nc.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          utils.GetLocalIp(),
		Port:        uint64(*webSerPort),
		ServiceName: cf.AppName,
		Weight:      1,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
	})
	if err2 != nil {
		panic(err2)
	}
	namingClient = _nc

}

func loadConfig(configs string) {
	BaseWebConfigVal = &BaseWebConfig{}
	if configs != "" {
		ids := strings.Split(configs, ",")
		configVal = make(map[string]string, len(ids))
		for _, id := range ids {
			if id != "" {
				conStr, _ := configClient.GetConfig(vo.ConfigParam{
					DataId: id,
					Group:  "DEFAULT_GROUP",
				})
				if configs != "" {
					config.ReadConfig(conStr, &BaseWebConfigVal)
				}
				configVal[id] = conStr
				configClient.ListenConfig(vo.ConfigParam{
					DataId: id,
					Group:  "DEFAULT_GROUP",
					OnChange: func(namespace, group, dataId, data string) {
						configVal[id] = conStr
					},
				})
			}

		}
	}
}

func GetAllConfig() map[string]string {
	return configVal
}
