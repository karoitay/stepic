package part1

import "testing"

func TestMedianString(t *testing.T) {
	type testinput struct {
		dna []string
		k   int
	}
	type testpair struct {
		input  testinput
		output string
	}
	tests := []testpair{
		{
			testinput{[]string{"AAATTGACGCAT", "GACGACCACGTT", "CGTCAGCGCCTG", "GCTGAGCACCGG", "AGTTCGGGACAG"}, 3},
			"GAC",
		},
	}
	for _, pair := range tests {
		v := MedianString(pair.input.k, pair.input.dna)
		if v != pair.output {
			t.Error("For", pair.input, "expected", pair.output, "got", v)
		}
	}
}
