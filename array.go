package utils

import (
	"sort"
)

func InArrayInt(i int, list []int) bool {
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})

	index := sort.SearchInts(list, i)
	return index >= 0 && index < len(list) && list[index] == i
}

func InArrayStr(s string, list []string) bool {
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})

	index := sort.SearchStrings(list, s)
	return index >= 0 && index < len(list) && list[index] == s
}
