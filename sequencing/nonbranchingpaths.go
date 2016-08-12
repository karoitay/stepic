package sequencing

// MaximalNonBranchingPaths return all paths that are maximal non branching
// paths in g.
func MaximalNonBranchingPaths(g *Graph) [][]Node {
	var res [][]Node
	unused := copyGraph(g)
	for n := range *g {
		if len((*g)[n]) == 0 || isOneToOneNode(n, g) {
			continue
		}
		for _, o := range (*g)[n] {
			next := o
			path := []Node{n, next}
			delete(*unused, n)
			delete(*unused, next)
			for isOneToOneNode(next, g) {
				next = (*g)[next][0]
				path = append(path, next)
				delete(*unused, next)
			}
			res = append(res, path)
		}
	}
	cycle := findOneToOneCycle(unused)
	for cycle != nil {
		res = append(res, cycle)
		cycle = findOneToOneCycle(unused)
	}
	return res
}

func isOneToOneNode(n Node, g *Graph) bool {
	return len((*g)[n]) == 1 && inDegree(n, g) == 1
}

func inDegree(n Node, g *Graph) int {
	d := 0
	for _, v := range *g {
		for _, to := range v {
			if n == to {
				d++
			}
		}
	}
	return d
}

func findOneToOneCycle(unused *Graph) []Node {
	var n Node
	for k := range *unused {
		n = k
		break
	}
	if n == nil {
		return nil
	}
	path := []Node{n}
	n, ok := getSingleChild(unused, n)
	for ok {
		if n == path[0] {
			if len(path) != 1 {
				path = append(path, n)
			}
			return path
		}
		path = append(path, n)
		n, ok = getSingleChild(unused, n)
	}
	return nil
}

func getSingleChild(unused *Graph, n Node) (Node, bool) {
	next, ok := (*unused)[n]
	delete(*unused, n)
	return next[0], ok
}
