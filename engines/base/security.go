package base

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"

	"github.com/astaxie/beego"
)

var aesCip cipher.Block
var hmacKey []byte

// AesEncrypt aes encrypt
func AesEncrypt(buf []byte) ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(aesCip, iv)
	val := make([]byte, len(buf))
	cfb.XORKeyStream(val, buf)

	return append(val, iv...), nil
}

// AesDecrypt aes decrypt
func AesDecrypt(buf []byte) ([]byte, error) {
	bln := len(buf)
	cln := bln - aes.BlockSize
	ct := buf[0:cln]
	iv := buf[cln:bln]

	cfb := cipher.NewCFBDecrypter(aesCip, iv)
	val := make([]byte, cln)
	cfb.XORKeyStream(val, ct)
	return val, nil
}

// HmacSum sum hmac
func HmacSum(plain []byte) []byte {
	mac := hmac.New(sha512.New, hmacKey)
	return mac.Sum(plain)
}

// HmacChk chk hmac
func HmacChk(plain, code []byte) bool {
	return hmac.Equal(HmacSum(plain), code)
}

func init() {
	var err error
	aesCip, err = aes.NewCipher([]byte(beego.AppConfig.String("aeskey")))
	if err != nil {
		beego.Error(err)
	}
	hmacKey = []byte(beego.AppConfig.String("hmackey"))
}
