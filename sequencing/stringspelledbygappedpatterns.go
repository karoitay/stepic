package sequencing

// StringSpelledByGappedPatterns returns the string spelled by the given gapped
// pattern or an empty string if none exists.
func StringSpelledByGappedPatterns(gappedPatterns []string, k, d int) string {
	firstPatterns, secondPatterns := split(gappedPatterns, k, d)
	prefixString := GenomePath(firstPatterns)
	suffixString := GenomePath(secondPatterns)
	if prefixString[k+d:] != suffixString[:len(suffixString)-k-d] {
		return ""
	}
	return prefixString[:k+d] + suffixString
}

func split(gappedPatterns []string, k, d int) ([]string, []string) {
	firstPatterns := make([]string, len(gappedPatterns))
	secondPatterns := make([]string, len(gappedPatterns))
	for i := 0; i < len(gappedPatterns); i++ {
		firstPatterns[i] = gappedPatterns[i][:k]
		secondPatterns[i] = gappedPatterns[i][k+1:]
	}
	return firstPatterns, secondPatterns
}
