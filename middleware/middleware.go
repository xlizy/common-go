package middleware

import (
	"github.com/google/uuid"
	"github.com/kataras/iris/v12/context"
	"github.com/xlizy/common-go/const/threadlocal"
	"github.com/xlizy/common-go/response"
)

var TraceId = func(ctx *context.Context) {
	traceId := ctx.GetHeader("X-Request-Id")
	if traceId == "" {
		traceId = uuid.New().String()
	}
	threadlocal.SetTraceId(traceId)
	ctx.Next()
}

var NeedLogin = func(ctx *context.Context) {
	userId := ctx.GetHeader("X-Request-UserId")
	if userId == "" {
		ctx.JSON(response.Error(500, "未登录", nil))
	} else {
		ctx.Next()
	}
}
