package crypt

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(plainText []byte) string {
	m := md5.New()
	m.Write(plainText)
	return hex.EncodeToString(m.Sum(nil))
}
