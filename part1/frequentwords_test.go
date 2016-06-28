package part1

import (
	"sort"
	"strings"
	"testing"
)

func TestFrequentWords(t *testing.T) {
	type testinput struct {
		text string
		k    int
	}
	type testpair struct {
		input  testinput
		output string
	}
	tests := []testpair{
		{testinput{"ACGTTGCATGTCGCATGATGCATGAGAGCT", 4}, "CATG GCAT"},
		{testinput{"CGGAGGACTCTAGGTAACGCTTATCAGGTCCATAGGACATTCA", 3}, "AGG"},
	}
	for _, pair := range tests {
		s := strings.Split(FrequentWords(pair.input.text, pair.input.k), " ")
		sort.Strings(s)
		v := strings.Join(s, " ")
		if v != pair.output {
			t.Error("For", pair.input, "expected", pair.output, "got", v)
		}
	}
}
