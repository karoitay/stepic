package part1

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
