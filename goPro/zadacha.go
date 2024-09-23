package main

import "sort"

func Permutations(input string) []string {
	if len(input) == 1 {
		return []string{input}
	}

	var result []string
	for i := 0; i < len(input); i++ {
		remainig := input[:i] + input[i+1:]
		for _, subPerm := range Permutations(remainig) {
			result = append(result, string(input[i])+subPerm)
		}
	}
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	return result
}
