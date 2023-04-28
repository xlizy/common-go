package goodstype

type GoodsType int

const (
	COURSE GoodsType = iota + 1
)

func (e GoodsType) Des() string {
	switch e {
	case COURSE:
		return "课程"
	}
	return "未知"
}
