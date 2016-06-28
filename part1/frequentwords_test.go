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
		s := FrequentWords(pair.input.text, pair.input.k)
		sort.Strings(s)
		v := strings.Join(s, " ")
		if v != pair.output {
			t.Error("For", pair.input, "expected", pair.output, "got", v)
		}
	}
}

func TestNeighbours(t *testing.T) {
	type testinput struct {
		text string
		d    int
	}
	type testpair struct {
		input  testinput
		output string
	}
	tests := []testpair{
		{testinput{"ACG", 1}, "AAG ACA ACC ACG ACT AGG ATG CCG GCG TCG"},
	}
	for _, pair := range tests {
		m := Neighbours(pair.input.text, pair.input.d)
		var s []string
		for k := range m {
			s = append(s, k)
		}
		sort.Strings(s)
		v := strings.Join(s, " ")
		if v != pair.output {
			t.Error("For", pair.input, "expected", pair.output, "got", v)
		}
	}
}

func TestFrequentWordsWithMismatch(t *testing.T) {
	type testinput struct {
		text string
		k, d int
	}
	type testpair struct {
		input  testinput
		output string
	}
	tests := []testpair{
		{testinput{"ACGTTGCATGTCGCATGATGCATGAGAGCT", 4, 1}, "ATGC ATGT GATG"},
	}
	for _, pair := range tests {
		s := FrequentWordsWithMismatch(pair.input.text, pair.input.k, pair.input.d)
		sort.Strings(s)
		v := strings.Join(s, " ")
		if v != pair.output {
			t.Error("For", pair.input, "expected", pair.output, "got", v)
		}
	}
}

func TestFrequentWordsWithMismatchAndReverseComplenent(t *testing.T) {
	type testinput struct {
		text string
		k, d int
	}
	type testpair struct {
		input  testinput
		output string
	}
	tests := []testpair{
		{testinput{"ACGTTGCATGTCGCATGATGCATGAGAGCT", 4, 1}, "ACAT ATGT"},
	}
	for _, pair := range tests {
		s := FrequentWordsWithMismatchAndReverseComplenent(pair.input.text, pair.input.k, pair.input.d)
		sort.Strings(s)
		v := strings.Join(s, " ")
		if v != pair.output {
			t.Error("For", pair.input, "expected", pair.output, "got", v)
		}
	}
}
