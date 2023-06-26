package true_false

import (
	"github.com/xlizy/common-go/enums"
	"strconv"
)

type TrueFalse int

const (
	TRUE TrueFalse = iota + 1
	FALSE
)

func (e TrueFalse) Val() int {
	return int(e)
}

func (e TrueFalse) Des() string {
	return enums.BE(e).Des
}

func (e TrueFalse) MarshalJSON() ([]byte, error) {
	switch e {
	case FALSE:
		return enums.JsonObj(int(e), "FALSE", "false")
	case TRUE:
		return enums.JsonObj(int(e), "TRUE", "true")
	}
	return []byte(strconv.Itoa(int(e))), nil
}
