package sequencing

import (
	"reflect"
	"sort"
	"testing"
)

func TestStringComposition(t *testing.T) {
	k := 5
	text := "CAATCCAAC"
	actual := StringComposition(k, text)

	expected := []string{"AATCC", "ATCCA", "CAATC", "CCAAC", "TCCAA"}
	sort.Strings(actual)
	if !reflect.DeepEqual(expected, actual) {
		t.Error("For {k:", k, "text:", text, "} expected", expected, "got", actual)
	}
}
