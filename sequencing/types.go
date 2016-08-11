package sequencing

import "strings"

// Node is a node in a graph.
type Node interface {
	ToString() string
}

// Nodes implements sort.Interface for []Node based on the node's ToString()
// value
type Nodes []Node

func (a Nodes) Len() int           { return len(a) }
func (a Nodes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Nodes) Less(i, j int) bool { return a[i].ToString() < a[j].ToString() }

// NodeParser parses a node.
type NodeParser interface {
	Parse(s string) Node
}

// Edge is an edge in a graph.
type Edge struct {
	From, To Node
}

// Graph that is represented as an adjacency list of nodes.
type Graph map[Node][]Node

// GraphFromMap creates a new graph from a map.
func GraphFromMap(m map[string][]string, p NodeParser) *Graph {
	g := Graph{}
	for k, v := range m {
		for _, n := range v {
			key := p.Parse(k)
			g[key] = append(g[key], p.Parse(n))
		}
	}
	return &g
}

// Read is a single read of length K
type Read string

// ToEdge creates an edge from a Read.
func (r Read) ToEdge() Edge {
	return Edge{From: r[:len(r)-1], To: r[1:]}
}

// ToString returns the read as a string
func (r Read) ToString() string {
	return string(r)
}

// ReadParser implements a NodeParser that parses a Read.
type ReadParser struct{}

// Parse parses a Read.
func (p ReadParser) Parse(s string) Node {
	return Read(s)
}

// ReadPair represents a gapped read of length 2k+d where Read1 and Read2 each
// of length k and gapped by an unknown sequence of length d.
type ReadPair struct {
	Read1, Read2 Read
	Gap          int
}

// ToEdge creates an edge from a ReadPair.
func (r ReadPair) ToEdge() Edge {
	e1 := r.Read1.ToEdge()
	e2 := r.Read2.ToEdge()
	from := ReadPair{
		Read1: e1.From.(Read),
		Read2: e2.From.(Read),
		Gap:   r.Gap + 1,
	}
	to := ReadPair{
		Read1: e1.To.(Read),
		Read2: e2.To.(Read),
		Gap:   r.Gap + 1,
	}
	return Edge{From: from, To: to}
}

// ToString returns the ReadPair as the string Read1|Read2
func (r ReadPair) ToString() string {
	return string(r.Read1) + "|" + string(r.Read2)
}

// ReadPairParser implements a NodeParser that parses a ReadPair.
type ReadPairParser struct {
	Gap int
}

// Parse parses a ReadPair.
func (p ReadPairParser) Parse(s string) Node {
	i := strings.Index(s, "|")
	r1 := Read(s[:i])
	r2 := Read(s[i+1:])
	return ReadPair{Read1: r1, Read2: r2, Gap: p.Gap}
}
