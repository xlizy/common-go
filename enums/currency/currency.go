package currency

type Currency int

const (
	CNY Currency = iota + 1
)

func (e Currency) Des() string {
	switch e {
	case CNY:
		return "人民币"
	}
	return "其他"
}
