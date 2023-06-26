package main

import (
	"fmt"
	"github.com/xlizy/common-go/enums/enabled"
	"github.com/xlizy/common-go/json"
	"time"
)

type Test struct {
	Id      int64           `json:"id"`
	Name    string          `json:"name"`
	Enable  enabled.Enabled `json:"enable" enum:"true"`
	RegTime time.Time       `json:"regTime"`
}

func main() {
	t := Test{
		Id:   1,
		Name: "xlizy",
	}
	fmt.Println(json.ToJsonStr(t))
	fmt.Println(json.ToJsonStr(t.RegTime))
	fmt.Println(json.ToJsonStr(t.RegTime.Unix()))
	fmt.Println(t.RegTime.Unix())
	fmt.Println(t.RegTime.UnixMilli())
	fmt.Println(t.RegTime.UnixMicro())
	fmt.Println(t.RegTime.UnixNano())
	b := time.Time{}
	fmt.Println("---------")
	fmt.Println(b.Unix())
	fmt.Println(b.UnixMilli())
	fmt.Println(b.UnixMicro())
	fmt.Println(b.UnixNano())
	fmt.Println(t.RegTime.Equal(b))
}
