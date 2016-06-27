package part1

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestPatternCount(t *testing.T) {
	type testinput struct {
		text, pattern string
	}
	type testpair struct {
		input  testinput
		output int
	}
	tests := []testpair{
		{testinput{"GCGCG", "GCG"}, 2},
		{testinput{"G", "GCG"}, 0},
		{testinput{"", "G"}, 0},
	}
	for _, pair := range tests {
		v := PatternCount(pair.input.text, pair.input.pattern)
		if v != pair.output {
			t.Error("For", pair.input, "expected", pair.output, "got", v)
		}
	}
}

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

func TestReverseComplement(t *testing.T) {
	type testpair struct {
		input  string
		output string
	}
	tests := []testpair{
		{"AAAACCCGGT", "ACCGGGTTTT"},
		{"CCAGATC", "GATCTGG"},
	}
	for _, pair := range tests {
		v := ReverseComplement(pair.input)
		if v != pair.output {
			t.Error("For", pair.input, "expected", pair.output, "got", v)
		}
	}
}

func TestPatternMatching(t *testing.T) {
	type testinput struct {
		text, pattern string
	}
	type testpair struct {
		input  testinput
		output []int
	}
	tests := []testpair{
		{testinput{"GATATATGCATATACTT", "ATAT"}, []int{1, 3, 9}},
		{testinput{"A", "A"}, []int{0}},
	}
	for _, pair := range tests {
		v := PatternMatching(pair.input.text, pair.input.pattern)
		if !reflect.DeepEqual(v, pair.output) {
			t.Error("For", pair.input, "expected", pair.output, "got", v)
		}
	}
}

func TestFindClumps(t *testing.T) {
	type testinput struct {
		text    string
		k, l, t int
	}
	type testpair struct {
		input  testinput
		output []string
	}
	tests := []testpair{
		{testinput{"CGGACTCGACAGATGTGAAGAACGACAATGTGAAGACTCGACACGACAGAGTGAAGAGAAGAGGAAACATTGTAA", 5, 50, 4}, []string{"CGACA", "GAAGA"}},
	}
	for _, pair := range tests {
		v := FindClumps(pair.input.text, pair.input.k, pair.input.l, pair.input.t)
		sort.Strings(v)
		if !reflect.DeepEqual(v, pair.output) {
			t.Error("For", pair.input, "expected", pair.output, "got", v)
		}
	}
}
