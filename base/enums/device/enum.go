package device

import (
	"github.com/xlizy/common-go/base/enums"
	"strconv"
)

type Device int

const (
	OTHER Device = -1
	WEB   Device = iota
	WAP
	IOS
	ANDROID
	WECHAT_BROWSER
	WECHAT_MINI_PROGRAM
)

func (e Device) Val() int {
	return int(e)
}
func (e Device) Des() string {
	return enums.BE(e).Des
}
func (e Device) MarshalJSON() ([]byte, error) {
	switch e {
	case OTHER:
		return enums.JsonObj(int(e), "OTHER", "未知")
	case WEB:
		return enums.JsonObj(int(e), "WEB", "Web")
	case WAP:
		return enums.JsonObj(int(e), "WAP", "Wap")
	case IOS:
		return enums.JsonObj(int(e), "IOS", "Ios")
	case ANDROID:
		return enums.JsonObj(int(e), "ANDROID", "Android")
	case WECHAT_BROWSER:
		return enums.JsonObj(int(e), "WECHAT_BROWSER", "微信浏览器")
	case WECHAT_MINI_PROGRAM:
		return enums.JsonObj(int(e), "WECHAT_MINI_PROGRAM", "微信小程序")
	}
	return []byte(strconv.Itoa(int(e))), nil
}

func (e *Device) UnmarshalJSON(value []byte) error {
	*e = Device(enums.UnmarshalEnum(value))
	return nil
}
