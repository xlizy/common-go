package dubbo

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/common/extension"
	"dubbo.apache.org/dubbo-go/v3/config"
	"dubbo.apache.org/dubbo-go/v3/filter"
	"dubbo.apache.org/dubbo-go/v3/protocol"
	"github.com/google/uuid"
	commonConfig "github.com/xlizy/common-go/config"
	"github.com/xlizy/common-go/const/threadlocal"
	"github.com/xlizy/common-go/json"
	"github.com/xlizy/common-go/zlog"
	"math/rand"
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
}

type Services struct {
	Consumer []Service
	Provider []Service
}

func InitDubbo(services Services) {
	nacosCfg := commonConfig.GetNacosCfg()
	extension.SetFilter("TraceIdFilter", NewTraceIdFilter)
	rc := config.GetRootConfig()
	rc.Logger = config.NewLoggerConfigBuilder().Build()
	rc.Logger.ZapConfig.Level = "error"
	rc.Logger.ZapConfig.OutputPaths = []string{commonConfig.GetLogCfg().Path}

	rc.Application.Name = commonConfig.GetNacosCfg().AppName
	rc.Registries["nacos"] = &config.RegistryConfig{
		Protocol:     "nacos",
		Address:      nacosCfg.Addr,
		Namespace:    nacosCfg.Namespace,
		RegistryType: "interface",
	}
	port := ""
	for i := 0; i < 10; i++ {
		pi := rand.Intn(16383) + 49152
		port = strconv.Itoa(pi)
		zlog.Info("端口:{}", port)
		if checkPort(port) {
			break
		}
	}
	rc.Protocols["tri"] = &config.ProtocolConfig{
		Name: "tri",
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
			Group:         "DEFAULT",
			Version:       "1.0.0",
			Filter:        "TraceIdFilter",
		}
	}
	for _, service := range services.Provider {
		config.SetProviderService(service.Interface)
		rc.Provider.Services[service.Name] = &config.ServiceConfig{
			Interface: service.InterfaceName,
			Group:     "DEFAULT",
			Version:   "1.0.0",
			Filter:    "TraceIdFilter",
		}
	}

	if err := config.Load(config.WithRootConfig(rc)); err != nil {
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
