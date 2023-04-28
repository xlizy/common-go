package messagetype

type MessageType int

const (
	PLATFORM MessageType = iota + 1
	EVENT_NOTIFICATION
	SYSTEM_MSG
)

func (e MessageType) Des() string {
	switch e {
	case PLATFORM:
		return "平台公告"
	case EVENT_NOTIFICATION:
		return "活动通知"
	case SYSTEM_MSG:
		return "系统消息"
	}
	return "默认"
}
