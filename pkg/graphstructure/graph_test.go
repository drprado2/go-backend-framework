package graphstructure

import "testing"

type GraphTestFixture struct {
	graph *Graph
}

func (fixture *GraphTestFixture) setup() {
	fixture.graph = NewGraph()
}

func (fixture *GraphTestFixture) teardown() {

}

func TestGraph_AddEmptyEdge(t *testing.T) {
	fixture := GraphTestFixture{}
	fixture.setup()
	defer fixture.teardown()

	edge := NewEdge(nil, nil, nil)
	fixture.graph.AddEdge(edge)

	edges := fixture.graph.GetEdges()
	vertices := fixture.graph.GetVertices()

	if len(vertices) != 0 {
		t.Errorf("Vertices length must be 0 got %v", len(vertices))
	}
	if len(edges) != 1 {
		t.Errorf("Edges length must be 1 got %v", len(edges))
	}
	addedEdge := edges[0]
	if addedEdge.ID != edge.ID || addedEdge.Tail != nil || addedEdge.Head != nil || addedEdge.GenericData != nil {
		t.Errorf("The edge added is invalid")
	}
}

func TestGraph_AddEdgeWithHead(t *testing.T) {
	fixture := GraphTestFixture{}
	fixture.setup()
	defer fixture.teardown()

	headVertice := NewVertice(nil)

	edge := NewEdge(headVertice, nil, nil)
	fixture.graph.AddEdge(edge)

	edges := fixture.graph.GetEdges()
	vertices := fixture.graph.GetVertices()

	if len(vertices) != 1 {
		t.Errorf("Vertices length must be 1 got %v", len(vertices))
	}
	addedVertice := vertices[0]
	if addedVertice.ID != headVertice.ID || addedVertice.GenericData != nil {
		t.Errorf("The vertice added is invalid")
	}

	if len(edges) != 1 {
		t.Errorf("Edges length must be 1 got %v", len(edges))
	}
	addedEdge := edges[0]
	if addedEdge.ID != edge.ID || addedEdge.Tail != nil || addedEdge.Head.ID != headVertice.ID || addedEdge.GenericData != nil {
		t.Errorf("The edge added is invalid")
	}
}

func TestGraph_AddEdgeWithTail(t *testing.T) {
	fixture := GraphTestFixture{}
	fixture.setup()
	defer fixture.teardown()

	tailVertice := NewVertice(nil)

	edge := NewEdge(nil, tailVertice, nil)
	fixture.graph.AddEdge(edge)

	edges := fixture.graph.GetEdges()
	vertices := fixture.graph.GetVertices()

	if len(vertices) != 1 {
		t.Errorf("Vertices length must be 1 got %v", len(vertices))
	}
	addedVertice := vertices[0]
	if addedVertice.ID != tailVertice.ID || addedVertice.GenericData != nil {
		t.Errorf("The vertice added is invalid")
	}

	if len(edges) != 1 {
		t.Errorf("Edges length must be 1 got %v", len(edges))
	}
	addedEdge := edges[0]
	if addedEdge.ID != edge.ID || addedEdge.Tail.ID != tailVertice.ID || addedEdge.Head != nil || addedEdge.GenericData != nil {
		t.Errorf("The edge added is invalid")
	}
}

func TestGraph_AddEdgeWithHeadTail(t *testing.T) {
	fixture := GraphTestFixture{}
	fixture.setup()
	defer fixture.teardown()

	tailVertice := NewVertice(nil)
	headVertice := NewVertice(nil)

	edge := NewEdge(headVertice, tailVertice, nil)
	fixture.graph.AddEdge(edge)

	edges := fixture.graph.GetEdges()
	vertices := fixture.graph.GetVertices()

	if len(vertices) != 2 {
		t.Errorf("Vertices length must be 2 got %v", len(vertices))
	}
	for _, vertice := range vertices {
		if (vertice.ID != tailVertice.ID && vertice.ID != headVertice.ID) || vertice.GenericData != nil {
			t.Errorf("The vertices added are invalid")
		}
	}

	if len(edges) != 1 {
		t.Errorf("Edges length must be 1 got %v", len(edges))
	}
	addedEdge := edges[0]
	if addedEdge.ID != edge.ID || addedEdge.Tail.ID != tailVertice.ID || addedEdge.Head.ID != headVertice.ID || addedEdge.GenericData != nil {
		t.Errorf("The edge added is invalid")
	}
}

