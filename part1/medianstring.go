package part1

func distance(pattern string, dna []string) int {
	d := 0
	k := len(pattern)
	for _, text := range dna {
		min := k
		for i := 0; i <= len(text)-len(pattern); i++ {
			h := HammingDistance(pattern, text[i:i+k])
			if h < min {
				min = h
			}
		}
		d += min
	}
	return d
}

func MedianString(k int, dna []string) string {
	iter := NewKmersIterator(k)
	min := k + 1
	pattern := ""
	for kmer := iter.next(); kmer != ""; kmer = iter.next() {
		d := distance(kmer, dna)
		if d < min {
			min = d
			pattern = kmer
		}
	}
	return pattern
}
