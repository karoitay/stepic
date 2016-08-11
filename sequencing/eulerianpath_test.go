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
	}, ReadParser{})

	actual := EulerianPath(input)

	expected := []Node{Read("6"), Read("7"), Read("8"), Read("9"), Read("6"),
		Read("3"), Read("0"), Read("2"), Read("1"), Read("3"), Read("4")}
	if !reflect.DeepEqual(expected, actual) {
		t.Error("For", input, "expected", expected, "got", actual)
	}
}
