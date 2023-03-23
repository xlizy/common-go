package threadlocal

import "github.com/timandy/routine"

var TraceId = routine.NewInheritableThreadLocal()
