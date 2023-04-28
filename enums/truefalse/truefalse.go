package truefalse

type TrueFalse int

const (
	FALSE TrueFalse = iota
	TRUE
)

func (e TrueFalse) Des() string {
	switch e {
	case TRUE:
		return "true"
	case FALSE:
		return "false"
	}
	return "未知"
}
