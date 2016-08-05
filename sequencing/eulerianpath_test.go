package sequencing

import (
	"reflect"
	"testing"
)

func TestEulerianPath(t *testing.T) {
	input := map[string][]string{
		"0": []string{"2"},
		"1": []string{"3"},
		"2": []string{"1"},
		"3": []string{"0", "4"},
		"6": []string{"3", "7"},
		"7": []string{"8"},
		"8": []string{"9"},
		"9": []string{"6"},
	}

	actual := EulerianPath(input)

	expected := []string{"6", "7", "8", "9", "6", "3", "0", "2", "1", "3", "4"}
	if !reflect.DeepEqual(expected, actual) {
		t.Error("For", input, "expected", expected, "got", actual)
	}
}
