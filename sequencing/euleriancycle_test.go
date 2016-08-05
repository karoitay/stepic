package sequencing

import (
	"reflect"
	"testing"
)

func TestEulerianCycle(t *testing.T) {
	input := map[string][]string{
		"0": []string{"3"},
		"1": []string{"0"},
		"2": []string{"1", "6"},
		"3": []string{"2"},
		"4": []string{"2"},
		"5": []string{"4"},
		"6": []string{"5", "8"},
		"7": []string{"9"},
		"8": []string{"7"},
		"9": []string{"6"},
	}

	actual := EulerianCycle(input)

	expected := []string{"6", "8", "7", "9", "6", "5", "4", "2", "1", "0", "3", "2", "6"}
	if !isCyclicEqual(expected, actual, t) {
		t.Error("For", input, "expected", expected, "got", actual)
	}
}

func isCyclicEqual(expected, actual []string, t *testing.T) bool {
	assertCyclic("expected", expected, t)
	assertCyclic("actual", actual, t)
	expected = expected[1:]
	actual = actual[1:]
	for i := 0; i < len(actual); i++ {
		cyclicShiftLeft(expected, 1)
		if reflect.DeepEqual(expected, actual) {
			return true
		}
	}
	return false
}

func assertCyclic(name string, buffer []string, t *testing.T) {
	if buffer[0] != buffer[len(buffer)-1] {
		t.Error(name, "is not cyclic", buffer)
	}
}
