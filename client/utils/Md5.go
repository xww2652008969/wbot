package utils

import (
	md52 "crypto/md5"
	"fmt"
)

func Gedmd5(date []byte) string {
	md5 := md52.Sum(date)
	md5str := fmt.Sprintf("%x", md5)
	return md5str
}
