package enums

import (
	"bytes"
	"fmt"
	"github.com/xlizy/common-go/json"
	"github.com/xlizy/common-go/models"
	"reflect"
	"strconv"
)

type BaseValEnum struct {
	Enum string `json:"enum"`
	Val  int    `json:"val"`
	Des  string `json:"des"`
}

func SelectOptions[T any](enums []T) []models.SelectOptions {
	res := make([]models.SelectOptions, 0)
	for _, enumVal := range enums {
		b := reflect.ValueOf(enumVal).MethodByName("MarshalJSON").Call(nil)[0]
		be := BaseValEnum{}
		json.ToObj(fmt.Sprintf("%s", b), &be)
		res = append(res, models.SelectOptions{Value: be.Val, Desc: be.Des})
	}
	return res
}

func JsonObj(val int, enum string, des string) ([]byte, error) {
	var buf bytes.Buffer
	str := "{\"val\":" + strconv.Itoa(val) + ",\"enum\":\"" + enum + "\",\"des\":\"" + des + "\"}"
	//_, err := buf.WriteString(json.ToJsonStr(BaseValEnum{Val: val, Enum: enum, Des: des}))
	_, err := buf.WriteString(str)
	return buf.Bytes(), err
}

func BE(val any) BaseValEnum {
	b := reflect.ValueOf(val).MethodByName("MarshalJSON").Call(nil)[0]
	be := BaseValEnum{}
	json.ToObj(fmt.Sprintf("%s", b), &be)
	return be
}
