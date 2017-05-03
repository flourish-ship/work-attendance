package tools

import (
	"testing"

	"fmt"
)

func Test_PinYin(t *testing.T) {
	str := "Matao"
	fmt.Println(ToPinYin(str))
}
