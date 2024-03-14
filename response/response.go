package response

import (
	constant "github.com/xlizy/common-go/const"
	"github.com/xlizy/common-go/const/threadlocal"
	"github.com/xlizy/common-go/enums/common_error"
	"github.com/xlizy/common-go/zlog"
	"reflect"
	"time"
)

type Response struct {
	Success      bool              `json:"success"`
	Code         int32             `json:"code"`
	Msg          string            `json:"msg,omitempty"`
	ResponseTime string            `json:"responseTime,omitempty"`
	TraceId      string            `json:"traceId,omitempty"`
	Data         any               `json:"data,omitempty"`
	Extend       map[string]string `json:"extend,omitempty"`
}

type PageResponse struct {
	Total    int64 `json:"total"`
	Data     any   `json:"data"`
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
	Pages    int   `json:"pages"`
	Size     int   `json:"size"`
}

func Success(msg string, data any) Response {
	return Response{Success: true, Code: 0, Msg: msg, ResponseTime: time.Now().Format(constant.DataFormat), TraceId: threadlocal.GetTraceId(), Data: data}
}

func SuccessCus(msg string, data any, extend map[string]string) Response {
	return Response{Success: true, Code: 0, Msg: msg, ResponseTime: time.Now().Format(constant.DataFormat), TraceId: threadlocal.GetTraceId(), Data: data, Extend: extend}
}

func ErrorCus(code int32, msg string, data any) Response {
	return Response{Success: false, Code: code, Msg: msg, ResponseTime: time.Now().Format(constant.DataFormat), TraceId: threadlocal.GetTraceId(), Data: data}
}

func Error(errType any, data any) (res Response) {
	res = Response{Success: false, Code: common_error.SYS_ERR_ENUM_ERROR.Code(), Msg: common_error.SYS_ERR_ENUM_ERROR.Des(), ResponseTime: time.Now().Format(constant.DataFormat), TraceId: threadlocal.GetTraceId(), Data: data}
	defer func(r *Response) {
		err := recover() // recover()内置函数，可以捕获到异常
		if err != nil {  //说明捕获到错误
		}
	}(&res)
	code := reflect.ValueOf(errType).MethodByName("Code").Call(nil)[0]
	des := reflect.ValueOf(errType).MethodByName("Des").Call(nil)[0]
	res.Code = int32(code.Int())
	res.Msg = des.String()
	return res

}

func RpcError(err error) Response {
	zlog.Error("调用dubbo服务异常:{}", err.Error())
	return Response{
		Success:      false,
		Code:         common_error.DUBBO_SERVICE_UNAVAILABLE.Code(),
		Msg:          common_error.DUBBO_SERVICE_UNAVAILABLE.Des(),
		ResponseTime: time.Now().Format(constant.DataFormat),
		TraceId:      threadlocal.GetTraceId(),
	}
}

func Page(total int64, data any, pageNum, pageSize, pages, size int) PageResponse {
	return PageResponse{Total: total, Data: data, PageNum: pageNum, PageSize: pageSize, Pages: pages, Size: size}
}
