package goods_type

import (
	"github.com/xlizy/common-go/enums"
	"strconv"
)

type GoodsType int

const (
	COURSE GoodsType = iota + 1
)

func (e GoodsType) Val() int {
	return int(e)
}

func (e GoodsType) Des() string {
	return enums.BE(e).Des
}

func (e GoodsType) MarshalJSON() ([]byte, error) {
	switch e {
	case COURSE:
		return enums.JsonObj(int(e), "COURSE", "课程")
	}
	return []byte(strconv.Itoa(int(e))), nil
}
