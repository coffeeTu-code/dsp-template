package helper_crypto

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(data string) string {
	dataSum := md5.Sum([]byte(data))
	return hex.EncodeToString(dataSum[0:])
}
