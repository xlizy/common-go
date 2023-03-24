package nacos

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/xlizy/common-go/config"
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
		"serviceName": {cfg.AppName},
		"encoding":    {"GBK"},
		"namespaceId": {cfg.Namespace},
	}
	targetUrl := baseHost + "/nacos/v1/ns/instance"
	http.PostForm(targetUrl, data)

	go func() {
		for {
			ip := utils.GetLocalIp()
			port := config.BootConfig.HttpPort
			beat := make(map[string]interface{})
			beat["cluster"] = "DEFAULT"
			beat["ip"] = ip
			beat["port"] = port
			beat["metadata"] = make(map[string]interface{})
			beat["scheduled"] = true
			beat["serviceName"] = cfg.AppName
			beat["weight"] = 1

			beatStr, _ := json.Marshal(&beat)

			params := url.Values{}
			params.Set("serviceName", cfg.AppName)
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
	b, _ := io.ReadAll(resp.Body)
	conStr = string(b)
	resp.Body.Close()
	if conStr != "" {
		configVal[dataId] = conStr
		config.ReadConfig(conStr, &BaseWebConfigVal)
	}
	return conStr
}

func listenConfig(namespace, dataId string) {
	for {
		targetUrl := baseHost + "/nacos/v1/cs/configs/listener"
		m := md5.New()
		m.Write([]byte(configVal[dataId]))
		md5Val := hex.EncodeToString(m.Sum(nil))
		dataStr := ""
		dataStr += dataId + "%02"
		dataStr += "DEFAULT_GROUP%02"
		dataStr += md5Val + "%02"
		dataStr += namespace + "%01"
		params := url.Values{}
		params.Set("Listening-Configs", dataStr)
		uri, _ := url.ParseRequestURI(targetUrl)
		uri.RawQuery = params.Encode()
		req, _ := http.NewRequest("POST", uri.String(), nil)
		req.Header.Set("Long-Pulling-Timeout", "30000")
		resp, _ := http.DefaultClient.Do(req)
		b, _ := io.ReadAll(resp.Body)
		if string(b) != "" {
			zlog.Info("nacos监听到namespace:%s,dataId:%s配置发生变化", namespace, dataId)
			getRemoteConfigContent(namespace, dataId)
		}
		resp.Body.Close()
	}
}

func GetAllConfig() map[string]string {
	return configVal
}

func LoadConfig(out interface{}) {
	for _, content := range configVal {
		yaml.Unmarshal([]byte(content), out)
	}
}
