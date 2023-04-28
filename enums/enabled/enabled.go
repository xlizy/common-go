package enabled

type Enabled int

const (
	FALSE Enabled = iota
	TRUE
)

func (e Enabled) Des() string {
	switch e {
	case TRUE:
		return "启用"
	case FALSE:
		return "未启用"
	}
	return "未知"
}
