package json

import (
	eJson "encoding/json"
	jsoniter "github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	"github.com/kataras/iris/v12/context"
)

var js = jsoniter.ConfigCompatibleWithStandardLibrary

func init() {
	extra.RegisterFuzzyDecoders() //开启PHP兼容模式
}

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

func ReadCtxJson(ctx *context.Context, obj any) {
	jsonStr := ""
	ctx.ReadBody(&jsonStr)
	b := []byte(jsonStr)
	if jsoniter.Valid(b) {
		js.Unmarshal(b, &obj)
	}
}
