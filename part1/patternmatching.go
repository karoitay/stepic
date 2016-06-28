package part1

import "strings"

func PatternMatching(text, pattern string) []int {
	indexFrom := func(i int) func(string, string) int {
		return func(text, pattern string) int {
			v := strings.Index(text[i:], pattern)
			if v == -1 {
				return v
			}
			return i + v
		}
	}
	var res []int
	i := strings.Index(text, pattern)
	for i != -1 {
		res = append(res, i)
		i = indexFrom(i+1)(text, pattern)
	}
	return res
}
