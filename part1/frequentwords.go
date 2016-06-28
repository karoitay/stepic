package part1

import "strings"

func FrequentWords(text string, k int) string {
	m := map[string]int{}
	max := 0
	for i := 0; i <= len(text)-k; i++ {
		pattern := text[i : i+k]
		v := m[pattern] + 1
		m[pattern] = v
		if v > max {
			max = v
		}
	}
	var res []string
	for k, v := range m {
		if v == max {
			res = append(res, k)
		}
	}
	return strings.Join(res, " ")
}
