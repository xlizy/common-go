package response

import (
	"github.com/xlizy/common-go/utils"
	rpc_api "github.com/xlizy/rpc-interface/pbs"
	"reflect"
)

func RpcRsp(rsp, data any) Response {
	rpcResValue := reflect.ValueOf(rsp).Elem().FieldByName("Result")
	rpcRes := rpcResValue.Interface().(*rpc_api.Result)
	res := Response{
		Success:      rpcRes.Success,
		Code:         rpcRes.Code,
		Msg:          rpcRes.Msg,
		ResponseTime: rpcRes.ResponseTime,
		TraceId:      rpcRes.TraceId,
	}
	rpcDataValue := reflect.ValueOf(rsp).Elem().FieldByName("Data")
	if !rpcDataValue.IsNil() {
		utils.DeepCopy(data, rpcDataValue.Interface())
	}
	res.Data = data
	return res
}
