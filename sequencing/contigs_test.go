package sequencing

import (
	"reflect"
	"sort"
	"testing"
)

func TestGetContigs(t *testing.T) {
	input := []string{
		"ATG", "ATG", "TGT", "TGG", "CAT", "GGA", "GAT", "AGA",
	}

	output := GetContigs(input)
	sort.Strings(output)

	expected := []string{
		"AGA", "ATG", "ATG", "CAT", "GAT", "TGGA", "TGT",
	}
	if !reflect.DeepEqual(expected, output) {
		t.Error("for", input, "expected", expected, "got", output)
	}
}
