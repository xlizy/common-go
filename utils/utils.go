package utils

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net"
	"net/http"
	"reflect"
	"strings"
)

const (
	character1 = "1234567890poiuytrewqlkjhgfdsamnbvcxzOIUYTREWQLKJHGFDSAMNBVCXZ"
)

// GetLocalPriorityIp 获取本机IP(根据规则优先选取IP)
func GetLocalPriorityIp(priorityNetWork []string) string {
	var ips []string
	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("fail to get net interface addrs: %v", err)
		return ""
	}

	for _, address := range interfaceAddr {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	if len(priorityNetWork) > 0 {
		for _, nk := range priorityNetWork {
			for _, ip := range ips {
				if strings.HasPrefix(ip, nk) {
					return ip
				}
			}
		}
	}
	return ips[0]
}

// GetLocalIp 获取本机IP
func GetLocalIp() string {
	var ips []string
	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("fail to get net interface addrs: %v", err)
		return ""
	}

	for _, address := range interfaceAddr {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips[0]
}

// Clone 浅克隆，可以克隆任意数据类型，对指针类型子元素无法克隆
// 获取类型：如果类型是指针类型，需要使用Elem()获取对象实际类型
// 获取实际值：如果值是指针类型，需要使用Elem()获取实际数据
// 说白了，Elem()就是获取反射数据的实际类型和实际值
func Clone(src interface{}) interface{} {
	typ := reflect.TypeOf(src)
	if typ.Kind() == reflect.Ptr { //如果是指针类型
		typ = typ.Elem()               //获取源实际类型(否则为指针类型)
		dst := reflect.New(typ).Elem() //创建对象
		data := reflect.ValueOf(src)   //源数据值
		data = data.Elem()             //源数据实际值（否则为指针）
		dst.Set(data)                  //设置数据
		dst = dst.Addr()               //创建对象的地址（否则返回值）
		return dst.Interface()         //返回地址
	} else {
		dst := reflect.New(typ).Elem() //创建对象
		data := reflect.ValueOf(src)   //源数据值
		dst.Set(data)                  //设置数据
		return dst.Interface()         //返回
	}
}

// DeepClone 深度克隆，可以克隆任意数据类型
func DeepClone(src interface{}) interface{} {
	typ := reflect.TypeOf(src)
	if typ.Kind() == reflect.Ptr { //如果是指针类型
		typ = typ.Elem()                          //获取源实际类型(否则为指针类型)
		dst := reflect.New(typ).Elem()            //创建对象
		b, _ := json.Marshal(src)                 //导出json
		json.Unmarshal(b, dst.Addr().Interface()) //json序列化
		return dst.Addr().Interface()             //返回指针
	} else {
		dst := reflect.New(typ).Elem()            //创建对象
		b, _ := json.Marshal(src)                 //导出json
		json.Unmarshal(b, dst.Addr().Interface()) //json序列化
		return dst.Interface()                    //返回值
	}
}

// DeepCopy 深度拷贝
func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func GetTypeByContentType(contentType string) string {
	res := "unknow"
	if contentType != "" {
		contentType = strings.ToLower(contentType)
		if contentType == "image/gif" {
			res = "gif"
		} else if contentType == "image/jpeg" {
			res = "jpeg"
		} else if contentType == "image/pjpeg" {
			res = "pjpeg"
		} else if contentType == "image/png" {
			res = "png"
		} else if contentType == "image/svg+xml" {
			res = "svg"
		} else if contentType == "image/tiff" {
			res = "tiff"
		}
	}
	return res
}

func GenRandomStr(length int) string {
	str := ""
	for i := 0; i < length; i++ {
		i := rand.IntN(len(character1))
		str += character1[i : i+1]
	}
	return str
}

func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}

func StructToMap(obj interface{}, tagName string) map[string]interface{} {
	if tagName == "" {
		tagName = "json"
	}
	result := make(map[string]interface{})
	value := reflect.ValueOf(obj).Elem() // 获取指针的值
	typ := value.Type()                  // 获取类型信息

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)         // 获取字段信息
		tag := field.Tag.Get(tagName) // 获取标签（如果有）
		if field.Anonymous {
			for j := 0; j < value.Field(i).Type().NumField(); j++ {
				field_an := value.Field(i).Type().Field(j)
				tag_an := field_an.Tag.Get(tagName)
				if tag_an != "" && !field_an.Anonymous { // 只处理非匿名字段且有标签的情况
					key_an := field_an.Name // 默认使用字段名作为Key
					if tag_an != "-" {      // 若标签不等于-则使用标签作为Key
						key_an = tag_an
						if tagName == "gorm" {
							tmp_an := strings.Split(tag_an, ";")
							for k := range tmp_an {
								if strings.HasPrefix(tmp_an[k], "column:") {
									key_an = strings.ReplaceAll(tmp_an[k], "column:", "")
								}
							}
						}
					}
					result[key_an] = value.Field(i).Field(j).Interface() // 存入Map
				}
			}
		}
		if tag != "" && !field.Anonymous { // 只处理非匿名字段且有标签的情况
			key := field.Name // 默认使用字段名作为Key
			if tag != "-" {   // 若标签不等于-则使用标签作为Key
				key = tag
				if tagName == "gorm" {
					tmp := strings.Split(tag, ";")
					for i := range tmp {
						if strings.HasPrefix(tmp[i], "column:") {
							key = strings.ReplaceAll(tmp[i], "column:", "")
						}
					}
				}
			}
			result[key] = value.Field(i).Interface() // 存入Map
		}
	}

	return result
}

func GenSalt(length int) string {
	salt := ""
	for i := 0; i < length; i++ {
		i := rand.IntN(len(character1))
		salt += character1[i : i+1]
	}
	return salt
}
