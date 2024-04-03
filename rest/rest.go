package rest

import (
	"github.com/go-resty/resty/v2"
	"github.com/xlizy/common-go/enums/common_error"
	"github.com/xlizy/common-go/response"
	"github.com/xlizy/common-go/zlog"
	"net/url"
	"time"
)

var _client *resty.Client

func init() {
	_client = resty.New()
}

func GetConn() *resty.Client {
	return _client
}

func GetJsonReq(timeout time.Duration) *resty.Request {
	return _client.SetTimeout(timeout).
		R().
		SetHeader("Content-Type", "application/json; charset=utf-8")
}

func ErrHandler(err error) response.Response {
	zlog.Error("http call error:{}", err.Error())
	ce := common_error.HTTP_CALL_ERROR
	if e1, ok1 := err.(*url.Error); ok1 {
		if e1.Timeout() {
			ce = common_error.HTTP_CALL_TIMEOUT
		}
	}
	return response.Error(ce, err.Error())
}
