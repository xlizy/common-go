package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

// PasswordHash 生成密码hash
func PasswordHash(pwd string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(bytes)
}

// PasswordVerify 验证密码
func PasswordVerify(pwd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err == nil
}

func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	md5Val := hex.EncodeToString(m.Sum(nil))
	return md5Val
}
