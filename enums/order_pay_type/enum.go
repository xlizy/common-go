package order_pay_type

import (
	"github.com/xlizy/common-go/enums"
	"strconv"
)

type OrderPayType int

const (
	NOT_SELECTED OrderPayType = -1
	WECHAT       OrderPayType = iota
	ALIPAY
)

func (e OrderPayType) Code() int {
	return int(e)
}

func (e OrderPayType) Des() string {
	return enums.BE(e).Des
}

func (e OrderPayType) MarshalJSON() ([]byte, error) {
	switch e {
	case NOT_SELECTED:
		return enums.JsonObj(int(e), "NOT_SELECTED", "未选择")
	case WECHAT:
		return enums.JsonObj(int(e), "WECHAT", "微信支付")
	case ALIPAY:
		return enums.JsonObj(int(e), "ALIPAY", "支付宝")
	}
	return []byte(strconv.Itoa(int(e))), nil
}
