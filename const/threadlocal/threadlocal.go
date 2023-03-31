package threadlocal

import "github.com/timandy/routine"

var traceID = routine.NewInheritableThreadLocal()

func SetTraceId(traceId string) {
	traceID.Set(traceId)
}

func GetTraceId() string {
	if traceID.Get() != nil {
		return traceID.Get().(string)
	} else {
		return "<nil>"
	}
}
