package sequencing

import (
	"reflect"
	"testing"
)

func TestEulerianPath(t *testing.T) {
	input := GraphFromMap(map[string][]string{
		"0": []string{"2"},
		"1": []string{"3"},
		"2": []string{"1"},
		"3": []string{"0", "4"},
		"6": []string{"3", "7"},
		"7": []string{"8"},
		"8": []string{"9"},
		"9": []string{"6"},
	}, KmerReadParser{})

	actual := EulerianPath(input)

	expected := []Node{KmerRead("6"), KmerRead("7"), KmerRead("8"),
		KmerRead("9"), KmerRead("6"), KmerRead("3"), KmerRead("0"),
		KmerRead("2"), KmerRead("1"), KmerRead("3"), KmerRead("4")}
	if !reflect.DeepEqual(expected, actual) {
		t.Error("For", input, "expected", expected, "got", actual)
	}
}
