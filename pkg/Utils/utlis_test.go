package Utils

import (
	"fmt"
	"strconv"
	"testing"
)

func TestMad5Hex(t *testing.T) {
	var data = make(map[string]int)
	for i := 0; i < 1000000; i++ {
		result := Md5ToHex([]byte(strconv.Itoa(i)))
		data[result]++
	}
	for _, result := range data {
		if result == 2 {
			fmt.Println("test fail")
		}
	}
}
