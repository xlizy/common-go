package device

type Device int

const (
	OTHER Device = iota
	WEB
	WAP
	IOS
	ANDROID
	WECHAT_BROWSER
	WECHAT_MINI_PROGRAM
)

func (e Device) Code() int32 {
	return int32(e)
}
func (e Device) Des() string {
	switch e {
	case OTHER:
		return "未知"
	case WEB:
		return "Web"
	case WAP:
		return "Wap"
	case IOS:
		return "Ios"
	case ANDROID:
		return "Android"
	case WECHAT_BROWSER:
		return "微信浏览器"
	case WECHAT_MINI_PROGRAM:
		return "微信小程序"
	}
	return "其他"
}