func TestGraph_AddUpdateEdge(t *testing.T) {
	fixture := GraphTestFixture{}
	fixture.setup()
	defer fixture.teardown()

	tailVertice := NewVertice(nil)
	headVertice := NewVertice(nil)

	edge := NewEdge(headVertice, nil, nil)
	fixture.graph.AddEdge(edge)

	edges := fixture.graph.GetEdges()
	vertices := fixture.graph.GetVertices()

	if len(vertices) != 1 {
		t.Errorf("Vertices length must be 1 got %v", len(vertices))
	}
	if len(edges) != 1 {
		t.Errorf("Edges length must be 1 got %v", len(edges))
	}

	edge.Tail = tailVertice
	headVertice.GenericData = 25
	fixture.graph.AddEdge(edge)

	edges = fixture.graph.GetEdges()
	vertices = fixture.graph.GetVertices()

	if len(vertices) != 2 {
		t.Errorf("Vertices length must be 2 got %v", len(vertices))
	}
	if len(edges) != 1 {
		t.Errorf("Edges length must be 1 got %v", len(edges))
	}

	findTailVertice := false
	for _, vertice := range vertices {
		if vertice.ID == headVertice.ID && vertice.GenericData != nil {
			t.Errorf("The data of head vertice was change, new value: %v", vertice.GenericData)
		}
		if vertice.ID == tailVertice.ID {
			findTailVertice = true
		}
	}
	if !findTailVertice {
		t.Errorf("The tail vertice was not added")
	}
}

func TestGraph_Encapsulation(t *testing.T) {
	fixture := GraphTestFixture{}
	fixture.setup()
	defer fixture.teardown()

	tailVertice := NewVertice(nil)
	headVertice := NewVertice(nil)

	edge := NewEdge(headVertice, tailVertice, nil)
	fixture.graph.AddEdge(edge)

	tailVertice.GenericData = 30
	headVertice.GenericData = 50

	vertices := fixture.graph.GetVertices()

	for _, vertice := range vertices {
		if vertice.ID == tailVertice.ID && vertice.GenericData != nil {
			t.Errorf("The tail vertice is not encapsulated")
		}
		if vertice.ID == headVertice.ID && vertice.GenericData != nil {
			t.Errorf("The head vertice is not encapsulated")
		}
	}
}

func TestGraph_AddEdgeWithGenericData(t *testing.T) {
	fixture := GraphTestFixture{}
	fixture.setup()
	defer fixture.teardown()

	edgeDataType := struct {
		name string
		age  int
	}{
		"Teste",
		25,
	}

	edge := NewEdge(nil, nil, edgeDataType)
	fixture.graph.AddEdge(edge)

	edges := fixture.graph.GetEdges()

	addedEdge := edges[0]

	genericData := addedEdge.GenericData.(struct {
		name string
		age  int
	})

	if genericData.age != edgeDataType.age || genericData.name != edgeDataType.name {
		t.Errorf("The generic data is invalid")
	}
}

func TestGraph_AddVertice(t *testing.T) {
	fixture := GraphTestFixture{}
	fixture.setup()

	verticeDataType := struct {
		name string
		age  int
	}{
		"Teste",
		25,
	}

	vertice := NewVertice(verticeDataType)

	fixture.graph.AddVertice(vertice)

	vertices := fixture.graph.GetVertices()
	if len(vertices) != 1 {
		t.Errorf("Vertices length must be 1 got %v", len(vertices))
	}
	addedVertice := vertices[0]

	if addedVertice.ID != vertice.ID {
		t.Errorf("The vertice ID added is invalid, value: %s", addedVertice.ID.String())
	}

	genericData := addedVertice.GenericData.(struct {
		name string
		age  int
	})

	if genericData.age != verticeDataType.age || genericData.name != verticeDataType.name {
		t.Errorf("The generic data is invalid")
	}
}
