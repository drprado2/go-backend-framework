package graphstructure

import (
	"fmt"
	"github.com/drprado2/go-backend-framework/pkg/queuestructure"
	"github.com/drprado2/go-backend-framework/pkg/stackstructure"
	"github.com/google/uuid"
)

type PathPoint struct {
	Vertex       Vertex
	WeightUpHere int
}

type Vertex struct {
	GenericData           interface{}
	ID                    uuid.UUID
	edgesAdjacentVertices []*Edge
}

func NewVertice(data interface{}) *Vertex {
	return &Vertex{
		GenericData:           data,
		ID:                    uuid.New(),
		edgesAdjacentVertices: make([]*Edge, 0, 10),
	}
}

type Edge struct {
	Head        *Vertex
	Tail        *Vertex
	Weight      int
	GenericData interface{}
}

type Graph struct {
	vertexes            map[string]*Vertex
	acceptCycles        bool
	isDirected          bool
	checkCycleOnAddEdge bool
}

func NewDirectedAcyclicGraph(checkCycleOnAddEdge bool) *Graph {
	return &Graph{
		vertexes:            make(map[string]*Vertex),
		acceptCycles:        false,
		isDirected:          true,
		checkCycleOnAddEdge: checkCycleOnAddEdge,
	}
}

func NewDirectedCyclicGraph() *Graph {
	return &Graph{
		vertexes:            make(map[string]*Vertex),
		acceptCycles:        true,
		isDirected:          true,
		checkCycleOnAddEdge: false,
	}
}

func NewUndirectedGraph() *Graph {
	return &Graph{
		vertexes:            make(map[string]*Vertex),
		acceptCycles:        true,
		isDirected:          false,
		checkCycleOnAddEdge: false,
	}
}

func (g *Graph) addVertice(vertice Vertex) error {
	if g.ContainsVertice(vertice.ID) {
		return fmt.Errorf("Vertex %s already exists, use UpdateVerticeData to change the vertice data", vertice.ID.String())
	}
	g.vertexes[vertice.ID.String()] = &vertice
	return nil
}

