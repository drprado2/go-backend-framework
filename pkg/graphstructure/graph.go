package graphstructure

import (
	"fmt"
	"github.com/google/uuid"
)

type Vertice struct {
	GenericData           interface{}
	ID                    uuid.UUID
	edgesAdjacentVertices []*Edge
}

func NewVertice(data interface{}) *Vertice {
	return &Vertice{
		GenericData:           data,
		ID:                    uuid.New(),
		edgesAdjacentVertices: make([]*Edge, 0, 10),
	}
}

type Edge struct {
	Head        *Vertice
	Tail        *Vertice
	GenericData interface{}
}

type Graph struct {
	vertices            map[string]*Vertice
	acceptCycles        bool
	isDirected          bool
	checkCycleOnAddEdge bool
}

func NewDirectedAcyclicGraph(checkCycleOnAddEdge bool) *Graph {
	return &Graph{
		vertices:            make(map[string]*Vertice),
		acceptCycles:        false,
		isDirected:          true,
		checkCycleOnAddEdge: checkCycleOnAddEdge,
	}
}

func NewDirectedCyclicGraph() *Graph {
	return &Graph{
		vertices:            make(map[string]*Vertice),
		acceptCycles:        true,
		isDirected:          true,
		checkCycleOnAddEdge: false,
	}
}

func NewUndirectedGraph() *Graph {
	return &Graph{
		vertices:            make(map[string]*Vertice),
		acceptCycles:        true,
		isDirected:          false,
		checkCycleOnAddEdge: false,
	}
}

func (g *Graph) addVertice(vertice Vertice) error {
	if g.ContainsVertice(vertice.ID) {
		return fmt.Errorf("Vertice %s already exists, use UpdateVerticeData to change the vertice data", vertice.ID.String())
	}
	g.vertices[vertice.ID.String()] = &vertice
	return nil
}

func (g *Graph) AddVertice(vertices ...Vertice) error {
	errors := ""
	for _, vertice := range vertices {
		err := g.addVertice(vertice)
		if err != nil {
			errors += err.Error()
		}
	}
	if len(errors) > 0 {
		return fmt.Errorf(errors)
	}
	return nil
}


func (g *Graph) ContainsVertice(verticeId uuid.UUID) bool {
	_, contains := g.vertices[verticeId.String()]
	return contains
}

// TODO implementar
func (g *Graph) ExistsCycle() bool {
	return false
}

func (g *Graph) AddEdge(fromVerticeId uuid.UUID, toVerticeId uuid.UUID, extraData interface{}) error {
	if !g.ContainsVertice(fromVerticeId) || !g.ContainsVertice(toVerticeId) {
		return fmt.Errorf("The vertice From and vertice To must be added to the directedAcycleGraph before creating the edge")
	}

	verticeFrom := g.vertices[fromVerticeId.String()]
	verticeTo := g.vertices[toVerticeId.String()]

	for _, currentEdge := range verticeFrom.edgesAdjacentVertices {
		if currentEdge.Head.ID == verticeTo.ID {
			return fmt.Errorf("Already exists one edge between %s and %s, use UpdateEdgeData to change the edge data", fromVerticeId.String(), toVerticeId.String())
		}
	}

	if !g.isDirected {
		backEdge := &Edge{
			Head:        verticeFrom,
			Tail:        verticeTo,
			GenericData: extraData,
		}
		verticeTo.edgesAdjacentVertices = append(verticeTo.edgesAdjacentVertices, backEdge)
	}

	edge := &Edge{
		Head:        verticeTo,
		Tail:        verticeFrom,
		GenericData: extraData,
	}
	verticeFrom.edgesAdjacentVertices = append(verticeFrom.edgesAdjacentVertices, edge)

	if g.checkCycleOnAddEdge && g.ExistsCycle() {
		verticeFrom.edgesAdjacentVertices = verticeFrom.edgesAdjacentVertices[:len(verticeFrom.edgesAdjacentVertices)-2]
		return fmt.Errorf("This edge will cause a cycle in the directedAcycleGraph")
	}

	return nil
}

func (g *Graph) UpdateEdgeData(fromVerticeId uuid.UUID, toVerticeId uuid.UUID, extraData interface{}) error {
	if !g.ContainsVertice(fromVerticeId) || !g.ContainsVertice(toVerticeId) {
		return fmt.Errorf("The vertice From and vertice To must be added to the directedAcycleGraph before creating the edge")
	}

	verticeFrom := g.vertices[fromVerticeId.String()]
	verticeTo := g.vertices[toVerticeId.String()]

	for _, currentEdge := range verticeFrom.edgesAdjacentVertices {
		if currentEdge.Head.ID == verticeTo.ID {
			currentEdge.GenericData = extraData
			break
		}
	}

	if !g.isDirected {
		for _, currentEdge := range verticeTo.edgesAdjacentVertices {
			if currentEdge.Head.ID == verticeFrom.ID {
				currentEdge.GenericData = extraData
				break
			}
		}
	}

	return nil
}

func (g *Graph) RemoveVertice(verticeId uuid.UUID) error {
	if !g.ContainsVertice(verticeId) {
		fmt.Errorf("The vertice %s not exists in the directedAcycleGraph", verticeId.String())
	}

	for _, vertice := range g.vertices {
		for i, edge := range vertice.edgesAdjacentVertices {
			if edge.Head.ID == verticeId {
				vertice.edgesAdjacentVertices = append(vertice.edgesAdjacentVertices[:i], vertice.edgesAdjacentVertices[i+1:]...)
			}
		}
	}

	delete(g.vertices, verticeId.String())
	return nil
}

func (g *Graph) RemoveEdge(fromVerticeId uuid.UUID, toVerticeId uuid.UUID) error {
	from, containsFrom := g.vertices[fromVerticeId.String()]
	to, containsTo := g.vertices[toVerticeId.String()]

	if !containsFrom || !containsTo {
		return fmt.Errorf("The verticeFrom or verticeTo doesn`t exist")
	}

	for i, edge := range from.edgesAdjacentVertices {
		if edge.Head.ID == toVerticeId {
			from.edgesAdjacentVertices = append(from.edgesAdjacentVertices[:i], from.edgesAdjacentVertices[i+1:]...)
			break
		}
	}
	if !g.isDirected {
		for i, edge := range to.edgesAdjacentVertices {
			if edge.Head.ID == fromVerticeId {
				to.edgesAdjacentVertices = append(to.edgesAdjacentVertices[:i], to.edgesAdjacentVertices[i+1:]...)
				break
			}
		}
	}
	return nil
}

func (g *Graph) GetEdges() []Edge {
	values := make([]Edge, 0, 10)

	for _, v := range g.vertices {
		for _, edge := range v.edgesAdjacentVertices {
			values = append(values, *edge)
		}
	}
	return values
}

func (g *Graph) GetVertices() []Vertice {
	vertices := make([]Vertice, 0, len(g.vertices))
	for _, v := range g.vertices {
		vertices = append(vertices, *v)
	}
	return vertices
}

// TODO implementar
func (g *Graph) HasAdjacency(verticeA *Vertice, verticeB *Vertice) {

}

// TODO implementar
func (g *Graph) GetNeighbors(vertice *Vertice) {

}
