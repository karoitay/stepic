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
