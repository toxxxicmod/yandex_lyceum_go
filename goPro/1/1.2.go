package main

import (
	"sort"
	"strings"
)

func AreAnagrams(str1, str2 string) bool {

	str1 = strings.ToLower(strings.ReplaceAll(str1, " ", ""))
	str2 = strings.ToLower(strings.ReplaceAll(str2, " ", ""))

	if len(str1) != len(str2) {
		return false
	}

	return sortString(str1) == sortString(str2)
}

func sortString(s string) string {
	slice := []rune(s)
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return string(slice)
}
