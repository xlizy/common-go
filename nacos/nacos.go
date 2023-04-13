package nacos

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/xlizy/common-go/config"
	"github.com/xlizy/common-go/json"
	"github.com/xlizy/common-go/utils"
	"github.com/xlizy/common-go/zlog"
	"gopkg.in/yaml.v3"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var baseHost = ""
var configVal map[string]string
var listenConfigs = make([]interface{}, 0)

func init() {
	BaseWebConfigVal = &BaseWebConfig{}
	BaseWebConfigVal.ServerConfig = WebServerConfig{}
}
func InitNacos() {
	cfg := config.GetNacosCfg()
	baseHost = "http://" + cfg.Addr
	loadRemoteConfigLight()
	go instanceRegister()
}

// 注册服务实例
func instanceRegister() {
	cfg := config.GetNacosCfg()

	for {
		address := net.JoinHostPort(utils.GetLocalIp(), config.BootConfig.HttpPort)
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

	data := url.Values{
		"port":        {config.BootConfig.HttpPort},
		"ip":          {utils.GetLocalIp()},
		"healthy":     {"true"},
		"weight":      {"1.0"},
		"serviceName": {appName(cfg.AppName)},
		"encoding":    {"GBK"},
		"namespaceId": {cfg.Namespace},
		"clusterName": {cfg.Cluster},
		"metadata":    {json.ToJsonStr(map[string]string{"availability-cluster": cfg.AvailabilityCluster})},
	}
	targetUrl := baseHost + "/nacos/v1/ns/instance"
	http.PostForm(targetUrl, data)

	go func() {
		for {
			ip := utils.GetLocalIp()
			port := config.BootConfig.HttpPort
			beat := make(map[string]interface{})
			beat["cluster"] = cfg.Cluster
			beat["ip"] = ip
			beat["port"] = port
			beat["metadata"] = map[string]string{"availability-cluster": cfg.AvailabilityCluster}
			beat["scheduled"] = true
			beat["serviceName"] = appName(cfg.AppName)
			beat["weight"] = 1

			beatStr := json.ToJsonStr(beat)

			params := url.Values{}
			params.Set("serviceName", appName(cfg.AppName))
			params.Set("ip", ip)
			params.Set("port", port)
			params.Set("namespaceId", cfg.Namespace)
			params.Set("beat", string(beatStr))
			beatTargetUrl := baseHost + "/nacos/v1/ns/instance/beat"
			uri, _ := url.ParseRequestURI(beatTargetUrl)
			uri.RawQuery = params.Encode()
			req, _ := http.NewRequest("PUT", uri.String(), nil)
			http.DefaultClient.Do(req)
			time.Sleep(15 * time.Second)
		}
	}()
}

func loadRemoteConfigLight() {

	cfg := config.GetNacosCfg()

	if cfg.DataIds != "" {
		ids := strings.Split(cfg.DataIds, ",")
		configVal = make(map[string]string, len(ids))
		for _, id := range ids {
			if id != "" {
				getRemoteConfigContent(cfg.Namespace, id)
				go listenConfig(cfg.Namespace, id)
			}
		}
	}
	if BaseWebConfigVal.ServerConfig.Port != "" {
		config.BootConfig.HttpPort = BaseWebConfigVal.ServerConfig.Port
	}
}

func getRemoteConfigContent(namespace, dataId string) string {
	var conStr = ""
	targetUrl := baseHost + "/nacos/v1/cs/configs"
	params := url.Values{}
	params.Set("tenant", namespace)
	params.Set("dataId", dataId)
	params.Set("group", "DEFAULT_GROUP")
	uri, _ := url.ParseRequestURI(targetUrl)
	uri.RawQuery = params.Encode()
	resp, _ := http.Get(uri.String())
	if resp != nil {
		b, _ := io.ReadAll(resp.Body)
		conStr = string(b)
		resp.Body.Close()
		if conStr != "" {
			if strings.Index(conStr, "config data not exist") != 0 {
				configVal[dataId] = conStr
				config.ReadConfig(conStr, &BaseWebConfigVal)
				for _, config := range listenConfigs {
					LoadConfig(config)
				}
			}
		}
	}
	return conStr
}

func listenConfig(namespace, dataId string) {
	for {
		targetUrl := baseHost + "/nacos/v1/cs/configs/listener"
		localContent := configVal[dataId]
		m := md5.New()
		m.Write([]byte(localContent))
		md5Val := hex.EncodeToString(m.Sum(nil))
		dataStr := ""
		dataStr += dataId + "%02"
		dataStr += "DEFAULT_GROUP%02"
		if localContent != "" {
			dataStr += md5Val + "%02"
		} else {
			dataStr += "" + "%02"
		}
		dataStr += namespace + "%01"
		params := url.Values{}
		params.Set("Listening-Configs", dataStr)
		uri, _ := url.ParseRequestURI(targetUrl)
		uri.RawQuery = params.Encode()
		req, _ := http.NewRequest("POST", uri.String(), nil)
		req.Header.Set("Long-Pulling-Timeout", "30000")
		resp, _ := http.DefaultClient.Do(req)
		if resp != nil {
			b, _ := io.ReadAll(resp.Body)
			res := string(b)
			if res != "" {
				//这种格式说明配置变了
				if res != dataStr {
					zlog.Info("nacos监听到namespace:%s,dataId:%s配置发生变化", namespace, dataId)
					getRemoteConfigContent(namespace, dataId)
				} else {
					//这种说明配置中心可能还没有这个配置文件
					time.Sleep(30 * time.Second)
				}
			}
			resp.Body.Close()
		}
	}
}

func AddListen(configs ...interface{}) {
	for _, config := range configs {
		listenConfigs = append(listenConfigs, config)
		LoadConfig(config)
	}
}

func GetAllConfig() map[string]string {
	return configVal
}

func LoadConfig(out interface{}) {
	for _, content := range configVal {
		if content != "" && strings.Index(content, "config data not exist") != 0 {
			yaml.Unmarshal([]byte(content), out)
		}
	}
}

func appName(name string) string {
	return "http:" + name
}
