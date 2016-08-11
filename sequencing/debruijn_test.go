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

	expected := Graph(map[Node][]Node{
		Node(Read("AAG")): []Node{Read("AGA"), Read("AGA")},
		Node(Read("AGA")): []Node{Read("GAT")},
		Node(Read("ATT")): []Node{Read("TTC")},
		Node(Read("CTA")): []Node{Read("TAA")},
		Node(Read("CTC")): []Node{Read("TCT")},
		Node(Read("GAT")): []Node{Read("ATT")},
		Node(Read("TAA")): []Node{Read("AAG")},
		Node(Read("TCT")): []Node{Read("CTA"), Read("CTC")},
		Node(Read("TTC")): []Node{Read("TCT")},
	})
	if !reflect.DeepEqual(expected, *actual) {
		t.Error("For {k:", k, "text:", text, "} expected", expected, "got", actual)
	}
}

func TestDeBruijnFromKmers(t *testing.T) {
	kmers := []string{"GAGG", "CAGG", "GGGG", "GGGA", "CAGG", "AGGG", "GGAG"}
	actual := DeBruijnFromKmers(kmers)
	for _, v := range *actual {
		sort.Sort(Nodes(v))
	}

	expected := Graph(map[Node][]Node{
		Node(Read("AGG")): []Node{Read("GGG")},
		Node(Read("CAG")): []Node{Read("AGG"), Read("AGG")},
		Node(Read("GAG")): []Node{Read("AGG")},
		Node(Read("GGA")): []Node{Read("GAG")},
		Node(Read("GGG")): []Node{Read("GGA"), Read("GGG")},
	})
	if !reflect.DeepEqual(expected, *actual) {
		t.Error("For", kmers, "expected", expected, "got", actual)
	}
}
