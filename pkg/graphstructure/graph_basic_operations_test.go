package graphstructure

import (
	"github.com/google/uuid"
	"testing"
)

type GraphTestFixture struct {
	directedAcycleGraph *Graph
}

func (fixture *GraphTestFixture) setup() {
	fixture.directedAcycleGraph = NewDirectedAcyclicGraph(false)
}

func (fixture *GraphTestFixture) teardown() {

}

func TestGraph_AddEmptyEdge(t *testing.T) {
	fixture := GraphTestFixture{}
	fixture.setup()
	defer fixture.teardown()

	err := fixture.directedAcycleGraph.AddEdge(uuid.New().String(), uuid.New().String(), 0, nil)
	if err == nil || err.Error() != "The vertice From and vertice To must be added to the directedAcycleGraph before creating the edge" {
		t.Errorf("Error must be not null got: %s", err)
	}
}

func TestGraph_AddEdgeWithHead(t *testing.T) {
	fixture := GraphTestFixture{}
	fixture.setup()
	defer fixture.teardown()

	headVertice := NewVertice(nil)
	fixture.directedAcycleGraph.AddVertice(*headVertice)

	err := fixture.directedAcycleGraph.AddEdge(uuid.New().String(), headVertice.ID, 0, nil)
	if err == nil || err.Error() != "The vertice From and vertice To must be added to the directedAcycleGraph before creating the edge" {
		t.Errorf("Error must be not null got: %s", err)
	}
}

func TestGraph_AddEdgeWithTail(t *testing.T) {
	fixture := GraphTestFixture{}
	fixture.setup()
	defer fixture.teardown()

	tailVertice := NewVertice(nil)

	fixture.directedAcycleGraph.AddVertice(*tailVertice)

	err := fixture.directedAcycleGraph.AddEdge(tailVertice.ID, uuid.New().String(), 0, nil)
	if err == nil || err.Error() != "The vertice From and vertice To must be added to the directedAcycleGraph before creating the edge" {
		t.Errorf("Error must be not null got: %s", err)
	}
}

func TestGraph_AddEdgeWithHeadTail(t *testing.T) {
	fixture := GraphTestFixture{}
	fixture.setup()
	defer fixture.teardown()

	tailVertice := NewVertice(nil)
	headVertice := NewVertice(nil)
	fixture.directedAcycleGraph.AddVertice(*tailVertice, *headVertice)

	fixture.directedAcycleGraph.AddEdge(tailVertice.ID, headVertice.ID, 0, nil)

	edges := fixture.directedAcycleGraph.GetEdges()
	vertices := fixture.directedAcycleGraph.GetVertices()

	if len(vertices) != 2 {
		t.Errorf("Vertices length must be 2 got %v", len(vertices))
	}
	for _, vertice := range vertices {
		if (vertice.ID != tailVertice.ID && vertice.ID != headVertice.ID) || vertice.GenericData != nil {
			t.Errorf("The vertexes added are invalid")
		}
	}

	if len(edges) != 1 {
		t.Errorf("Edges length must be 1 got %v", len(edges))
	}
	addedEdge := edges[0]
	if addedEdge.Tail.ID != tailVertice.ID || addedEdge.Head.ID != headVertice.ID || addedEdge.GenericData != nil {
		t.Errorf("The edge added is invalid")
	}
}

func TestGraph_UpdateEdgeData(t *testing.T) {
	fixture := GraphTestFixture{}
	fixture.setup()
	defer fixture.teardown()

	tailVertice := NewVertice(nil)
	headVertice := NewVertice(nil)
	fixture.directedAcycleGraph.AddVertice(*tailVertice, *headVertice)

	fixture.directedAcycleGraph.AddEdge(tailVertice.ID, headVertice.ID, 0, nil)

	edges := fixture.directedAcycleGraph.GetEdges()
	vertices := fixture.directedAcycleGraph.GetVertices()

	if len(vertices) != 2 {
		t.Errorf("Vertices length must be 2 got %v", len(vertices))
	}
	if len(edges) != 1 {
		t.Errorf("Edges length must be 1 got %v", len(edges))
	}

	fixture.directedAcycleGraph.UpdateEdgeData(tailVertice.ID, headVertice.ID, 29)

	edges = fixture.directedAcycleGraph.GetEdges()
	vertices = fixture.directedAcycleGraph.GetVertices()

	if len(vertices) != 2 {
		t.Errorf("Vertices length must be 2 got %v", len(vertices))
	}
	if len(edges) != 1 {
		t.Errorf("Edges length must be 1 got %v", len(edges))
	}

	updatedEdge := edges[0]
	if updatedEdge.GenericData != 29 {
		t.Errorf("Generic data after update must be 29 got %v", updatedEdge.GenericData)
	}
}

func TestGraph_Encapsulation(t *testing.T) {
	fixture := GraphTestFixture{}
	fixture.setup()
	defer fixture.teardown()

	tailVertice := NewVertice(nil)
	headVertice := NewVertice(nil)
	fixture.directedAcycleGraph.AddVertice(*tailVertice, *headVertice)

	fixture.directedAcycleGraph.AddEdge(tailVertice.ID, headVertice.ID, 0, nil)

	tailVertice.GenericData = 30
	headVertice.GenericData = 50

	vertices := fixture.directedAcycleGraph.GetVertices()

	for _, vertice := range vertices {
		if vertice.ID == tailVertice.ID && vertice.GenericData != nil {
			t.Errorf("The tail vertice is not encapsulated")
		}
		if vertice.ID == headVertice.ID && vertice.GenericData != nil {
			t.Errorf("The head vertice is not encapsulated")
		}
	}

	vertices[0].GenericData = 66

	vertices = fixture.directedAcycleGraph.GetVertices()

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

	tailVertice := NewVertice(nil)
	headVertice := NewVertice(nil)
	fixture.directedAcycleGraph.AddVertice(*tailVertice, *headVertice)

	fixture.directedAcycleGraph.AddEdge(tailVertice.ID, headVertice.ID, 0, edgeDataType)

	edges := fixture.directedAcycleGraph.GetEdges()

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

	fixture.directedAcycleGraph.AddVertice(*vertice)

	vertices := fixture.directedAcycleGraph.GetVertices()
	if len(vertices) != 1 {
		t.Errorf("Vertices length must be 1 got %v", len(vertices))
	}
	addedVertice := vertices[0]

	if addedVertice.ID != vertice.ID {
		t.Errorf("The vertice ID added is invalid, value: %s", addedVertice.ID)
	}

	genericData := addedVertice.GenericData.(struct {
		name string
		age  int
	})

	if genericData.age != verticeDataType.age || genericData.name != verticeDataType.name {
		t.Errorf("The generic data is invalid")
	}
}

func TestGraph_UpdateVerticeData(t *testing.T) {
	fixture := GraphTestFixture{}
	fixture.setup()

	verticeDataType := struct {
		name string
		age  int
	}{
		"Teste",
		25,
	}

	vertice := NewVertice(nil)
	fixture.directedAcycleGraph.AddVertice(*vertice)

	fixture.directedAcycleGraph.UpdateVerticeData(vertice.ID, verticeDataType)

	vertices := fixture.directedAcycleGraph.GetVertices()
	addedVertice := vertices[0]

	if addedVertice.ID != vertice.ID {
		t.Errorf("The vertice ID added is invalid, value: %s", addedVertice.ID)
	}

	genericData := addedVertice.GenericData.(struct {
		name string
		age  int
	})

	if genericData.age != verticeDataType.age || genericData.name != verticeDataType.name {
		t.Errorf("The generic data is invalid")
	}
}
