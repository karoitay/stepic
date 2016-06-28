package part1

import (
	"reflect"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	type testinput struct {
		text1, text2 string
	}
	type testpair struct {
		input  testinput
		output int
	}
	tests := []testpair{
		{testinput{"GGGCCGTTGGT", "GGACCGTTGAC"}, 3},
	}
	for _, pair := range tests {
		v := HammingDistance(pair.input.text1, pair.input.text2)
		if !reflect.DeepEqual(v, pair.output) {
			t.Error("For", pair.input, "expected", pair.output, "got", v)
		}
	}
}
