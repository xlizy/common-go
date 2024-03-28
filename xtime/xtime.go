package xtime

import (
	"database/sql/driver"
	"fmt"
	constant "github.com/xlizy/common-go/const"
	"strings"
	"time"
)

type Time struct {
	time.Time
}

func (t Time) MarshalJSON() ([]byte, error) {
	r := make([]byte, 0)
	if t.Time.IsZero() {
		return r, nil
	}
	return []byte("\"" + t.Time.Format(constant.DataFormat) + "\""), nil
}

func (t *Time) UnmarshalJSON(value []byte) (err error) {
	var date time.Time
	layouts := []string{"2006-01-02T15:04:05-0700", "2006-01-02 15:04:05-0700", "2006-01-02 15:04:05", "2006-01-02"}
	for _, layout := range layouts {
		date, err = time.Parse(layout, strings.Replace(string(value), "\"", "", -1))
		if err == nil {
			t.Time = date
			break
		}
	}
	return
}

func (t Time) Value() (driver.Value, error) {
	if t.Time.IsZero() {
		return nil, nil
	} else {
		return t.Time.Format("2006-01-02 15:04:05"), nil
	}
}

func (t *Time) Scan(value interface{}) (err error) {
	if value == nil {
		return nil
	}

	tm, ok := value.(time.Time)
	if !ok {
		var date time.Time
		layouts := []string{"2006-01-02T15:04:05-0700", "2006-01-02 15:04:05-0700", "2006-01-02 15:04:05", "2006-01-02"}
		for _, layout := range layouts {
			date, err = time.Parse(layout, fmt.Sprintf("%v", value))
			if err == nil {
				t.Time = date
				break
			}
		}
		return
	}
	t.Time = tm
	return nil
}
