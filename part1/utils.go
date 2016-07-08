package part1

type stringPredicate func(text1, text2 string) bool

func exactStringMatch(text1, text2 string) bool {
	return text1 == text2
}

func maxHammingDistanceMatch(d int) stringPredicate {
	return func(text1, text2 string) bool {
		return HammingDistance(text1, text2) <= d
	}
}
