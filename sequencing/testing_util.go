package sequencing

import (
	"reflect"
	"sort"
)

type nodes []Node

func (a nodes) Len() int           { return len(a) }
func (a nodes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a nodes) Less(i, j int) bool { return a[i].ToString() < a[j].ToString() }

func (g *Graph) equals(other *Graph) bool {
	if g == nil && other == nil {
		return true
	} else if g == nil || other == nil {
		return false
	}
	o := copyGraph(other)
	t := copyGraph(g)
	for _, v := range *o {
		sort.Sort(nodes(v))
	}
	for _, v := range *t {
		sort.Sort(nodes(v))
	}
	return reflect.DeepEqual(t, o)
}

type edges []Edge

func (es edges) Len() int      { return len(es) }
func (es edges) Swap(i, j int) { es[i], es[j] = es[j], es[i] }
func (es edges) Less(i, j int) bool {
	from1 := es[i].From.ToString()
	from2 := es[j].From.ToString()
	if from1 != from2 {
		return from1 < from2
	}
	return es[i].To.ToString() < es[j].To.ToString()
}

type path []Node

func (p path) copy() path {
	path := make([]Node, len(p))
	for i := 0; i < len(p); i++ {
		path[i] = p[i]
	}
	return path
}

func (p path) asEdges() edges {
	es := make([]Edge, len(p)-1)
	for i := 1; i < len(p); i++ {
		es[i-1] = Edge{From: p[i-1], To: p[i]}
	}
	return es
}

type paths []path

func (ps paths) Len() int      { return len(ps) }
func (ps paths) Swap(i, j int) { ps[i], ps[j] = ps[j], ps[i] }
func (ps paths) Less(i, j int) bool {
	if len(ps[i]) != len(ps[j]) {
		return len(ps[i]) < len(ps[j])
	}
	es1 := ps[i].asEdges()
	es2 := ps[j].asEdges()
	sort.Sort(es1)
	sort.Sort(es2)
	for k := 0; k < len(es1); k++ {
		if es1[k] != es2[k] {
			return edges([]Edge{es1[k], es2[k]}).Less(0, 1)
		}
	}
	return true
}

func (ps paths) copy() paths {
	paths := make([]path, len(ps))
	for i := 0; i < len(ps); i++ {
		paths[i] = ps[i].copy()
	}
	return paths
}

func (ps paths) equals(other paths) bool {
	if len(ps) != len(other) {
		return false
	}
	ps1 := ps.copy()
	ps2 := other.copy()
	sort.Sort(ps1)
	sort.Sort(ps2)
	for i := 0; i < len(ps); i++ {
		es1 := ps1[i].asEdges()
		es2 := ps2[i].asEdges()
		sort.Sort(es1)
		sort.Sort(es2)
		if !reflect.DeepEqual(es1, es2) {
			return false
		}
	}
	return true
}
