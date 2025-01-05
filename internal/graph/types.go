package graph

// Node and Edge structs represent the graph data
type Node struct {
	Data map[string]string `json:"data"`
}

type Edge struct {
	Data map[string]string `json:"data"`
}

type Graph struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}
