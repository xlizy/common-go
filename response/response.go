package response

import (
	"github.com/kataras/iris/v12"
	"time"
)

const (
	TrafficKey = "X-Request-Id"
	Success    = 0
)

var Default = &response{}

func base(c iris.Context) Responses {
	res := Default.Clone()
	res.SetTraceID(GenerateMsgIDFromContext(c))
	res.SetResponseTime(time.Now().Format("2006-01-02T15:04:05-0700"))
	return res
}

// Error 失败数据处理
func Error(c iris.Context, code int, msg string, err error) {
	res := base(c)
	res.SetSuccess(false)
	res.SetCode(int32(code))
	if err != nil {
		res.SetMsg(err.Error())
	}
	if msg != "" {
		res.SetMsg(msg)
	} else {
		res.SetMsg("处理异常")
	}
	c.JSON(res)
}

// OK 通常成功数据处理
func OK(c iris.Context, data interface{}, msg string) {
	res := base(c)
	res.SetSuccess(true)
	res.SetCode(Success)
	if msg != "" {
		res.SetMsg(msg)
	} else {
		res.SetMsg("处理成功")
	}
	res.SetData(data)
	c.JSON(res)
}

// PageOK 分页数据处理
func PageOK(c iris.Context, result interface{}, count int, pageIndex int, pageSize int, msg string) {
	var res page
	res.List = result
	res.Count = count
	res.PageIndex = pageIndex
	res.PageSize = pageSize
	OK(c, res, msg)
}

func GenerateMsgIDFromContext(c iris.Context) string {
	return c.Values().GetString("__trace_id")
}
