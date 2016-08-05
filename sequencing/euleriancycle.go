package sequencing

import "fmt"

// EulerianCycle finds an Eulerian cycle in a graph assuming that such a cycle
// exists.
func EulerianCycle(graph map[string][]string) []string {
	path := make([]string, 1)
	for k := range graph {
		path[0] = k
		break
	}
	unusedEdges := copyGraph(graph)
	for {
		path = expand(path, unusedEdges)
		if path[0] != path[len(path)-1] {
			panic("expand returned but no cycle formed")
		}
		if len(unusedEdges) == 0 {
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
func expand(path []string, unusedEdges map[string][]string) []string {
	for {
		lastNodeInPath := path[len(path)-1]
		if _, ok := unusedEdges[lastNodeInPath]; !ok {
			return path
		}
		available := unusedEdges[lastNodeInPath]
		path = append(path, available[0])
		if len(available) == 1 {
			delete(unusedEdges, lastNodeInPath)
		} else {
			unusedEdges[lastNodeInPath] = available[1:]
		}
	}
}

func nextStartingIndex(path []string, unusedEdges map[string][]string) int {
	for i := 0; i < len(path); i++ {
		if _, ok := unusedEdges[path[i]]; ok {
			return i
		}
	}
	panic("No node in path with unused outgoing edge.")
}

func cyclicShiftLeft(path []string, shiftBy int) {
	if shiftBy == 0 {
		return
	}
	buffer := make([]string, shiftBy)
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

func copyGraph(graph map[string][]string) map[string][]string {
	newGraph := map[string][]string{}
	for k, v := range graph {
		newVal := make([]string, len(v), len(v))
		for i := 0; i < len(v); i++ {
			newVal[i] = v[i]
		}
		newGraph[k] = newVal
	}
	return newGraph
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
