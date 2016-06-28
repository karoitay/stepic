package part1

var others = map[string][]string{
	"A": []string{"C", "G", "T"},
	"C": []string{"A", "G", "T"},
	"G": []string{"A", "C", "T"},
	"T": []string{"A", "C", "G"},
}

func FrequentWords(text string, k int) []string {
	m := map[string]int{}
	max := 0
	res := []string{}
	for i := 0; i <= len(text)-k; i++ {
		pattern := text[i : i+k]
		v := m[pattern] + 1
		m[pattern] = v
		if v > max {
			max = v
			res = []string{pattern}
		} else if v == max {
			res = append(res, pattern)
		}
	}
	return res
}

func Neighbours(pattern string, d int) map[string]int {
	m := make(map[string]int)
	if len(pattern) == 1 {
		m[pattern] = 0
		for _, n := range others[pattern] {
			m[n] = 1
		}
		return m
	}
	firstChar := pattern[0:1]
	for p, distance := range Neighbours(pattern[1:], d) {
		m[firstChar+p] = distance
		if distance < d {
			for _, n := range others[firstChar] {
				m[n+p] = distance + 1
			}
		}
	}
	return m
}

func FrequentWordsWithMismatch(text string, k, d int) []string {
	m := map[string]int{}
	max := 0
	res := []string{}
	for i := 0; i <= len(text)-k; i++ {
		for pattern := range Neighbours(text[i:i+k], d) {
			v := m[pattern] + 1
			m[pattern] = v
			if v > max {
				max = v
				res = []string{pattern}
			} else if v == max {
				res = append(res, pattern)
			}
		}
	}
	return res
}

func FrequentWordsWithMismatchAndReverseComplenent(text string, k, d int) []string {
	m := map[string]int{}
	max := 0
	resmap := map[string]bool{}
	for i := 0; i <= len(text)-k; i++ {
		for pattern := range Neighbours(text[i:i+k], d) {
			rpattern := ReverseComplement(pattern)
			v1, v2 := m[pattern]+1, m[rpattern]+1
			m[pattern], m[rpattern] = v1, v2
			if v1 > max {
				max = v1
				resmap = map[string]bool{pattern: true}
			}
			if v2 > max {
				max = v2
				resmap = map[string]bool{rpattern: true}
			}
			if v1 == max {
				resmap[pattern] = true
			}
			if v2 == max {
				resmap[rpattern] = true
			}
		}
	}
	res := make([]string, len(resmap), len(resmap))
	i := 0
	for k := range resmap {
		res[i] = k
		i++
	}
	return res
}
