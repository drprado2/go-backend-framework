package queuestructure

import "testing"

func TestQueue_Enqueue(t *testing.T) {
	queue := NewQueue(0)
	elemA := 12
	elemB := 20

	queue.Enqueue(elemA)
	queue.Enqueue(elemB)

	if queue.elements[0] != 12 {
		t.Errorf("Element at 0 must be 12 got %v", queue.elements[0])
	}
	if queue.elements[1] != 20 {
		t.Errorf("Element at 1 must be 20 got %v", queue.elements[1])
	}
}

func TestQueue_Next(t *testing.T) {
	queue := NewQueue(0)
	elemA := 12
	elemB := 20

	queue.Enqueue(elemA)
	queue.Enqueue(elemB)

	resultA := queue.Next()
	if len(queue.elements) != 1 {
		t.Errorf("Queue length must be 1 got %v", len(queue.elements))
	}

	resultB := queue.Next()
	if len(queue.elements) != 0 {
		t.Errorf("Queue length must be 0 got %v", len(queue.elements))
	}

	resultC := queue.Next()
	if len(queue.elements) != 0 {
		t.Errorf("Queue length must be 0 got %v", len(queue.elements))
	}

	if resultA != 12 {
		t.Errorf("Element at 0 must be 12 got %v", resultA)
	}
	if resultB != 20 {
		t.Errorf("Element at 1 must be 20 got %v", resultB)
	}
	if resultC != nil {
		t.Errorf("Element at 2 must be nil got %v", resultC)
	}
}

func TestQueue_Length(t *testing.T) {
	queue := NewQueue(0)
	elemA := 12
	elemB := 20

	queue.Enqueue(elemA)
	queue.Enqueue(elemB)

	len := queue.Length()
	if len != 2 {
		t.Errorf("Length must be 2 got %v", len)
	}

	queue.Next()
	len = queue.Length()
	if len != 1 {
		t.Errorf("Length must be 1 got %v", len)
	}
}

type testType struct {
	name string
	age  int
}

func TestQueue_PositionOfElement(t *testing.T) {
	queue := NewQueue(0)
	elemA := testType{
		name: "Adriano",
		age:  27,
	}
	elemB := testType{
		name: "Pedro",
		age:  36,
	}
	elemC := testType{
		name: "Carlos",
		age:  44,
	}

	queue.Enqueue(elemA)
	queue.Enqueue(elemB)

	compareFunc := func(elemA interface{}, elemB interface{}) bool {
		castA := elemA.(testType)
		castB := elemB.(testType)

		if castA.name == castB.name {
			return true
		}
		return false
	}

	resultC := queue.PositionOfElement(elemC, compareFunc)
	resultB := queue.PositionOfElement(elemB, compareFunc)
	resultA := queue.PositionOfElement(elemA, compareFunc)

	if resultC != -1 {
		t.Errorf("Position of elementC must be -1 got %v", resultC)
	}
	if resultB != 1 {
		t.Errorf("Position of elementB must be 1 got %v", resultB)
	}
	if resultA != 0 {
		t.Errorf("Position of elementC must be 0 got %v", resultA)
	}
}
