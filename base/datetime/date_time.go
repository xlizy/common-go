package datetime

import (
	"strings"
	"time"
)

type DateTime struct {
	time.Time
}

func (t *DateTime) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte("null"), nil
	}
	return t.Time.MarshalJSON()
}

func (t *DateTime) UnmarshalJSON(b []byte) (err error) {
	var date time.Time
	layouts := []string{"2006-01-02T15:04:05-0700", "2006-01-02 15:04:05-0700", "2006-01-02 15:04:05", "2006-01-02"}
	for _, layout := range layouts {
		date, err = time.Parse(layout, strings.Replace(string(b), "\"", "", -1))
		if err == nil {
			t.Time = date
			break
		}
	}
	return
}
