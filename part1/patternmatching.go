package part1

func indicesOf(text, pattern string, f stringPredicate) []int {
	var res []int
	l := len(pattern)
	for i := 0; i <= len(text)-l; i++ {
		if f(pattern, text[i:i+l]) {
			res = append(res, i)
		}
	}
	return res
}

// PatternMatching returns the indices int text that pattern can be found in.
func PatternMatching(text, pattern string) []int {
	return indicesOf(text, pattern, exactStringMatch)
}

// ApproximatePatternMatching returns the indices int text that pattern can be
// found in (with hamming distance match).
func ApproximatePatternMatching(text, pattern string, d int) []int {
	return indicesOf(text, pattern, maxHammingDistanceMatch(d))
}
