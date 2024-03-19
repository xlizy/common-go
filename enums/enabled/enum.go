package enabled

import (
	"github.com/xlizy/common-go/enums"
	"github.com/xlizy/common-go/models"
	"strconv"
)

type Enabled int

const (
	TRUE Enabled = iota + 1
	FALSE
)

func (e Enabled) Val() int {
	return int(e)
}

func (e Enabled) Des() string {
	return enums.BE(e).Des
}

func (e Enabled) MarshalJSON() ([]byte, error) {
	switch e {
	case TRUE:
		return enums.JsonObj(int(e), "TRUE", "启用")
	case FALSE:
		return enums.JsonObj(int(e), "FALSE", "未启用")
	}
	return []byte(strconv.Itoa(int(e))), nil
}

func (e Enabled) Enum() string {
	return enums.BE(e).Enum
}

func SelectOptions() []models.SelectOptions {
	return enums.SelectOptions([]any{TRUE, FALSE})
}
