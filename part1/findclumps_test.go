package part1

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindClumps(t *testing.T) {
	type testinput struct {
		text    string
		k, l, t int
	}
	type testpair struct {
		input  testinput
		output []string
	}
	tests := []testpair{
		{testinput{"CGGACTCGACAGATGTGAAGAACGACAATGTGAAGACTCGACACGACAGAGTGAAGAGAAGAGGAAACATTGTAA", 5, 50, 4}, []string{"CGACA", "GAAGA"}},
	}
	for _, pair := range tests {
		v := FindClumps(pair.input.text, pair.input.k, pair.input.l, pair.input.t)
		sort.Strings(v)
		if !reflect.DeepEqual(v, pair.output) {
			t.Error("For", pair.input, "expected", pair.output, "got", v)
		}
	}
}
