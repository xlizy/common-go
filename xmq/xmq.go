package xmq

import (
	"context"
	"github.com/xlizy/common-go/config"
	"github.com/xlizy/common-go/const/threadlocal"
	"github.com/xlizy/common-go/dubbo"
	"github.com/xlizy/common-go/response"
	"github.com/xlizy/common-go/utils"
	"github.com/xlizy/common-go/zlog"
	rpc_api "github.com/xlizy/rpc-interface/pbs"
)

var MQServiceClientImpl = new(rpc_api.MQServiceClientImpl)

func GetMQConsume() dubbo.Service {
	return dubbo.Service{
		Name:          "MQServiceClientImpl",
		InterfaceName: "com.rehod.mq.MQServiceClientImpl",
		Interface:     MQServiceClientImpl,
	}
}

func Send(topic, msg string) response.Response {
	zlog.Info("发送MQ消息,topic:{},msg:{}", topic, msg)
	rsp, err := MQServiceClientImpl.Send(context.TODO(), &rpc_api.SendReq{
		TraceId:  threadlocal.GetTraceId(),
		Topic:    topic,
		Msg:      msg,
		ClientIp: utils.GetLocalPriorityIp(config.PriorityNetwork.Networks),
		AppName:  config.GetNacosCfg().AppName,
	})
	if err != nil {
		zlog.Error("发送MQ失败:{}", err.Error())
		return response.Error("", nil)
	}
	if rsp.Success {
		zlog.Error("发送MQ成功")
		return response.Success("", nil)
	} else {
		zlog.Error("发送MQ失败:{}", rsp.Msg)
		return response.ErrorCus(rsp.Code, rsp.Msg, nil)
	}
}