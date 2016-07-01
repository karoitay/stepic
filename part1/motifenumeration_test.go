package part1

import (
	"reflect"
	"sort"
	"testing"
)

func TestMotifEnumeration(t *testing.T) {
	type testinput struct {
		dna  []string
		k, d int
	}
	type testpair struct {
		input  testinput
		output []string
	}
	tests := []testpair{
		{
			testinput{[]string{"ATTTGGC", "TGCCTTA", "CGGTATC", "GAAAATT"}, 3, 1},
			[]string{"ATA", "ATT", "GTT", "TTT"},
		},
	}
	for _, pair := range tests {
		v := MotifEnumeration(pair.input.dna, pair.input.k, pair.input.d)
		sort.Strings(v)
		if !reflect.DeepEqual(v, pair.output) {
			t.Error("For", pair.input, "expected", pair.output, "got", v)
		}
	}
}
