package utils

import (
	"fmt"
	"strings"
)

func TrimZero(v string) string {
	vals := strings.Split(v, ".")
	if len(vals) < 2 {
		return v
	}

	str := vals[1]
	i := len(str) - 1
	for ; i >= 0; i-- {
		if string(str[i]) != "0" {
			break
		}
	}

	if i <= 0 {
		return vals[0]
	} else {
		s := ""
		for j := 0; j < i+1; j++ {
			s += string(str[j])
		}
		return fmt.Sprintf("%s.%s", vals[0], s)
	}
}
