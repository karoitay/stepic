package part1

import (
	"reflect"
	"testing"
)

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

func TestApproximatePatternMatching(t *testing.T) {
	type testinput struct {
		text, pattern string
		maxDistance   int
	}
	type testpair struct {
		input  testinput
		output []int
	}
	tests := []testpair{
		{testinput{"CGCCCGAATCCAGAACGCATTCCCATATTTCGGGACCACTGGCCTCCACGGTACGGACGTCAATCAAAT", "ATTCTGGA", 3}, []int{6, 7, 26, 27}},
	}
	for _, pair := range tests {
		v := ApproximatePatternMatching(pair.input.text, pair.input.pattern, pair.input.maxDistance)
		if !reflect.DeepEqual(v, pair.output) {
			t.Error("For", pair.input, "expected", pair.output, "got", v)
		}
	}
}
