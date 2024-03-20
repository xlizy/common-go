package dubbo

import (
	"dubbo.apache.org/dubbo-go/v3/cluster/loadbalance"
	"dubbo.apache.org/dubbo-go/v3/common/extension"
	"dubbo.apache.org/dubbo-go/v3/protocol"
	"github.com/xlizy/common-go/config"
	"math/rand/v2"
	"strconv"
	"strings"
)

// ClusterWeightedRandomRobinLoadBalance 基于按集群顺序优先的甲醛随机负载均衡策略
type ClusterWeightedRandomRobinLoadBalance struct {
}

func (lb ClusterWeightedRandomRobinLoadBalance) Select(invokers []protocol.Invoker, invocation protocol.Invocation) protocol.Invoker {
	if len(invokers) == 0 {
		return nil
	}
	if len(invokers) == 1 {
		return invokers[0]
	}
	clusters := make([]string, 0)
	ac := config.GetNacosCfg().AvailabilityCluster
	if ac != "" {
		clusters = strings.Split(ac, ",")
	}
	tmpInvoker := make([]protocol.Invoker, 0)
	for _, cluster := range clusters {
		for _, invoker := range invokers {
			if invoker.GetURL().GetParam("group", "DEFAULT") == cluster {
				tmpInvoker = append(tmpInvoker, invoker)
			}
		}
		if len(tmpInvoker) > 0 {
			break
		}
	}
	if len(tmpInvoker) == 0 {
		for _, invoker := range invokers {
			tmpInvoker = append(tmpInvoker, invoker)
		}
	}
	if len(tmpInvoker) == 1 {
		return tmpInvoker[0]
	}
	type weight struct {
		Min float64
		Max float64
	}
	score := 0.00
	his := 0.00
	temp := make([]weight, len(tmpInvoker))

	for _, invoker := range tmpInvoker {
		appWeightStr := invoker.GetURL().GetParam("appWeight", "1")
		appWeight, _ := strconv.Atoi(appWeightStr)
		score += float64(appWeight)
	}

	for index, invoker := range tmpInvoker {
		appWeightStr := invoker.GetURL().GetParam("appWeight", "1")
		appWeight, _ := strconv.Atoi(appWeightStr)
		temp[index] = weight{
			Min: his,
			Max: his + float64(appWeight)/score*10000,
		}
		his = temp[index].Max
	}

	r := rand.IntN(10000)
	for index, t := range temp {
		if int(t.Min) <= r && r <= int(t.Max) {
			instance := tmpInvoker[index]
			return instance
		}
	}
	return invokers[0]
}

func init() {
	extension.SetLoadbalance("clusterWeightedRandomRobinLoadBalance", NewCustomLoadBalance)
}

func NewCustomLoadBalance() loadbalance.LoadBalance {
	return &ClusterWeightedRandomRobinLoadBalance{}
}
