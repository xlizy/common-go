package middleware

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/kataras/iris/v12/context"
	"github.com/xlizy/common-go/const/threadlocal"
	common_error "github.com/xlizy/common-go/enums/common_error"
	"github.com/xlizy/common-go/response"
	"github.com/xlizy/common-go/validator"
	"github.com/xlizy/common-go/zlog"
	"runtime"
)

var Default = func(ctx *context.Context) {
	traceId := ctx.GetHeader("X-Request-Id")
	if traceId == "" {
		traceId = uuid.New().String()
	}
	threadlocal.SetTraceId(traceId)
	userId := ctx.GetHeader("X-Request-UserId")
	if userId != "" {
		threadlocal.SetUserId(userId)
	}
	validator.TransInit(ctx)
	ctx.Next()
}

var CrossDomain = func(ctx *context.Context) {
	ctx.Header("Access-Control-Allow-Origin", ctx.Request().Header.Get("Origin"))
	ctx.Header("Access-Control-Allow-Headers", "Content-Type,X-Request-Id,X-Request-Some")
	ctx.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS,PUT,DELETE,PATCH,HEAD,TRACE")
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
		ctx.JSON(response.Error(common_error.NOT_LOGGED_IN, nil))
	} else {
		if userId != "" {
			threadlocal.SetUserId(userId)
			ctx.Next()
		} else {
			ctx.JSON(response.Error(common_error.NOT_LOGGED_IN, nil))
		}
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

			errMsg := fmt.Sprintf("%s", err)
			zlog.Info("异常Ctl入口:{}", ctx.HandlerName())
			zlog.Info("ErrorInfo:{}", errMsg)
			zlog.Info("ErrorStacktrace:{}", stacktrace)
			// 返回错误信息
			ctx.JSON(response.Error(common_error.SYSTEM_ERROR, errMsg))
			ctx.StatusCode(500)
			ctx.StopExecution()
		}
	}()
	ctx.Next()
}
