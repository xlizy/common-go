package json

import (
	eJson "encoding/json"
)

func ToJsonStr(obj any) string {
	if obj == nil {
		return "{}"
	}
	res, err := eJson.Marshal(&obj)
	if err != nil {
		return "{}"
	} else {
		return string(res)
	}
}

func ToObj(jsonStr string, obj any) {
	eJson.Unmarshal([]byte(jsonStr), &obj)
}
