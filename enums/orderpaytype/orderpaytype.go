package orderpaytype

type OrderPayType int

const (
	NOT_SELECTED OrderPayType = iota
	WECHAT
	ALIPAY
)

func (e OrderPayType) Des() string {
	switch e {
	case NOT_SELECTED:
		return "未选择"
	case WECHAT:
		return "微信支付"
	case ALIPAY:
		return "支付宝"
	}
	return "未知"
}
