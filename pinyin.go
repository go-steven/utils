package utils

import (
	"github.com/mozillazg/go-pinyin"
	"strings"
)

func ConvertToPinyin(s string) string {
	a := pinyin.NewArgs()
	pinyins := pinyin.LazyPinyin(s, a)
	if len(pinyins) == 0 {
		return s
	}
	return strings.Join(pinyins, "")
}
