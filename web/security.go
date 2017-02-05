package web

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
)

// Security security helper
type Security interface {
	Encrypt(buf []byte) ([]byte, error)
	Decrypt(buf []byte) ([]byte, error)
	Sum(plain []byte) []byte
	Chk(plain, code []byte) bool
}

// NewAesHmacSecurity new aes-hmac security
func NewAesHmacSecurity(ak, hk []byte) (Security, error) {
	cip, err := aes.NewCipher(ak)
	if err != nil {
		return nil, err
	}
	return &AesHmacSecurity{cip: cip, key: hk}, nil
}

// AesHmacSecurity aes-hmac
type AesHmacSecurity struct {
	cip cipher.Block
	key []byte
}

// Encrypt aes encrypt
func (p *AesHmacSecurity) Encrypt(buf []byte) ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(p.cip, iv)
	val := make([]byte, len(buf))
	cfb.XORKeyStream(val, buf)

	return append(val, iv...), nil
}

// Decrypt aes decrypt
func (p *AesHmacSecurity) Decrypt(buf []byte) ([]byte, error) {
	bln := len(buf)
	cln := bln - aes.BlockSize
	ct := buf[0:cln]
	iv := buf[cln:bln]

	cfb := cipher.NewCFBDecrypter(p.cip, iv)
	val := make([]byte, cln)
	cfb.XORKeyStream(val, ct)
	return val, nil
}

// Sum sum hmac
func (p *AesHmacSecurity) Sum(plain []byte) []byte {
	mac := hmac.New(sha512.New, p.key)
	return mac.Sum(plain)
}

// Chk chk hmac
func (p *AesHmacSecurity) Chk(plain, code []byte) bool {
	return hmac.Equal(p.Sum(plain), code)
}