func (g *Graph) AddVertice(vertices ...Vertex) error {
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

func (g *Graph) UpdateVerticeData(verticeId uuid.UUID, data interface{}) error {
	vertice, ok := g.vertexes[verticeId.String()]
	if !ok {
		return fmt.Errorf("Vertex not found")
	}
	vertice.GenericData = data
	return nil
}

func (g *Graph) ContainsVertice(verticeId uuid.UUID) bool {
	_, contains := g.vertexes[verticeId.String()]
	return contains
}

func (g *Graph) AddEdge(fromVerticeId uuid.UUID, toVerticeId uuid.UUID, weight int, extraData interface{}) error {
	if !g.ContainsVertice(fromVerticeId) || !g.ContainsVertice(toVerticeId) {
		return fmt.Errorf("The vertice From and vertice To must be added to the directedAcycleGraph before creating the edge")
	}

	verticeFrom := g.vertexes[fromVerticeId.String()]
	verticeTo := g.vertexes[toVerticeId.String()]

	for _, currentEdge := range verticeFrom.edgesAdjacentVertices {
		if currentEdge.Head.ID == verticeTo.ID {
			return fmt.Errorf("Already exists one edge between %s and %s, use UpdateEdgeData to change the edge data", fromVerticeId.String(), toVerticeId.String())
		}
	}

	if !g.isDirected {
		backEdge := &Edge{
			Head:        verticeFrom,
			Tail:        verticeTo,
			Weight:      weight,
			GenericData: extraData,
		}
		verticeTo.edgesAdjacentVertices = append(verticeTo.edgesAdjacentVertices, backEdge)
	}

	edge := &Edge{
		Head:        verticeTo,
		Tail:        verticeFrom,
		Weight:      weight,
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

	verticeFrom := g.vertexes[fromVerticeId.String()]
	verticeTo := g.vertexes[toVerticeId.String()]

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

	for _, vertice := range g.vertexes {
		for i, edge := range vertice.edgesAdjacentVertices {
			if edge.Head.ID == verticeId {
				vertice.edgesAdjacentVertices = append(vertice.edgesAdjacentVertices[:i], vertice.edgesAdjacentVertices[i+1:]...)
			}
		}
	}

	delete(g.vertexes, verticeId.String())
	return nil
}

func (g *Graph) RemoveEdge(fromVerticeId uuid.UUID, toVerticeId uuid.UUID) error {
	from, containsFrom := g.vertexes[fromVerticeId.String()]
	to, containsTo := g.vertexes[toVerticeId.String()]

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

	for _, v := range g.vertexes {
		for _, edge := range v.edgesAdjacentVertices {
			values = append(values, *edge)
		}
	}
	return values
}

func (g *Graph) GetVertices() []Vertex {
	vertices := make([]Vertex, 0, len(g.vertexes))
	for _, v := range g.vertexes {
		vertices = append(vertices, *v)
	}
	return vertices
}

func (g *Graph) ExistsCycle() bool {
	if !g.isDirected {
		return true
	}

	alreadyVisited := make([]uuid.UUID, 0, len(g.vertexes))
	visitedControl := stackstructure.NewStack(3)

	for _, vertex := range g.vertexes {
		if g.existsCycle(vertex, alreadyVisited, visitedControl) {
			return true
		}
	}
	return false
}

func (g *Graph) existsCycle(vertex *Vertex, alreadyVisiteds []uuid.UUID, visitedInCurrentScan *stackstructure.Stack) bool {
	if visitedInCurrentScan.PositionOfElement(vertex.ID, func(elemA interface{}, elemB interface{}) bool {
		idA := elemA.(uuid.UUID)
		idB := elemB.(uuid.UUID)
		return idA == idB
	}) > -1 {
		return true
	}
	for _, id := range alreadyVisiteds {
		if vertex.ID == id {
			return false
		}
	}

	visitedInCurrentScan.StackUp(vertex.ID)
	alreadyVisiteds = append(alreadyVisiteds, vertex.ID)

	for _, edge := range vertex.edgesAdjacentVertices {
		if g.existsCycle(edge.Head, alreadyVisiteds, visitedInCurrentScan) {
			return true
		}
	}

	visitedInCurrentScan.Unstack()
	return false
}

func (g *Graph) GetCycles() [][]uuid.UUID {
	alreadyVisited := make([]uuid.UUID, 0, len(g.vertexes))
	visitedControl := stackstructure.NewStack(3)
	result := make([][]uuid.UUID, 0)

	for _, vertex := range g.vertexes {
		g.getCycles(vertex, alreadyVisited, visitedControl, result)
	}

	return result
}

func (g *Graph) getCycles(vertex *Vertex, alreadyVisited []uuid.UUID, visitedInCurrentScan *stackstructure.Stack, result [][]uuid.UUID) {
	if visitedInCurrentScan.PositionOfElement(vertex.ID, func(elemA interface{}, elemB interface{}) bool {
		idA := elemA.(uuid.UUID)
		idB := elemB.(uuid.UUID)
		return idA == idB
	}) > -1 {
		cycle := visitedInCurrentScan.CopyToSlice()
		cycleIds := make([]uuid.UUID, 0, len(cycle)+1)
		for i, id := range cycle {
			cycleIds[i] = id.(uuid.UUID)
		}
		cycleIds[len(cycle)+1] = vertex.ID
		result = append(result, cycleIds)
		return
	}
	for _, id := range alreadyVisited {
		if id == vertex.ID {
			return
		}
	}
	visitedInCurrentScan.StackUp(vertex)
	alreadyVisited = append(alreadyVisited, vertex.ID)

	for _, edge := range vertex.edgesAdjacentVertices {
		g.getCycles(edge.Head, alreadyVisited, visitedInCurrentScan, result)
	}

	visitedInCurrentScan.Unstack()
}

func (g *Graph) BreadthFirstSearch(fromVertexId uuid.UUID, toVertexId uuid.UUID) (*Vertex, error) {
	vertexFrom, findFrom := g.vertexes[fromVertexId.String()]
	_, findTo := g.vertexes[toVertexId.String()]

	if !findFrom || !findTo {
		return nil, fmt.Errorf("VertexFrom or vertexTo doen`s exist in the graph")
	}

	alreadyVisitedIds := make([]uuid.UUID, 0, len(g.vertexes))
	vertexResult := g.breadthFirstSearch(vertexFrom, toVertexId, alreadyVisitedIds)
	if vertexResult == nil {
		return nil, nil
	}

	var result Vertex
	result = *vertexResult
	return &result, nil
}

func (g *Graph) breadthFirstSearch(vertexFrom *Vertex, toVertexId uuid.UUID, alreadyVisitedIds []uuid.UUID) *Vertex {
	for _, el := range alreadyVisitedIds {
		if vertexFrom.ID == el {
			return nil
		}
	}
	alreadyVisitedIds = append(alreadyVisitedIds, vertexFrom.ID)

	elementsToVisit := queuestructure.NewQueue(len(vertexFrom.edgesAdjacentVertices))

	for _, edge := range vertexFrom.edgesAdjacentVertices {
		if edge.Head.ID == toVertexId {
			return edge.Head
		}
		elementsToVisit.Enqueue(edge.Head)
	}
	for nextEl := elementsToVisit.Next(); nextEl != nil; {
		if result := g.breadthFirstSearch(nextEl.(*Vertex), toVertexId, alreadyVisitedIds); result != nil {
			return result
		}
	}
	return nil
}

func (g *Graph) DepthFirstSearch(fromVertexId uuid.UUID, toVertexId uuid.UUID) (*Vertex, error) {
	vertexFrom, findFrom := g.vertexes[fromVertexId.String()]
	_, findTo := g.vertexes[toVertexId.String()]

	if !findFrom || !findTo {
		return nil, fmt.Errorf("VertexFrom or vertexTo doen`s exist in the graph")
	}

	alreadyVisitedIds := make([]uuid.UUID, 0, len(g.vertexes))
	vertexFound := g.depthFirstSearch(vertexFrom, toVertexId, alreadyVisitedIds)

	if vertexFound == nil {
		return nil, nil
	}

	var result Vertex
	result = *vertexFound
	return &result, nil
}

func (g *Graph) depthFirstSearch(vertexFrom *Vertex, vertexToId uuid.UUID, alreadyVisitedIds []uuid.UUID) *Vertex {
	for _, el := range alreadyVisitedIds {
		if vertexFrom.ID == el {
			return nil
		}
	}
	alreadyVisitedIds = append(alreadyVisitedIds, vertexFrom.ID)

	for _, edge := range vertexFrom.edgesAdjacentVertices {
		if edge.Head.ID == vertexToId {
			return edge.Head
		}
		if result := g.depthFirstSearch(edge.Head, vertexToId, alreadyVisitedIds); result != nil {
			return result
		}
	}
	return nil
}

// TODO implementar
func (g *Graph) GetRelatedVertices(fromVertices []uuid.UUID) ([]Vertex, error) {
	return make([][]uuid.UUID, 0), nil
}

func (g *Graph) FindShortestPath(fromVerticeId uuid.UUID, toVerticeId uuid.UUID) ([]PathPoint, error) {
	verticeFrom, foundFrom := g.vertexes[fromVerticeId.String()]
	_, foundTo := g.vertexes[toVerticeId.String()]

	if !foundFrom || !foundTo {
		return nil, fmt.Errorf("VerticeFrom or verticeTo doesn`t exist in the graph")
	}

	sealedVertexes := make([]uuid.UUID, 0, 3)
	pathPoints := make([]PathPoint, 0, 3)
	pathPoints[0] = PathPoint{
		Vertex:       *verticeFrom,
		WeightUpHere: 0,
	}

}

func (g *Graph) findShortestPath(vertex *Vertex, toVerticeId uuid.UUID, sealedVertexes []uuid.UUID) ([]PathPoint, error) {
	for _, sealed := range sealedVertexes{
		if vertex.ID == sealed{
			continue
		}
	}
	for _, edge := range vertex.edgesAdjacentVertices {

		if edge.Head.

	}

}
