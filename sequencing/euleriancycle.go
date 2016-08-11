package sequencing

import "fmt"

// EulerianCycle finds an Eulerian cycle in a graph assuming that such a cycle
// exists.
func EulerianCycle(graph *Graph) []Node {
	path := make([]Node, 1)
	for k := range *graph {
		path[0] = k
		break
	}
	unusedEdges := copyGraph(graph)
	for {
		path = expand(path, unusedEdges)
		if path[0] != path[len(path)-1] {
			panic("expand returned but no cycle formed")
		}
		if len(*unusedEdges) == 0 {
			return path
		}
		path = path[1:]
		i := nextStartingIndex(path, unusedEdges)
		cyclicShiftLeft(path, i)
		path = append(path, path[0])
	}
}

// expand returns an expanded by by adding nodes to the end of the given path.
// New nodes are added by removing edges from unusedEdges.
func expand(path []Node, unusedEdges *Graph) []Node {
	for {
		lastNodeInPath := path[len(path)-1]
		if _, ok := (*unusedEdges)[lastNodeInPath]; !ok {
			return path
		}
		available := (*unusedEdges)[lastNodeInPath]
		path = append(path, available[0])
		if len(available) == 1 {
			delete(*unusedEdges, lastNodeInPath)
		} else {
			(*unusedEdges)[lastNodeInPath] = available[1:]
		}
	}
}

func nextStartingIndex(path []Node, unusedEdges *Graph) int {
	for i := 0; i < len(path); i++ {
		if _, ok := (*unusedEdges)[path[i]]; ok {
			return i
		}
	}
	panic("No node in path with unused outgoing edge.")
}

func cyclicShiftLeft(path []Node, shiftBy int) {
	if shiftBy == 0 {
		return
	}
	buffer := make([]Node, shiftBy)
	for i := 0; i < len(path); i++ {
		if i < shiftBy {
			buffer[i] = path[i]
		}
		newValIndex := i + shiftBy
		if newValIndex < len(path) {
			path[i] = path[newValIndex]
		} else {
			path[i] = buffer[newValIndex%len(path)]
		}
	}
}

func copyGraph(graph *Graph) *Graph {
	newGraph := Graph{}
	for k, v := range *graph {
		newVal := make([]Node, len(v), len(v))
		for i := 0; i < len(v); i++ {
			newVal[i] = v[i]
		}
		newGraph[k] = newVal
	}
	return &newGraph
}

func print(unusedEdges map[string][]string, path []string) {
	fmt.Println(" unused edges ")
	fmt.Println("--------------")
	for k, v := range unusedEdges {
		fmt.Printf("%s -> %v\n", k, v)
	}
	fmt.Println("     path     ")
	fmt.Println("--------------")
	fmt.Println(path)
	fmt.Println()
}
