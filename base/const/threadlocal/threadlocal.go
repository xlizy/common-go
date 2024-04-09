package threadlocal

import "github.com/timandy/routine"

var traceID = routine.NewInheritableThreadLocal()
var userID = routine.NewInheritableThreadLocal()

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

func SetUserId(userId string) {
	userID.Set(userId)
}

func GetUserId() string {
	if userID.Get() != nil {
		return userID.Get().(string)
	} else {
		return ""
	}
}
