package sequencing

import "strconv"

// EulerianPath finds an Eulerian path in a graph
func EulerianPath(graph *Graph) []Node {
	from, to := findEdgeToAdd(graph)
	if from == nil || to == nil {
		return EulerianCycle(graph)
	}
	(*graph)[*from] = append((*graph)[*from], *to)
	path := EulerianCycle(graph)[1:]
	// now we need to remove from->to from the path.
	i := 1
	for i < len(path) {
		if path[i-1] == *from && path[i] == *to {
			cyclicShiftLeft(path, i)
			break
		}
		i++
	}
	return path
}

func unbalancedNodes(graph *Graph) map[Node]int {
	m := map[Node]int{}
	for k, v := range *graph {
		addDegree(m, k, len(v))
		for _, n := range v {
			addDegree(m, n, -1)
		}
	}
	if len(m) != 2 && len(m) != 0 {
		panic("There are " + strconv.Itoa(len(m)) + " unbalanced nodes")
	}
	return m
}

func addDegree(m map[Node]int, node Node, degree int) {
	m[node] = m[node] + degree
	if m[node] == 0 {
		delete(m, node)
	}
}

func findEdgeToAdd(graph *Graph) (*Node, *Node) {
	un := unbalancedNodes(graph)
	if len(un) != 2 {
		return nil, nil
	}
	var from, to Node
	for k, v := range un {
		if v == 1 {
			to = k
		} else if v == -1 {
			from = k
		} else {
			panic(k.ToString() + " degree is not 1 but " + strconv.Itoa(v))
		}
	}
	if from == nil || to == nil {
		panic("could not build a cyclic graph")
	}
	return &from, &to
}
