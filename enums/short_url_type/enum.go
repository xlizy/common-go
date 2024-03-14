package short_url_type

import (
	"github.com/xlizy/common-go/enums"
	"strconv"
)

type ShortUrlType int

const (
	PERPETUAL ShortUrlType = iota + 1
	TIMES
	TIME_LIMIT
	TIME_LIMIT_AND_TIMES
)

func (e ShortUrlType) Val() int {
	return int(e)
}

func (e ShortUrlType) Des() string {
	return enums.BE(e).Des
}

func (e ShortUrlType) MarshalJSON() ([]byte, error) {
	switch e {
	case PERPETUAL:
		return enums.JsonObj(int(e), "PERPETUAL", "永久")
	case TIMES:
		return enums.JsonObj(int(e), "TIMES", "次数限制")
	case TIME_LIMIT:
		return enums.JsonObj(int(e), "TIME_LIMIT", "限时")
	case TIME_LIMIT_AND_TIMES:
		return enums.JsonObj(int(e), "TIME_LIMIT_AND_TIMES", "限时次数")
	}
	return []byte(strconv.Itoa(int(e))), nil
}
