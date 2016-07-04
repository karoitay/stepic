package part1

import (
	"reflect"
	"testing"
)

func TestProfileMostProbable(t *testing.T) {
	type testinput struct {
		text    string
		k       int
		profile map[byte][]float64
	}
	type testpair struct {
		input  testinput
		output string
	}
	tests := []testpair{
		{
			testinput{"ACCTGTTTATTGCCTAAGTTCCGAACAAACCCAATATAGCCCGAGGGCCT", 5, map[byte][]float64{
				'A': []float64{0.2, 0.2, 0.3, 0.2, 0.3},
				'C': []float64{0.4, 0.3, 0.1, 0.5, 0.1},
				'G': []float64{0.3, 0.3, 0.5, 0.2, 0.4},
				'T': []float64{0.1, 0.2, 0.1, 0.1, 0.2},
			}},
			"CCGAG",
		},
	}
	for _, pair := range tests {
		v := ProfileMostProbable(pair.input.text, pair.input.k, pair.input.profile)
		if v != pair.output {
			t.Error("For", pair.input, "expected", pair.output, "got", v)
		}
	}
}

func TestGreedyMotifSearch(t *testing.T) {
	type testinput struct {
		k   int
		dna []string
	}
	type testpair struct {
		input  testinput
		output []string
	}
	tests := []testpair{
		{
			testinput{3, []string{"GGCGTTCAGGCA", "AAGAATCAGTCA", "CAAGGAGTTCGC", "CACGTCAATCAC", "CAATAATATTCG"}},
			[]string{"CAG", "CAG", "CAA", "CAA", "CAA"},
		},
	}
	for _, pair := range tests {
		v := GreedyMotifSearch(pair.input.dna, pair.input.k)
		if !reflect.DeepEqual(v, pair.output) {
			t.Error("For", pair.input, "expected", pair.output, "got", v)
		}
	}
}

func TestPseudoCountGreedyMotifSearch(t *testing.T) {
	type testinput struct {
		k   int
		dna []string
	}
	type testpair struct {
		input  testinput
		output []string
	}
	tests := []testpair{
		{
			testinput{3, []string{"GGCGTTCAGGCA", "AAGAATCAGTCA", "CAAGGAGTTCGC", "CACGTCAATCAC", "CAATAATATTCG"}},
			[]string{"TTC", "ATC", "TTC", "ATC", "TTC"},
		},
	}
	for _, pair := range tests {
		v := PseudoCountGreedyMotifSearch(pair.input.dna, pair.input.k)
		if !reflect.DeepEqual(v, pair.output) {
			t.Error("For", pair.input, "expected", pair.output, "got", v)
		}
	}
}
