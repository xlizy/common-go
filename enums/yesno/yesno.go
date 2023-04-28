package yesno

type YesNo int

const (
	No YesNo = iota
	Yes
)

func (e YesNo) Des() string {
	switch e {
	case Yes:
		return "是"
	case No:
		return "否"
	}
	return "未知"
}
