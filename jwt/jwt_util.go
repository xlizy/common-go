package jwt

import (
	"fmt"
	gJwt "github.com/golang-jwt/jwt/v4"
	constant "github.com/xlizy/common-go/const"
	"time"
)

const (
	defaultKey = "EcALispjpvf4JrsucfmNtOGQkSni6kDU08aYLfmRHoCXn6M93Nj3wJMSxv2H3E0TvF85oVnpDKswtuFR1R8UF7rXzRe8SMoZv93XPqZCcM0I8ZpnDLZRCKBbz9NOgpCA"
)

type Payload struct {
	IdMask string
	Roles  []string
	Expire string
}
type PayloadClaims struct {
	Payload Payload
	gJwt.RegisteredClaims
}

func GenJwt(payload Payload, key string) string {
	if key == "" {
		key = defaultKey
	}
	if payload.Expire == "" {
		payload.Expire = time.Now().Add(24 * time.Hour).Format(constant.DataFormat)
	}
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

func GetPayload(signingString, key string) Payload {
	if key == "" {
		key = defaultKey
	}
	// 根据Token字符串解析成Claims结构体
	claims, err := gJwt.ParseWithClaims(signingString, &PayloadClaims{}, func(token *gJwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	return claims.Claims.(*PayloadClaims).Payload
}
