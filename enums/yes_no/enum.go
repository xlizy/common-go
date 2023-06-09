package yes_no

import (
	"github.com/xlizy/common-go/enums"
	"strconv"
)

type YesNo int

const (
	Yes YesNo = iota + 1
	No
)

func (e YesNo) Val() int {
	return int(e)
}

func (e YesNo) Des() string {
	return enums.BE(e).Des
}

func (e YesNo) MarshalJSON() ([]byte, error) {
	switch e {
	case No:
		return enums.JsonObj(int(e), "No", "否")
	case Yes:
		return enums.JsonObj(int(e), "Yes", "是")
	}
	return []byte(strconv.Itoa(int(e))), nil
}
