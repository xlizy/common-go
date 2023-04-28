package crypto

import (
	"fmt"
	"github.com/xlizy/common-go/config"
	constant "github.com/xlizy/common-go/const"
	"github.com/xlizy/common-go/json"
	"sort"
	"time"
)

func GenSign(obj any, app string) string {
	m := make(map[string]any)
	keys := make([]string, 0)
	jsonStr := json.ToJsonStr(obj)
	json.ToObj(jsonStr, &m)
	for k, _ := range m {
		if k != "sign" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	o := ""
	for _, k := range keys {
		o += k
		o += fmt.Sprintf("%v", m[k])
	}

	if config.AppSign.Sign[app] == "" {
		o += "W4v0s1yoyvpYADUHeNpgbrTpu6a2cgEC"
	} else {
		o += config.AppSign.Sign[app]
	}
	return Md5(o)
}

func CheckSign(obj any, app string) bool {
	m := make(map[string]any)
	keys := make([]string, 0)
	jsonStr := json.ToJsonStr(obj)
	json.ToObj(jsonStr, &m)
	for k, _ := range m {
		if k != "sign" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	o := ""
	for _, k := range keys {
		o += k
		o += fmt.Sprintf("%v", m[k])
	}

	if app == "" {
		o += "W4v0s1yoyvpYADUHeNpgbrTpu6a2cgEC"
	} else {
		o += config.AppSign.Sign[app]
	}
	sign := Md5(o)
	if m["timestamp"] != "" {
		t, e := time.Parse(constant.DataFormat, fmt.Sprintf("%v", m["timestamp"]))
		if e != nil {
			return false
		}
		if time.Now().After(t.Add(5 * time.Minute)) {
			return false
		}
	}
	if m["sign"] == sign {
		return true
	}
	return false
}
