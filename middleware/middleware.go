package middleware

import (
	"github.com/google/uuid"
	"github.com/kataras/iris/v12/context"
	"github.com/xlizy/common-go/const/threadlocal"
)

var TraceId = func(ctx *context.Context) {
	traceId := uuid.New().String()
	threadlocal.TraceId.Set(traceId)
	ctx.Next()
}
