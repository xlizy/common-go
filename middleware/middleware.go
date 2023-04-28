package middleware

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/kataras/iris/v12/context"
	"github.com/xlizy/common-go/const/threadlocal"
	common_error_type "github.com/xlizy/common-go/enums/commonerrortype"
	"github.com/xlizy/common-go/response"
	"github.com/xlizy/common-go/zlog"
	"runtime"
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
	ctx.Header("Access-Control-Allow-Headers", "Content-Type,X-Request-Id,X-Request-Some")
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
		ctx.JSON(response.Error(common_error_type.NOT_LOGGED_IN, nil))
	} else {
		ctx.Next()
	}
}

var GobalRecover = func(ctx *context.Context) {
	defer func() {
		if err := recover(); err != nil {
			if ctx.IsStopped() {
				return
			}

			var stacktrace string
			for i := 1; ; i++ {
				_, f, l, got := runtime.Caller(i)
				if !got {
					break
				}
				stacktrace += fmt.Sprintf("%s:%d\n", f, l)
			}

			errMsg := fmt.Sprintf("错误信息: %s", err)
			// when stack finishes
			logMessage := fmt.Sprintf("从错误中回复：('%s')\n", ctx.HandlerName())
			logMessage += errMsg + "\n"
			logMessage += fmt.Sprintf("\n%s", stacktrace)
			// 打印错误日志
			zlog.Error(logMessage)
			// 返回错误信息
			ctx.JSON(response.Error(common_error_type.SYSTEM_ERROR, logMessage))
			ctx.StatusCode(500)
			ctx.StopExecution()
		}
	}()
	ctx.Next()
}
