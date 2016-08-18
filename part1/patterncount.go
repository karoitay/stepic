package part1

func patternCount(text, pattern string, f stringPredicate) int {
	c := 0
	l := len(pattern)
	for i := 0; i <= len(text)-l; i++ {
		if f(pattern, text[i:i+l]) {
			c++
		}
	}
	return c
}

// PatternCount returns how many times pattern appears in text.
func PatternCount(text, pattern string) int {
	return patternCount(text, pattern, exactStringMatch)
}

// ApproximatePatternCount returns how many times pattern appears in text
// (using hamming distance match).
func ApproximatePatternCount(text, pattern string, d int) int {
	return patternCount(text, pattern, maxHammingDistanceMatch(d))
}
