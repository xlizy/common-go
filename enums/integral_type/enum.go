package integral_type

import (
	"github.com/xlizy/common-go/enums"
	"strconv"
)

type IntegralType int

const (
	MAIN IntegralType = iota + 1
	QING_NIAO
)

func (e IntegralType) Val() int {
	return int(e)
}

func (e IntegralType) Des() string {
	return enums.BE(e).Des
}

func (e IntegralType) MarshalJSON() ([]byte, error) {
	switch e {
	case MAIN:
		return enums.JsonObj(int(e), "MAIN", "主要")
	case QING_NIAO:
		return enums.JsonObj(int(e), "QING_NIAO", "青鸟送信")
	}
	return []byte(strconv.Itoa(int(e))), nil
}
