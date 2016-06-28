package part1

import "testing"

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

func TestApproximatePatternCount(t *testing.T) {
	type testinput struct {
		text, pattern string
		maxDistance   int
	}
	type testpair struct {
		input  testinput
		output int
	}
	tests := []testpair{
		{testinput{"AACAAGCTGATAAACATTTAAAGAG", "AAAAA", 2}, 11},
		{testinput{"TTTAGAGCCTTCAGAGG", "GAGG", 2}, 4},
	}
	for _, pair := range tests {
		v := ApproximatePatternCount(pair.input.text, pair.input.pattern, pair.input.maxDistance)
		if v != pair.output {
			t.Error("For", pair.input, "expected", pair.output, "got", v)
		}
	}
}
