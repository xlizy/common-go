package nacos

import (
	"errors"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	config2 "github.com/xlizy/common-go/base/config"
	"github.com/xlizy/common-go/utils/common"
	"github.com/xlizy/common-go/utils/zlog"
	"gopkg.in/yaml.v3"
	"math/rand/v2"
	"net"
	"strconv"
	"strings"
	"time"
)

var _namingClient naming_client.INamingClient
var configVal map[string]string

var listenConfigs = make([]interface{}, 0)

func init() {
	BaseWebConfigVal = &BaseWebConfig{}
	BaseWebConfigVal.ServerConfig = WebServerConfig{}
}

func GetNamingClient() naming_client.INamingClient {
	return _namingClient
}

func InitNacos() {
	zlog.Info("InitNacos...")

	cfg := config2.GetNacosCfg()

	nacosAddr := cfg.Addr
	t := strings.Split(nacosAddr, ":")
	host := t[0]
	port, _ := strconv.Atoi(t[1])

	serverConfigs := make([]constant.ServerConfig, 0)
	serverConfigs = append(serverConfigs, constant.ServerConfig{
		IpAddr: host,
		Port:   uint64(port),
	})
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(cfg.Namespace),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("error"),
	)
	clientParam := vo.NacosClientParam{
		ClientConfig:  &clientConfig,
		ServerConfigs: serverConfigs,
	}

	loadRemoteConfig(clientParam)
	AddListen(config2.AppEnv)
	AddListen(*config2.AppSign)
	AddListen(config2.PriorityNetwork)
	go instanceRegister(clientParam, cfg.AppName, cfg.AvailabilityCluster, cfg.Cluster)

}

func instanceRegister(clientParam vo.NacosClientParam, appName, availabilityCluster, cluster string) {

	for {
		address := net.JoinHostPort(common.GetLocalPriorityIp(config2.PriorityNetwork.Networks), config2.BootConfig.HttpPort)
		// 3 秒超时
		conn, err := net.DialTimeout("tcp", address, 3*time.Second)
		if err != nil {
			continue
		} else {
			if conn != nil {
				_ = conn.Close()
				break
			} else {
				continue
			}
		}

	}

	namingClient, err := clients.NewNamingClient(clientParam)
	if err != nil {
		zlog.Error("连接Nacos异常:{}", err.Error())
		panic(err)
	} else {
		_namingClient = namingClient
		port, _ := strconv.Atoi(config2.BootConfig.HttpPort)
		_, _ = namingClient.RegisterInstance(vo.RegisterInstanceParam{
			Ip:          common.GetLocalPriorityIp(config2.PriorityNetwork.Networks),
			Port:        uint64(port),
			ServiceName: "http:" + appName,
			Weight:      1,
			Enable:      true,
			Healthy:     true,
			Ephemeral:   true,
			Metadata:    map[string]string{"availability-cluster": availabilityCluster},
			ClusterName: cluster,
			GroupName:   "DEFAULT_GROUP",
		})
	}
}

func loadRemoteConfig(clientParam vo.NacosClientParam) {
	cfg := config2.GetNacosCfg()

	configClient, err := clients.NewConfigClient(clientParam)

	if err != nil {
		zlog.Error("连接Nacos异常:{}", err.Error())
		panic(err)
	} else {
		if cfg.DataIds != "" {
			ids := strings.Split(cfg.DataIds, ",")
			configVal = make(map[string]string, len(ids))
			for _, id := range ids {
				if id != "" {
					conStr, _ := configClient.GetConfig(vo.ConfigParam{
						DataId: id,
						Group:  "DEFAULT_GROUP",
					})
					if conStr != "" {
						if strings.Index(conStr, "config data not exist") != 0 {
							configVal[id] = conStr
							config2.ReadConfig(conStr, &BaseWebConfigVal)
							for _, config := range listenConfigs {
								LoadConfig(config)
							}
							if BaseWebConfigVal.ServerConfig.Port != "" {
								config2.BootConfig.HttpPort = BaseWebConfigVal.ServerConfig.Port
							}
						}
					}

					configClient.ListenConfig(vo.ConfigParam{
						DataId: id,
						Group:  "DEFAULT_GROUP",
						OnChange: func(namespace, group, dataId, data string) {
							if data != "" {
								if strings.Index(data, "config data not exist") != 0 {
									configVal[dataId] = data
									config2.ReadConfig(data, &BaseWebConfigVal)
									for _, config := range listenConfigs {
										LoadConfig(config)
									}
									if BaseWebConfigVal.ServerConfig.Port != "" {
										config2.BootConfig.HttpPort = BaseWebConfigVal.ServerConfig.Port
									}
								}
							}
						},
					})
				}
			}
		}

	}
}

func AddListen(configs ...interface{}) {
	for _, config := range configs {
		listenConfigs = append(listenConfigs, config)
		LoadConfig(config)
	}
}

func LoadConfig(out interface{}) {
	for _, content := range configVal {
		if content != "" && strings.Index(content, "config data not exist") != 0 {
			yaml.Unmarshal([]byte(content), out)
		}
	}
}

func GetAppIns(serviceName string) (string, error) {
	ac := config2.GetNacosCfg().AvailabilityCluster
	clusters := make([]string, 0)
	if ac != "" {
		clusters = strings.Split(ac, ",")
	}
	nc := GetNamingClient()
	if nc == nil {
		zlog.Info("nacos namingClient is nil")
		return "", errors.New("nacos namingClient is nil")
	} else {
		instances, err := nc.SelectInstances(vo.SelectInstancesParam{
			ServiceName: serviceName,
			HealthyOnly: true,
		})
		if err != nil {
			zlog.Info("selectInstances error:{}", err.Error())
			return "", err
		} else {
			if len(instances) == 0 {
				zlog.Info("no available instance:{}", err.Error())
				return "", errors.New("no available instance")
			} else if len(instances) == 1 {
				instance := instances[0]
				return instance.Ip + ":" + strconv.Itoa(int(instance.Port)), nil
			} else {
				tmpInstances := make([]model.Instance, 0)
				for _, cluster := range clusters {
					for _, instance := range instances {
						if instance.ClusterName == cluster {
							tmpInstances = append(tmpInstances, instance)
						}
					}
					if len(tmpInstances) > 0 {
						break
					}
				}
				if len(tmpInstances) == 0 {
					for _, instance := range instances {
						tmpInstances = append(tmpInstances, instance)
					}
				}
				if len(tmpInstances) == 1 {
					instance := tmpInstances[0]
					return instance.Ip + ":" + strconv.Itoa(int(instance.Port)), nil
				}

				type weight struct {
					Min float64
					Max float64
				}

				score := 0.00
				his := 0.00
				temp := make([]weight, len(tmpInstances))

				for _, server := range tmpInstances {
					score += server.Weight
				}

				for index, server := range tmpInstances {
					temp[index] = weight{
						Min: his,
						Max: his + server.Weight/score*10000,
					}
					his = temp[index].Max
				}

				r := rand.IntN(10000)
				for index, t := range temp {
					if int(t.Min) <= r && r <= int(t.Max) {
						instance := tmpInstances[index]
						return instance.Ip + ":" + strconv.Itoa(int(instance.Port)), nil
					}
				}

				instance := tmpInstances[0]
				return instance.Ip + ":" + strconv.Itoa(int(instance.Port)), nil
			}
		}
	}
}
