package part1

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
