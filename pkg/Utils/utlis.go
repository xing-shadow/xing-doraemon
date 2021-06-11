package Utils

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"strconv"
)

func Md5ToHex(input []byte) (result string) {
	data := md5.New()
	data.Write(input)
	return hex.EncodeToString(data.Sum(nil))
}

func MustToInt(src string) int {
	result, _ := strconv.Atoi(src)
	return result
}

func IsFileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
