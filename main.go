package main

import (
	"os"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

func main() {
	// this is a directed acyclic graph of int hashes
	g := graph.New(graph.IntHash, graph.Directed(), graph.Acyclic())

	// adding 5 vertices
	_ = g.AddVertex(1, graph.VertexAttribute("label", "commit 1"))
	_ = g.AddVertex(2, graph.VertexAttribute("label", "commit 2"))
	_ = g.AddVertex(3, graph.VertexAttribute("label", "commit 3"))
	_ = g.AddVertex(4, graph.VertexAttribute("label", "commit 4"))
	_ = g.AddVertex(5, graph.VertexAttribute("label", "commit 1"))
	_ = g.AddVertex(6, graph.VertexAttribute("label", "commit 5"))
	_ = g.AddVertex(7, graph.VertexAttribute("label", "commit 2"))
	_ = g.AddVertex(8, graph.VertexAttribute("label", "merge into main, commit 6"))
	_ = g.AddVertex(9, graph.VertexAttribute("label", "commit 7"))

	// adding edges between vertices
	_ = g.AddEdge(1, 2, graph.EdgeAttribute("label", "main"))
	_ = g.AddEdge(2, 3, graph.EdgeAttribute("label", "main"))
	_ = g.AddEdge(3, 4, graph.EdgeAttribute("label", "main"))
	_ = g.AddEdge(3, 5, graph.EdgeAttribute("label", "feature"))
	_ = g.AddEdge(5, 7, graph.EdgeAttribute("label", "feature"))
	_ = g.AddEdge(4, 6, graph.EdgeAttribute("label", "main"))
	_ = g.AddEdge(6, 8, graph.EdgeAttribute("label", "main"))
	_ = g.AddEdge(7, 8, graph.EdgeAttribute("label", "merge"))
	_ = g.AddEdge(8, 9, graph.EdgeAttribute("label", "main"))

	// create empty file
	file, _ := os.Create("./DAG-graph.gv")

	// draw graph to file
	_ = draw.DOT(g, file)
}
