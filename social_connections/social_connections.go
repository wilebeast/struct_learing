package social_connections

// GetVisibleProfilesCount returns the size of the connected component
// for each queried user in an undirected graph.
func GetVisibleProfilesCount(nodes int32, u []int32, v []int32, queries []int32) []int32 {
	return getVisibleProfilesCount(nodes, u, v, queries)
}

func getVisibleProfilesCount(nodes int32, u []int32, v []int32, queries []int32) []int32 {
	if nodes <= 0 {
		return nil
	}

	parent := make([]int32, nodes+1)
	size := make([]int32, nodes+1)
	for i := int32(1); i <= nodes; i++ {
		parent[i] = i
		size[i] = 1
	}

	for i := 0; i < len(u) && i < len(v); i++ {
		union(parent, size, u[i], v[i])
	}

	result := make([]int32, len(queries))
	for i, query := range queries {
		root := find(parent, query)
		result[i] = size[root]
	}
	return result
}

func find(parent []int32, x int32) int32 {
	if parent[x] != x {
		parent[x] = find(parent, parent[x])
	}
	return parent[x]
}

func union(parent, size []int32, a, b int32) {
	rootA := find(parent, a)
	rootB := find(parent, b)
	if rootA == rootB {
		return
	}

	if size[rootA] < size[rootB] {
		rootA, rootB = rootB, rootA
	}

	parent[rootB] = rootA
	size[rootA] += size[rootB]
}
