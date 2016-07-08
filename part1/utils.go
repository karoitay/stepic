package part1

import "strings"

type stringPredicate func(text1, text2 string) bool

func exactStringMatch(text1, text2 string) bool {
	return text1 == text2
}

func maxHammingDistanceMatch(d int) stringPredicate {
	return func(text1, text2 string) bool {
		return HammingDistance(text1, text2) <= d
	}
}

type kmersIterator struct {
	k          int
	last, prev string
}

func NewKmersIterator(k int) *kmersIterator {
	iter := kmersIterator{k: k, last: "", prev: ""}
	for i := 0; i < k; i++ {
		iter.last += "T"
	}
	return &iter
}

func (iter *kmersIterator) Next() string {
	if iter.prev == iter.last {
		return ""
	} else if iter.prev == "" {
		for i := 0; i < iter.k; i++ {
			iter.prev += "A"
		}
		return iter.prev
	} else {
		i := strings.LastIndexAny(iter.prev, "ACG")
		// i must me => 0 since if all are 'T' we would have returned on the first if
		c := iter.prev[i]
		iter.prev = iter.prev[:i]
		if c == 'A' {
			iter.prev += "C"
		} else if c == 'C' {
			iter.prev += "G"
		} else if c == 'G' {
			iter.prev += "T"
		}
		for len(iter.prev) < iter.k {
			iter.prev += "A"
		}
		return iter.prev
	}
}
