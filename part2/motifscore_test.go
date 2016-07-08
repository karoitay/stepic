package part2

import (
	"math"
	"testing"
)

func TestMotifsEntropy(t *testing.T) {
	type testpair struct {
		input  []string
		output float64
	}
	tests := []testpair{
		{
			[]string{
				"TCGGGGGTTTTT",
				"CCGGTGACTTAC",
				"ACGGGGATTTTC",
				"TTGGGGACTTTT",
				"AAGGGGACTTCC",
				"TTGGGGACTTCC",
				"TCGGGGATTCAT",
				"TCGGGGATTCCT",
				"TAGGGGAACTAC",
				"TCGGGTATAACC",
			},
			9.91629,
		},
	}
	acceptedError := 0.0000001
	for _, pair := range tests {
		v := MotifsEntropy(pair.input)
		if math.Abs(v-pair.output) > acceptedError {
			t.Error("For", pair.input, "expected", pair.output, "within error", acceptedError, "got", v)
		}
	}
}

func TestMotifsScore(t *testing.T) {
	type testpair struct {
		input  []string
		output int
	}
	tests := []testpair{
		{
			[]string{
				"TCGGGGGTTTTT",
				"CCGGTGACTTAC",
				"ACGGGGATTTTC",
				"TTGGGGACTTTT",
				"AAGGGGACTTCC",
				"TTGGGGACTTCC",
				"TCGGGGATTCAT",
				"TCGGGGATTCCT",
				"TAGGGGAACTAC",
				"TCGGGTATAACC",
			},
			30,
		},
	}
	for _, pair := range tests {
		v := MotifsScore(pair.input)
		if v != pair.output {
			t.Error("For", pair.input, "expected", pair.output, "got", v)
		}
	}
}
