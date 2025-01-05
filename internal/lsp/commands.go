package lsp

import (
	"zeta/internal/graph"
)

func (s *Server) showGraph() error {

	// Retrieve files and links from the document manager
	files, links, err := s.docManager.GetGraph()
	if err != nil {
		return err
	}

	// Create a graph from files and links
	g := graph.Graph{
		Nodes: []graph.Node{},
		Edges: []graph.Edge{},
	}

	// Populate nodes from the files
	for _, file := range files {
		g.Nodes = append(g.Nodes, graph.Node{
			Data: map[string]string{
				"id":    file.Path,
				"label": file.Path,
			},
		})
	}

	// Populate edges from the links
	for _, link := range links {
		g.Edges = append(g.Edges, graph.Edge{
			Data: map[string]string{
				"source": link.SourcePath,
				"target": link.TargetPath,
			},
		})
	}

	graph.ShowGraph(g)
	return nil
}
