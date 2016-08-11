package sequencing

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

// Edge is an edge in a graph.
type Edge struct {
	From, To Node
}

// Graph that is represented as an adjacency list of nodes.
type Graph map[Node][]Node

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
