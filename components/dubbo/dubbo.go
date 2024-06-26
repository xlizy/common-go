package dubbo

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/extension"
	"dubbo.apache.org/dubbo-go/v3/config"
	"dubbo.apache.org/dubbo-go/v3/filter"
	"dubbo.apache.org/dubbo-go/v3/protocol"
	"dubbo.apache.org/dubbo-go/v3/registry/directory"
	"github.com/google/uuid"
	config2 "github.com/xlizy/common-go/base/config"
	"github.com/xlizy/common-go/base/const/threadlocal"
	"github.com/xlizy/common-go/utils/common"
	"github.com/xlizy/common-go/utils/json"
	"github.com/xlizy/common-go/utils/zlog"
	"math/rand/v2"
	"net"
	"strconv"
	"time"

	_ "dubbo.apache.org/dubbo-go/v3/cluster/cluster/failover"
	_ "dubbo.apache.org/dubbo-go/v3/cluster/loadbalance/random"
	_ "dubbo.apache.org/dubbo-go/v3/config_center/nacos"
	//_ "dubbo.apache.org/dubbo-go/v3/imports"
	_ "dubbo.apache.org/dubbo-go/v3/filter/accesslog"
	_ "dubbo.apache.org/dubbo-go/v3/filter/active"
	_ "dubbo.apache.org/dubbo-go/v3/filter/adaptivesvc"
	_ "dubbo.apache.org/dubbo-go/v3/filter/auth"
	_ "dubbo.apache.org/dubbo-go/v3/filter/echo"
	_ "dubbo.apache.org/dubbo-go/v3/filter/exec_limit"
	_ "dubbo.apache.org/dubbo-go/v3/filter/generic"
	_ "dubbo.apache.org/dubbo-go/v3/filter/graceful_shutdown"
	_ "dubbo.apache.org/dubbo-go/v3/filter/hystrix"
	_ "dubbo.apache.org/dubbo-go/v3/filter/metrics"
	_ "dubbo.apache.org/dubbo-go/v3/filter/otel/trace"
	_ "dubbo.apache.org/dubbo-go/v3/filter/polaris/limit"
	_ "dubbo.apache.org/dubbo-go/v3/filter/seata"
	_ "dubbo.apache.org/dubbo-go/v3/filter/sentinel"
	_ "dubbo.apache.org/dubbo-go/v3/filter/token"
	_ "dubbo.apache.org/dubbo-go/v3/filter/tps"
	_ "dubbo.apache.org/dubbo-go/v3/filter/tps/limiter"
	_ "dubbo.apache.org/dubbo-go/v3/filter/tps/strategy"
	_ "dubbo.apache.org/dubbo-go/v3/filter/tracing"
	_ "dubbo.apache.org/dubbo-go/v3/metadata/report/nacos"
	_ "dubbo.apache.org/dubbo-go/v3/metadata/service/exporter/configurable"
	_ "dubbo.apache.org/dubbo-go/v3/metadata/service/local"
	_ "dubbo.apache.org/dubbo-go/v3/protocol/dubbo3"
	_ "dubbo.apache.org/dubbo-go/v3/protocol/dubbo3/health"
	_ "dubbo.apache.org/dubbo-go/v3/protocol/dubbo3/reflection"
	_ "dubbo.apache.org/dubbo-go/v3/proxy/proxy_factory"
	_ "dubbo.apache.org/dubbo-go/v3/registry/nacos"
	_ "dubbo.apache.org/dubbo-go/v3/registry/protocol"
)

type Service struct {
	Name          string
	InterfaceName string
	Interface     interface{}
	Weight        int
}

type Services struct {
	Consumer []Service
	Provider []Service
}

func InitDubbo(services Services) {
	nacosCfg := config2.GetNacosCfg()
	extension.SetFilter("TraceIdFilter", NewTraceIdFilter)
	rc := config.GetRootConfig()
	rc.Logger = config.NewLoggerConfigBuilder().Build()
	rc.Logger.Level = "error"
	rc.Logger.File = &config.File{
		//Name:   "./logs/dubbo.log",
		MaxAge: 30,
	}

	rc.Application.Name = config2.GetNacosCfg().AppName
	rc.Registries["nacos"] = &config.RegistryConfig{
		Protocol:     "nacos",
		Address:      nacosCfg.Addr,
		Namespace:    nacosCfg.Namespace,
		RegistryType: "interface",
	}
	extension.SetDirectory("nacos", directory.NewRegistryDirectory)
	port := ""
	for i := 0; i < 10; i++ {
		pi := rand.IntN(16383) + 49152
		port = strconv.Itoa(pi)
		zlog.Info("端口:{}", port)
		if checkPort(port) {
			break
		}
	}
	rc.Protocols["tri"] = &config.ProtocolConfig{
		Name: "tri",
		Ip:   common.GetLocalPriorityIp(config2.PriorityNetwork.Networks),
		Port: port,
	}

	check := false
	for _, service := range services.Consumer {
		config.SetConsumerService(service.Interface)
		rc.Consumer.References[service.Name] = &config.ReferenceConfig{
			InterfaceName: service.InterfaceName,
			Protocol:      "tri",
			Check:         &check,
			Retries:       "0",
			Group:         config2.GetNacosCfg().Cluster,
			Version:       "1.0.0",
			Filter:        "TraceIdFilter",
			Loadbalance:   "clusterWeightedRandomRobinLoadBalance",
		}
	}
	for _, service := range services.Provider {
		config.SetProviderService(service.Interface)
		servicesParams := make(map[string]string)
		if service.Weight <= 0 {
			servicesParams["appWeight"] = "1"
		} else {
			servicesParams["appWeight"] = strconv.Itoa(service.Weight)
		}
		rc.Provider.Services[service.Name] = &config.ServiceConfig{
			Interface:   service.InterfaceName,
			Group:       config2.GetNacosCfg().Cluster,
			Params:      servicesParams,
			Version:     "1.0.0",
			Filter:      "TraceIdFilter",
			Loadbalance: "clusterWeightedRandomRobinLoadBalance",
		}
	}

	if err := config.Load(config.WithRootConfig(rc)); err != nil {
		zlog.Error("初始化Dubbo异常:{}", err.Error())
		panic(err)
	}
}

type traceIdFilter struct {
}

func (t traceIdFilter) Invoke(ctx context.Context, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
	traceId := "<nil>"
	traceIdKey := "_dubbo_trace_id"
	filterVal, ok := invocation.GetAttachment(traceIdKey)
	if ok {
		traceId = filterVal
	} else {
		traceId = threadlocal.GetTraceId()
	}
	if traceId == "<nil>" {
		traceId = uuid.New().String()
	}
	invocation.SetAttachment(traceIdKey, traceId)
	threadlocal.SetTraceId(traceId)
	zlog.Info("dubbo-call-start,methodName:{},invocation:{}", invocation.MethodName(), json.ToJsonStr(invocation))
	return invoker.Invoke(ctx, invocation)
}

func (t traceIdFilter) OnResponse(ctx context.Context, result protocol.Result, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {
	zlog.Info("dubbo-call-end,response:{}", json.ToJsonStr(result))
	return result
}

func NewTraceIdFilter() filter.Filter {
	return &traceIdFilter{}
}

// 检测端口
func checkPort(port string) bool {
	conn, err := net.DialTimeout("tcp", "127.0.0.1:"+port, 3*time.Second)
	if err != nil {
		return true
	} else {
		if conn != nil {
			conn.Close()
			return false
		} else {
			return true
		}
	}
}
