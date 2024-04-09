package crypto

import (
	"crypto/aes"
	"encoding/base64"
)

// 加密
func AesEncryptECB(origDataStr, keyStr string) string {
	origData := []byte(origDataStr)
	key := []byte(keyStr)
	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	return base64.StdEncoding.EncodeToString(encrypted)
}

// 解密
func AesDecryptECB(encryptedStr, keyStr string) string {
	defer func() {
		if err := recover(); err != nil {
		}
	}()
	encrypted, _ := base64.StdEncoding.DecodeString(encryptedStr)
	key := []byte(keyStr)
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted := make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return string(decrypted[:trim])
}
func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 32)
	copy(genKey, key)
	for i := 32; i < len(key); {
		for j := 0; j < 32 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}
