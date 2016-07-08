package part2

import "github.com/karoitay/stepic/part1"

func distance(pattern string, dna []string) int {
	d := 0
	k := len(pattern)
	for _, text := range dna {
		min := k
		for i := 0; i <= len(text)-len(pattern); i++ {
			h := part1.HammingDistance(pattern, text[i:i+k])
			if h < min {
				min = h
			}
		}
		d += min
	}
	return d
}

func MedianString(k int, dna []string) string {
	iter := part1.NewKmersIterator(k)
	min := k + 1
	pattern := ""
	for kmer := iter.Next(); kmer != ""; kmer = iter.Next() {
		d := distance(kmer, dna)
		if d < min {
			min = d
			pattern = kmer
		}
	}
	return pattern
}
