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
	_ = g.AddVertex(1)
	_ = g.AddVertex(2)
	_ = g.AddVertex(3)
	_ = g.AddVertex(4)
	_ = g.AddVertex(5)

	// adding edges between vertices
	_ = g.AddEdge(1, 2)
	_ = g.AddEdge(2, 3)
	_ = g.AddEdge(3, 4)
	_ = g.AddEdge(3, 5)
	_ = g.AddEdge(5, 6)
	_ = g.AddEdge(4, 6)

	// create empty file
	file, _ := os.Create("./DAG-graph.gv")

	// draw graph to file
	_ = draw.DOT(g, file)
}
