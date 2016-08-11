package sequencing

import (
	"reflect"
	"sort"
	"testing"
)

func TestDeBruijnFromText(t *testing.T) {
	k := 4
	text := "AAGATTCTCTAAGA"
	actual := DeBruijnFromText(k, text)
	for _, v := range *actual {
		sort.Sort(Nodes(v))
	}

	expected := GraphFromMap(map[string][]string{
		"AAG": []string{"AGA", "AGA"},
		"AGA": []string{"GAT"},
		"ATT": []string{"TTC"},
		"CTA": []string{"TAA"},
		"CTC": []string{"TCT"},
		"GAT": []string{"ATT"},
		"TAA": []string{"AAG"},
		"TCT": []string{"CTA", "CTC"},
		"TTC": []string{"TCT"},
	}, ReadParser{})
	if !reflect.DeepEqual(expected, actual) {
		t.Error("For {k:", k, "text:", text, "} expected", expected, "got", actual)
	}
}

func TestDeBruijnFromKmers(t *testing.T) {
	kmers := []string{"GAGG", "CAGG", "GGGG", "GGGA", "CAGG", "AGGG", "GGAG"}
	actual := DeBruijnFromKmers(kmers)
	for _, v := range *actual {
		sort.Sort(Nodes(v))
	}

	expected := GraphFromMap(map[string][]string{
		"AGG": []string{"GGG"},
		"CAG": []string{"AGG", "AGG"},
		"GAG": []string{"AGG"},
		"GGA": []string{"GAG"},
		"GGG": []string{"GGA", "GGG"},
	}, ReadParser{})
	if !reflect.DeepEqual(expected, actual) {
		t.Error("For", kmers, "expected", expected, "got", actual)
	}
}
