package part1

func indexesOf(text, pattern string, f stringPredicate) []int {
	var res []int
	l := len(pattern)
	for i := 0; i <= len(text)-l; i++ {
		if f(pattern, text[i:i+l]) {
			res = append(res, i)
		}
	}
	return res
}

func PatternMatching(text, pattern string) []int {
	return indexesOf(text, pattern, exactStringMatch)
}

func ApproximatePatternMatching(text, pattern string, d int) []int {
	return indexesOf(text, pattern, maxHammingDistanceMatch(d))
}
