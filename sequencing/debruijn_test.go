package sequencing

import (
	"reflect"
	"sort"
	"testing"
)

func TestDeBruijn(t *testing.T) {
	k := 4
	text := "AAGATTCTCTAAGA"
	actual := DeBruijn(k, text)
	for _, v := range actual {
		sort.Strings(v)
	}

	expected := map[string][]string{
		"AAG": []string{"AGA", "AGA"},
		"AGA": []string{"GAT"},
		"ATT": []string{"TTC"},
		"CTA": []string{"TAA"},
		"CTC": []string{"TCT"},
		"GAT": []string{"ATT"},
		"TAA": []string{"AAG"},
		"TCT": []string{"CTA", "CTC"},
		"TTC": []string{"TCT"},
	}
	if !reflect.DeepEqual(expected, actual) {
		t.Error("For {k:", k, "text:", text, "} expected", expected, "got", actual)
	}
}
