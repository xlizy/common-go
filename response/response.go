package response

import (
	"github.com/xlizy/common-go/const/threadlocal"
	"time"
)

type Response struct {
	Success      bool   `json:"success"`
	Code         int32  `json:"code"`
	Msg          string `json:"msg,omitempty"`
	ResponseTime string `json:"responseTime,omitempty"`
	TraceId      string `json:"traceId,omitempty"`
	Data         any    `json:"data,omitempty"`
}

type PageResponse struct {
	Total    int64 `json:"total"`
	Data     []any `json:"data"`
	PageNum  int32 `json:"pageNum"`
	PageSize int32 `json:"pageSize"`
	Pages    int32 `json:"pages"`
	Size     int32 `json:"size"`
}

func Success(msg string, data any) Response {
	return Response{Success: true, Code: 0, Msg: msg, ResponseTime: time.Now().Format("2006-01-02T15:04:05-0700"), TraceId: threadlocal.GetTraceId(), Data: data}
}

func Error(code int32, msg string, data any) Response {
	return Response{Success: false, Code: code, Msg: msg, ResponseTime: time.Now().Format("2006-01-02T15:04:05-0700"), TraceId: threadlocal.GetTraceId(), Data: data}
}

func Page(total int64, data []any, pageNum, pageSize, pages, size int32) PageResponse {
	return PageResponse{Total: total, Data: data, PageNum: pageNum, PageSize: pageSize, Pages: pages, Size: size}
}
