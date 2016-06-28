package part1

import "testing"

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
