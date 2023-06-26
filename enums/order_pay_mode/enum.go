package order_pay_mode

import (
	"github.com/xlizy/common-go/enums"
	"strconv"
)

type OrderPayMode int

const (
	DEFAULT  OrderPayMode = -1
	REDIRECT OrderPayMode = iota
	FORM
)

func (e OrderPayMode) Val() int {
	return int(e)
}

func (e OrderPayMode) Des() string {
	return enums.BE(e).Des
}

func (e OrderPayMode) MarshalJSON() ([]byte, error) {
	switch e {
	case DEFAULT:
		return enums.JsonObj(int(e), "DEFAULT", "未知")
	case REDIRECT:
		return enums.JsonObj(int(e), "REDIRECT", "跳转")
	case FORM:
		return enums.JsonObj(int(e), "FORM", "表单")
	}
	return []byte(strconv.Itoa(int(e))), nil
}
