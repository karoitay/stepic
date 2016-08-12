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
	}, KmerReadParser{})
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
	}, KmerReadParser{})
	if !reflect.DeepEqual(expected, actual) {
		t.Error("For", kmers, "expected", expected, "got", actual)
	}
}

func TestDeBruijnFromReads(t *testing.T) {
	readPairs := []string{
		"AAT|CAT", "ATG|ATG", "ATG|ATG", "CAT|GAT", "CCA|GGA", "GCC|GGG", "GGG|GTT", "TAA|CCA", "TGC|TGG", "TGG|TGT",
	}
	actual := DeBruijnFromReads(readPairs, ReadPairParser{Gap: 2})
	for _, v := range *actual {
		sort.Sort(Nodes(v))
	}

	expected := GraphFromMap(map[string][]string{
		"AA|CA": {"AT|AT"},
		"AT|AT": {"TG|TG", "TG|TG"},
		"CA|GA": {"AT|AT"},
		"CC|GG": {"CA|GA"},
		"GC|GG": {"CC|GG"},
		"GG|GT": {"GG|TT"},
		"TA|CC": {"AA|CA"},
		"TG|TG": {"GC|GG", "GG|GT"},
	}, ReadPairParser{Gap: 3})
	if !reflect.DeepEqual(expected, actual) {
		t.Error("For", readPairs, "expected", expected, "got", actual)
	}
}
