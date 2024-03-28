package message_type

import (
	"github.com/xlizy/common-go/enums"
	"strconv"
)

type MessageType int

const (
	PLATFORM MessageType = iota + 1
	EVENT_NOTIFICATION
	SYSTEM_MSG
)

func (e MessageType) Val() int {
	return int(e)
}

func (e MessageType) Des() string {
	return enums.BE(e).Des
}

func (e MessageType) MarshalJSON() ([]byte, error) {
	switch e {
	case PLATFORM:
		return enums.JsonObj(int(e), "PLATFORM", "平台公告")
	case EVENT_NOTIFICATION:
		return enums.JsonObj(int(e), "EVENT_NOTIFICATION", "活动通知")
	case SYSTEM_MSG:
		return enums.JsonObj(int(e), "SYSTEM_MSG", "系统消息")
	}
	return []byte(strconv.Itoa(int(e))), nil
}

func (e *MessageType) UnmarshalJSON(value []byte) error {
	*e = MessageType(enums.UnmarshalEnum(value))
	return nil
}
