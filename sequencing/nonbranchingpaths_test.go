package sequencing

import (
	"sort"
	"testing"
)

func TestMaximalNonBranchingPaths(t *testing.T) {
	g := GraphFromMap(map[string][]string{
		"1": []string{"2"},
		"2": []string{"3"},
		"3": []string{"4", "5"},
		"6": []string{"7"},
		"7": []string{"6"},
	}, KmerReadParser{})

	a := MaximalNonBranchingPaths(g)
	actual := paths(make([]path, len(a)))
	for i := 0; i < len(a); i++ {
		actual[i] = a[i]
	}

	expected := paths([]path{
		path{KmerRead("1"), KmerRead("2"), KmerRead("3")},
		path{KmerRead("3"), KmerRead("4")},
		path{KmerRead("3"), KmerRead("5")},
		path{KmerRead("7"), KmerRead("6"), KmerRead("7")},
	})
	sort.Sort(actual)
	sort.Sort(expected)
	if !actual.equals(expected) {
		t.Error("for", g, "expected", expected, "got", actual)
	}
}
