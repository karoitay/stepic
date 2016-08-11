package sequencing

import (
	"reflect"
	"testing"
)

func TestEulerianCycle(t *testing.T) {
	input := GraphFromMap(map[string][]string{
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
	}, ReadParser{})

	actual := EulerianCycle(input)

	expected := []Node{Read("6"), Read("8"), Read("7"), Read("9"),
		Read("6"), Read("5"), Read("4"), Read("2"), Read("1"), Read("0"),
		Read("3"), Read("2"), Read("6")}
	if !isCyclicEqual(expected, actual, t) {
		t.Error("For", input, "expected", expected, "got", actual)
	}
}

func isCyclicEqual(expected, actual []Node, t *testing.T) bool {
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

func assertCyclic(name string, buffer []Node, t *testing.T) {
	if buffer[0] != buffer[len(buffer)-1] {
		t.Error(name, "is not cyclic", buffer)
	}
}
