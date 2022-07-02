package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

type Graph struct {
}

const (
	// Start - стартовая вершина
	Start = 3
	// Vert - кол-во вершин
	Vert = 7
)

var (
	used  = make(map[int]bool)
	graph = [Vert][]int{
		{1, 2},
		{0, 2},
		{0, 1, 3, 4, 5},
		{2},
		{2, 6},
		{2, 6},
		{4, 5},
	}
)

func main() {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := graph.Close(); err != nil {
			log.Fatal(err)
		}
		g.Close()
	}()

	gnr, err := graph.CreateNode(strconv.Itoa(Start))
	if err != nil {
		log.Fatal(err)
	}
	dfs(Start, graph, gnr)

	var buf bytes.Buffer
	if err := g.Render(graph, "dot", &buf); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())

}

func dfs(start int, g *cgraph.Graph, gnr *cgraph.Node) {
	used[start] = true
	for _, n := range graph {
		for _, v := range n {
			gn, err := g.CreateNode(strconv.Itoa(v))
			if err != nil {
				log.Fatal(err)
			}

			e, err := g.CreateEdge("", gnr, gn)
			if err != nil {
				log.Fatal(err)
			}
			e.SetLabel(strconv.Itoa(v))
			if _, ok := used[v]; !ok {
				dfs(v, g, gn)
			}
		}
	}
}
