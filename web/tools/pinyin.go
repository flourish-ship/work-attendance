package tools

import (
	"strings"

	"github.com/struCoder/Go-pinyin"
)

func ToPinYin(value string, args ...interface{}) string {
	style := pinyingo.STYLE_FIRST_LETTER
	segment := pinyingo.NO_SEGMENT
	sep := ""

	switch len(args) {
	case 1:
		style = args[0].(int)
	case 2:
		style = args[0].(int)
		segment = args[1].(bool)
	case 3:
		style = args[0].(int)
		segment = args[1].(bool)
		sep = args[2].(string)
	default:

	}

	return strings.Join(pinyingo.NewPy(style, segment).Convert(value), sep)
}
