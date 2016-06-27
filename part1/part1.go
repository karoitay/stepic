package part1

import "strings"

var rev = map[byte]byte{
	'A': 'T',
	'C': 'G',
	'G': 'C',
	'T': 'A',
}

func PatternCount(text, pattern string) int {
	cnt := 0
	len_pat := len(pattern)
	for i := 0; i <= len(text)-len_pat; i++ {
		if text[i:i+len_pat] == pattern {
			cnt += 1
		}
	}
	return cnt
}

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

func ReverseComplement(text string) string {
	l := len(text)
	res := make([]byte, 0)
	for i := l - 1; i >= 0; i-- {
		res = append(res, rev[text[i]])
	}
	return string(res)
}

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

func FindClumps(text string, k, l, t int) []string {
	m := map[string][]int{}
	found := map[string]bool{}
	res := []string{}
	l = l - k
	for i := 0; i <= len(text)-k; i++ {
		pattern := text[i : i+k]
		if !found[pattern] {
			indexes := append(m[pattern], i)
			m[pattern] = indexes
			c := len(indexes)
			if c >= t && indexes[c-1]-indexes[c-t] <= l {
				found[pattern] = true
				res = append(res, pattern)
			}
		}
	}
	return res
}
