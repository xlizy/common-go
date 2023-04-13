package middleware

import (
	"github.com/google/uuid"
	"github.com/kataras/iris/v12/context"
	"github.com/xlizy/common-go/const/threadlocal"
	"github.com/xlizy/common-go/response"
	"github.com/xlizy/common-go/zlog"
)

var TraceId = func(ctx *context.Context) {
	traceId := ctx.GetHeader("X-Request-Id")
	if traceId == "" {
		traceId = uuid.New().String()
	}
	threadlocal.SetTraceId(traceId)
	ctx.Next()
}

var CrossDomain = func(ctx *context.Context) {
	ctx.Header("Access-Control-Allow-Origin", ctx.Request().Header.Get("Origin"))
	zlog.Info("head,%s", ctx.Request().Header.Get("Origin"))
	ctx.Header("Access-Control-Allow-Headers", "Content-Type,X-Request-Id")
	ctx.Header("Access-Control-Allow-Methods", "*")
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Access-Control-Max-Age", "7200")
	if ctx.Request().Method == "OPTIONS" {
		ctx.StatusCode(200)
	} else {
		ctx.Next()
	}
}

var NeedLogin = func(ctx *context.Context) {
	userId := ctx.GetHeader("X-Request-UserId")
	if userId == "" {
		ctx.JSON(response.Error(500, "未登录", nil))
	} else {
		ctx.Next()
	}
}
