package currency

import (
	"github.com/xlizy/common-go/enums"
	"strconv"
)

type Currency int

const (
	CNY Currency = iota + 1
)

func (e Currency) Val() int {
	return int(e)
}
func (e Currency) Des() string {
	return enums.BE(e).Des
}
func (e Currency) MarshalJSON() ([]byte, error) {
	switch e {
	case CNY:
		return enums.JsonObj(int(e), "CNY", "人民币")
	}
	return []byte(strconv.Itoa(int(e))), nil
}

func (e *Currency) UnmarshalJSON(value []byte) error {
	*e = Currency(enums.UnmarshalEnum(value))
	return nil
}
