package sequencing

import (
	"reflect"
	"testing"
)

func TestOverlapGraph(t *testing.T) {
	input := []string{"ATGCG", "GCATG", "CATGC", "AGGCA", "GGCAT"}
	actual := OverlapGraph(input)

	expected := map[string][]string{
		"CATGC": []string{"ATGCG"},
		"GCATG": []string{"CATGC"},
		"GGCAT": []string{"GCATG"},
		"AGGCA": []string{"GGCAT"},
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Error("For", input, "expected", expected, "got", actual)
	}
}
