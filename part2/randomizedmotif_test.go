package part2

import (
	"reflect"
	"testing"
)

const testSeed = int64(20)

func TestRandomizedMotifSearch(t *testing.T) {
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
			testinput{8, []string{
				"CGCCCCTCTCGGGGGTGTTCAGTAAACGGCCA",
				"GGGCGAGGTATGTGTAAGTGCCAAGGTGCCAG",
				"TAGTACCGAGACCGAAAGAAGTATACAGGCGT",
				"TAGATCAAGTTTCAGGTGCACGTCGGTGAACC",
				"AATCCACCAGCTCCACGTGCAATGTTGGCCTA",
			}},
			[]string{"TCTCGGGG", "CCAAGGTG", "TACAGGCG", "TTCAGGTG", "TCCACGTG"},
		},
	}
	for _, pair := range tests {
		v := randomizedMotifSearch(pair.input.dna, pair.input.k, testSeed)
		if !reflect.DeepEqual(v, pair.output) {
			t.Error("For", pair.input, "expected", pair.output, "got", v)
		}
	}
}
