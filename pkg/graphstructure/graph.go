package graphstructure

import (
	"fmt"
	"github.com/drprado2/go-backend-framework/pkg/queuestructure"
	"github.com/drprado2/go-backend-framework/pkg/stackstructure"
	"github.com/drprado2/go-backend-framework/pkg/structdoublylinkedlist"
	"github.com/google/uuid"
	"reflect"
)

type PathPoint struct {
	Vertex       Vertex
	WeightUpHere int
}

type Vertex struct {
	GenericData           interface{}
	ID                    string
	edgesAdjacentVertices []*Edge
}

func NewVertice(data interface{}) *Vertex {
	return &Vertex{
		GenericData:           data,
		ID:                    uuid.New().String(),
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
		return fmt.Errorf("Vertex %s already exists, use UpdateVerticeData to change the vertice data", vertice.ID)
	}
	g.vertexes[vertice.ID] = &vertice
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

func (g *Graph) UpdateVerticeData(verticeId string, data interface{}) error {
	vertice, ok := g.vertexes[verticeId]
	if !ok {
		return fmt.Errorf("Vertex not found")
	}
	vertice.GenericData = data
	return nil
}

func (g *Graph) ContainsVertice(verticeId string) bool {
	_, contains := g.vertexes[verticeId]
	return contains
}

func (g *Graph) AddEdge(fromVerticeId string, toVerticeId string, weight int, extraData interface{}) error {
	if !g.ContainsVertice(fromVerticeId) || !g.ContainsVertice(toVerticeId) {
		return fmt.Errorf("The vertice From and vertice To must be added to the directedAcycleGraph before creating the edge")
	}

	verticeFrom := g.vertexes[fromVerticeId]
	verticeTo := g.vertexes[toVerticeId]

	for _, currentEdge := range verticeFrom.edgesAdjacentVertices {
		if currentEdge.Head.ID == verticeTo.ID {
			return fmt.Errorf("Already exists one edge between %s and %s, use UpdateEdgeData to change the edge data", fromVerticeId, toVerticeId)
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

func (g *Graph) UpdateEdgeData(fromVerticeId string, toVerticeId string, extraData interface{}) error {
	if !g.ContainsVertice(fromVerticeId) || !g.ContainsVertice(toVerticeId) {
		return fmt.Errorf("The vertice From and vertice To must be added to the directedAcycleGraph before creating the edge")
	}

	verticeFrom := g.vertexes[fromVerticeId]
	verticeTo := g.vertexes[toVerticeId]

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

func (g *Graph) RemoveVertice(verticeId string) error {
	if !g.ContainsVertice(verticeId) {
		fmt.Errorf("The vertice %s not exists in the directedAcycleGraph", verticeId)
	}

	for _, vertice := range g.vertexes {
		for i, edge := range vertice.edgesAdjacentVertices {
			if edge.Head.ID == verticeId {
				vertice.edgesAdjacentVertices = append(vertice.edgesAdjacentVertices[:i], vertice.edgesAdjacentVertices[i+1:]...)
			}
		}
	}

	delete(g.vertexes, verticeId)
	return nil
}

func (g *Graph) RemoveEdge(fromVerticeId string, toVerticeId string) error {
	from, containsFrom := g.vertexes[fromVerticeId]
	to, containsTo := g.vertexes[toVerticeId]

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

	alreadyVisited := make([]string, 0, len(g.vertexes))
	visitedControl := stackstructure.NewStack(3)

	for _, vertex := range g.vertexes {
		if g.existsCycle(vertex, alreadyVisited, visitedControl) {
			return true
		}
	}
	return false
}

func (g *Graph) existsCycle(vertex *Vertex, alreadyVisiteds []string, visitedInCurrentScan *stackstructure.Stack) bool {
	if visitedInCurrentScan.PositionOfElement(vertex.ID, func(elemA interface{}, elemB interface{}) bool {
		idA := elemA.(string)
		idB := elemB.(string)
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

func (g *Graph) GetCycles() [][]string {
	alreadyVisited := make([]string, 0, len(g.vertexes))
	visitedControl := stackstructure.NewStack(3)
	result := make([][]string, 0)

	for _, vertex := range g.vertexes {
		g.getCycles(vertex, alreadyVisited, visitedControl, result)
	}

	return result
}

func (g *Graph) getCycles(vertex *Vertex, alreadyVisited []string, visitedInCurrentScan *stackstructure.Stack, result [][]string) {
	if visitedInCurrentScan.PositionOfElement(vertex.ID, func(elemA interface{}, elemB interface{}) bool {
		idA := elemA.(string)
		idB := elemB.(string)
		return idA == idB
	}) > -1 {
		cycle := visitedInCurrentScan.CopyToSlice()
		cycleIds := make([]string, 0, len(cycle)+1)
		for i, id := range cycle {
			cycleIds[i] = id.(string)
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

func (g *Graph) BreadthFirstSearch(fromVertexId string, toVertexId string) (*Vertex, error) {
	vertexFrom, findFrom := g.vertexes[fromVertexId]
	_, findTo := g.vertexes[toVertexId]

	if !findFrom || !findTo {
		return nil, fmt.Errorf("VertexFrom or vertexTo doen`s exist in the graph")
	}

	alreadyVisitedIds := make([]string, 0, len(g.vertexes))
	vertexResult := g.breadthFirstSearch(vertexFrom, toVertexId, alreadyVisitedIds)
	if vertexResult == nil {
		return nil, nil
	}

	var result Vertex
	result = *vertexResult
	return &result, nil
}

func (g *Graph) breadthFirstSearch(vertexFrom *Vertex, toVertexId string, alreadyVisitedIds []string) *Vertex {
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

func (g *Graph) DepthFirstSearch(fromVertexId string, toVertexId string) (*Vertex, error) {
	vertexFrom, findFrom := g.vertexes[fromVertexId]
	_, findTo := g.vertexes[toVertexId]

	if !findFrom || !findTo {
		return nil, fmt.Errorf("VertexFrom or vertexTo doen`s exist in the graph")
	}

	alreadyVisitedIds := make([]string, 0, len(g.vertexes))
	vertexFound := g.depthFirstSearch(vertexFrom, toVertexId, alreadyVisitedIds)

	if vertexFound == nil {
		return nil, nil
	}

	var result Vertex
	result = *vertexFound
	return &result, nil
}

func (g *Graph) depthFirstSearch(vertexFrom *Vertex, vertexToId string, alreadyVisitedIds []string) *Vertex {
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

func getVertexesList() *structdoublylinkedlist.List {
	equalityFunc := func(elemA interface{}, elemB interface{}) bool {
		vertexA, okA := elemA.(*Vertex)
		vertexB, okB := elemB.(*Vertex)

		if !okA || !okB {
			return false
		}

		return vertexA.ID == vertexB.ID
	}
	return structdoublylinkedlist.NewList(equalityFunc, reflect.TypeOf(&Vertex{}))
}

func (g *Graph) GetDependents(fromVertices []string) ([]Vertex, error) {
	vertexes := make([]*Vertex, 0, len(fromVertices))
	for i, id := range fromVertices {
		v, ok := g.vertexes[id]
		if !ok {
			return nil, fmt.Errorf("The vertex %v doens`t exist in the graph", id)
		}
		vertexes[i] = v
	}

	edges := g.GetEdges()
	resultList := getVertexesList()
	for _, vertex := range vertexes {
		g.getDependents(vertex, edges, resultList)
	}
	result := make([]Vertex, 0, resultList.Lenght())
	for elem := resultList.Unshift(); elem != nil; elem = resultList.Unshift() {
		vertex := elem.(*Vertex)
		result = append(result, *vertex)
	}
	return result, nil
}

func (g *Graph) getDependents(vertex *Vertex, edges []Edge, dependents *structdoublylinkedlist.List) {
	if dependents.Exists(vertex) {
		dependents.Remove(vertex)
		dependents.Add(vertex)
		return
	}
	dependents.Add(vertex)

	for _, edge := range edges {
		if edge.Head.ID == vertex.ID {
			g.getDependents(edge.Tail, edges, dependents)
		}
	}
}

func (g *Graph) FindShortestPath(fromVerticeId string, toVerticeId string) ([]PathPoint, error) {
	verticeFrom, foundFrom := g.vertexes[fromVerticeId]
	_, foundTo := g.vertexes[toVerticeId]

	if !foundFrom || !foundTo {
		return nil, fmt.Errorf("VerticeFrom or verticeTo doesn`t exist in the graph")
	}

	sealedVertexes := make(map[string]PathPoint)
	pathPoints := make(map[string]PathPoint)
	pathPoints[verticeFrom.ID] = PathPoint{
		Vertex:       *verticeFrom,
		WeightUpHere: 0,
	}
	if finalPath := g.findShortestPath(pathPoints[verticeFrom.ID], toVerticeId, pathPoints, sealedVertexes); finalPath == nil {
		return nil, fmt.Errorf("There is no way from %s to %s", fromVerticeId, toVerticeId)
	}
	result := make([]PathPoint, 0, 10)
	for path := sealedVertexes[toVerticeId]; path.Vertex.ID != fromVerticeId; path = sealedVertexes[path.Vertex.ID] {
		result = append(result, path)
	}
	result = append(result, sealedVertexes[fromVerticeId])
	for x, z := 0, len(result)-1; x < z; x, z = x+1, z-1 {
		result[x], result[z] = result[z], result[x]
	}
	return result, nil
}

func getPathPointSortedList() *structdoublylinkedlist.List {
	equalityFunc := func(elemA interface{}, elemB interface{}) bool {
		pathA, okA := elemA.(PathPoint)
		pathB, okB := elemB.(PathPoint)

		if !okA || !okB {
			return false
		}

		return pathA.Vertex.ID == pathB.Vertex.ID
	}
	sortFunc := func(elemA interface{}, elemB interface{}) int {
		pathA := elemA.(PathPoint)
		pathB := elemB.(PathPoint)

		return pathA.WeightUpHere - pathB.WeightUpHere
	}
	return structdoublylinkedlist.NewSortedList(equalityFunc, sortFunc, reflect.TypeOf(PathPoint{}))
}

func (g *Graph) findShortestPath(currentPath PathPoint, toVerticeId string, paths map[string]PathPoint, sealedVertexes map[string]PathPoint) *PathPoint {
	sealedVertexes[currentPath.Vertex.ID] = currentPath

	pathList := getPathPointSortedList()
	for _, edge := range currentPath.Vertex.edgesAdjacentVertices {
		if _, sealed := sealedVertexes[edge.Head.ID]; sealed {
			continue
		}
		distance := currentPath.WeightUpHere + edge.Weight
		if estimate, ok := paths[edge.Head.ID]; !ok || distance < estimate.WeightUpHere {
			paths[edge.Head.ID] = PathPoint{
				Vertex:       currentPath.Vertex,
				WeightUpHere: distance,
			}
		}
		if edge.Head.ID == toVerticeId {
			path := paths[edge.Head.ID]
			return &path
		}
		pathList.Add(paths[edge.Head.ID])
	}
	for path := pathList.Unshift(); path != nil; path = pathList.Unshift() {
		cPath := path.(PathPoint)
		if finalPath := g.findShortestPath(cPath, toVerticeId, paths, sealedVertexes); finalPath != nil {
			return finalPath
		}
	}
	return nil
}
