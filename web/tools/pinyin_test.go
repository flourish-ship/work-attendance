package tools

import (
	"testing"

	"fmt"
)

func Test_PinYin(t *testing.T) {
	str := "中国人"
	fmt.Println(FullPinYin(str))
}
