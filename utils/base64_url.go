package utils

import (
	"encoding/base64"
	"strings"
)

func EncodeBase64Url(data []byte) string {
	str := base64.StdEncoding.EncodeToString(data)
	str = strings.Replace(str, "+", "-", -1)
	str = strings.Replace(str, "/", "_", -1)
	str = strings.Replace(str, "=", "", -1)
	return str
}

func DecodeBase64Url(str string) ([]byte, error) {
	str = strings.Replace(str, "-", "+", -1)
	str = strings.Replace(str, "_", "/", -1)
	for len(str)%4 != 0 {
		str += "="
	}
	return base64.StdEncoding.DecodeString(str)
}
