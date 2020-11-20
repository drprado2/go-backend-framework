package graphstructure

import (
	"fmt"
	"github.com/google/uuid"
)

type Vertice struct {
	GenericData interface{}
	ID          uuid.UUID
}

func NewVertice(data interface{}) *Vertice {
	return &Vertice{
		GenericData: data,
		ID:          uuid.New(),
	}
}

type Edge struct {
	Head        *Vertice
	Tail        *Vertice
	GenericData interface{}
	ID          uuid.UUID
}

func NewEdge(head *Vertice, tail *Vertice, data interface{}) *Edge {
	return &Edge{
		Head:        head,
		Tail:        tail,
		GenericData: data,
		ID:          uuid.New(),
	}
}

type Graph struct {
	vertices map[string]Vertice
	edges    map[string]Edge
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[string]Vertice),
		edges:    make(map[string]Edge),
	}
}

func (g *Graph) AddVertice(vertice *Vertice) error {
	if vertice == nil {
		return fmt.Errorf("Vertice must be not nil")
	}
	g.vertices[vertice.ID.String()] = *vertice
	return nil
}

func (g *Graph) ContainsVertice(verticeId uuid.UUID) bool {
	_, contains := g.vertices[verticeId.String()]
	return contains
}

func (g *Graph) AddEdge(edge *Edge) error {
	if edge == nil {
		return fmt.Errorf("Edge must be not nil")
	}
	if edge.Head != nil && !g.ContainsVertice(edge.Head.ID) {
		if err := g.AddVertice(edge.Head); err != nil {
			return fmt.Errorf("Problem adding head vertice\nError: %s", err)
		}
	}
	if edge.Tail != nil && !g.ContainsVertice(edge.Tail.ID) {
		if err := g.AddVertice(edge.Tail); err != nil {
			return fmt.Errorf("Problem adding tail vertice\nError: %s", err)
		}
	}

	g.edges[edge.ID.String()] = *edge
	return nil
}

func (g *Graph) GetEdges() []Edge {
	values := make([]Edge, 0, len(g.edges))
	for _, v := range g.edges {
		values = append(values, v)
	}
	return values
}

func (g *Graph) GetVertices() []Vertice {
	vertices := make([]Vertice, 0, len(g.vertices))
	for _, v := range g.vertices {
		vertices = append(vertices, v)
	}
	return vertices
}
