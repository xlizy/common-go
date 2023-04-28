package orderpaymode

type OrderPayMode int

const (
	DEFAULT OrderPayMode = iota
	REDIRECT
	FORM
)

func (e OrderPayMode) Des() string {
	switch e {
	case DEFAULT:
		return "未知"
	case REDIRECT:
		return "跳转"
	case FORM:
		return "表单"
	}
	return "未知"
}
