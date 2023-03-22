package jwt

import (
	"fmt"
	gJwt "github.com/golang-jwt/jwt/v4"
	"time"
)

const (
	key = "FxH4OkoAAHRtldFPn1FuXD3Gc7qlh25p"
)

type Payload struct {
	Id     int64
	Roles  string
	Expire string
}
type PayloadClaims struct {
	Payload Payload
	gJwt.RegisteredClaims
}

func GenJwt(payload Payload) string {
	payload.Expire = time.Now().Add(24 * time.Hour).Format("2006-01-02T15:04:05-0700")
	// 创建Token结构体
	claims := gJwt.NewWithClaims(gJwt.SigningMethodHS256, PayloadClaims{
		Payload: payload,
	})
	// 调用加密方法，生成Token字符串
	signingString, err := claims.SignedString([]byte(key))
	if err != nil {
		fmt.Println(err.Error())
	}
	return signingString
}

func GetPayload(signingString string) Payload {
	// 根据Token字符串解析成Claims结构体
	claims, err := gJwt.ParseWithClaims(signingString, &PayloadClaims{}, func(token *gJwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	return claims.Claims.(*PayloadClaims).Payload
}
