package xmq

import (
	"context"
	config2 "github.com/xlizy/common-go/base/config"
	"github.com/xlizy/common-go/base/enums/common_error"
	"github.com/xlizy/common-go/base/response"
	"github.com/xlizy/common-go/components/dubbo"
	"github.com/xlizy/common-go/utils/common"
	"github.com/xlizy/common-go/utils/zlog"
	rpcApi "github.com/xlizy/rpc-interface/pbs"
)

var MQServiceClientImpl = new(rpcApi.MQServiceClientImpl)

func GetMQConsume() dubbo.Service {
	return dubbo.Service{
		Name:          "MQServiceClientImpl",
		InterfaceName: "com.rehod.mq.MQServiceClientImpl",
		Interface:     MQServiceClientImpl,
	}
}

func Send(topic, msg string) *response.Response {
	zlog.Info("发送MQ消息,topic:{},msg:{}", topic, msg)
	rsp, err := MQServiceClientImpl.SendMQ(context.TODO(), &rpcApi.SendMQReq{
		Topic:    topic,
		Msg:      msg,
		ClientIp: common.GetLocalPriorityIp(config2.PriorityNetwork.Networks),
		AppName:  config2.GetNacosCfg().AppName,
	})
	if err != nil {
		zlog.Error("发送MQ失败:{}", err.Error())
		return response.Error(common_error.RPC_CALL_ERROR, nil)
	}
	if rsp.Success {
		zlog.Error("发送MQ成功")
		return response.Succ()
	} else {
		zlog.Error("发送MQ失败:{}", rsp.Msg)
		return response.ErrorCus(rsp.Code, rsp.Msg, nil)
	}
}
