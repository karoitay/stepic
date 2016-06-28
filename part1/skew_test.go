package part1

import (
	"reflect"
	"testing"
)

func TestSkew(t *testing.T) {
	type testpair struct {
		input  string
		output []int
	}
	tests := []testpair{
		{"CATGGGCATCGGCCATACGCC", []int{0, -1, -1, -1, 0, 1, 2, 1, 1, 1, 0, 1, 2, 1, 0, 0, 0, 0, -1, 0, -1, -2}},
		{"GAGCCACCGCGATA", []int{0, 1, 1, 2, 1, 0, 0, -1, -2, -1, -2, -1, -1, -1, -1}},
	}
	for _, pair := range tests {
		v := Skew(pair.input)
		if !reflect.DeepEqual(v, pair.output) {
			t.Error("For", pair.input, "expected", pair.output, "got", v)
		}
	}
}

func TestMinimumSkew(t *testing.T) {
	type testpair struct {
		input  string
		output []int
	}
	tests := []testpair{
		{"TAAAGACTGCCGAGAGGCCAACACGAGTGCTAGAACGAGGGGCGTAAACGCGGGTCCGAT", []int{11, 24}},
	}
	for _, pair := range tests {
		v := MinimumSkew(pair.input)
		if !reflect.DeepEqual(v, pair.output) {
			t.Error("For", pair.input, "expected", pair.output, "got", v)
		}
	}
}
