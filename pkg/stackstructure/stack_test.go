package stackstructure

import "testing"

func TestStack_StackUp(t *testing.T) {
	stack := NewStack(0)

	elemA := 30
	elemB := 45

	stack.StackUp(elemA)
	stack.StackUp(elemB)

	stack.Unstack()

	stack.StackUp(elemA)
	stack.StackUp(elemB)
	stack.StackUp(elemB)

	if len(stack.elements) != 4 {
		t.Errorf("Lenght of stack must be 4 got %v", len(stack.elements))
	}

	if stack.elements[0] != elemA || stack.elements[1] != elemA || stack.elements[2] != elemB || stack.elements[3] != elemB {
		t.Errorf("Elements from stack are invalid")
	}
}

func TestStack_Unstack(t *testing.T) {
	stack := NewStack(0)

	elemA := 30
	elemB := 45

	stack.StackUp(elemA)
	stack.StackUp(elemB)

	elem1 := stack.Unstack()

	stack.StackUp(elemA)
	stack.StackUp(elemB)
	stack.StackUp(elemB)

	elem2 := stack.Unstack()
	elem3 := stack.Unstack()
	elem4 := stack.Unstack()
	elem5 := stack.Unstack()
	elem6 := stack.Unstack()

	if elem1 != elemB {
		t.Errorf("Element 1 must be %v got %v", elemB, elem1)
	}
	if elem2 != elemB {
		t.Errorf("Element 2 must be %v got %v", elemB, elem2)
	}
	if elem3 != elemB {
		t.Errorf("Element 3 must be %v got %v", elemB, elem3)
	}
	if elem4 != elemA {
		t.Errorf("Element 4 must be %v got %v", elemA, elem4)
	}
	if elem5 != elemA {
		t.Errorf("Element 5 must be %v got %v", elemA, elem5)
	}
	if elem6 != nil {
		t.Errorf("Element 6 must be nil got %v", elem6)
	}
}

func TestStack_Lenght(t *testing.T) {
	stack := NewStack(0)

	elemA := 30
	elemB := 45

	stack.StackUp(elemA)
	stack.StackUp(elemB)
	stack.StackUp(elemA)
	stack.StackUp(elemB)
	stack.StackUp(elemB)

	len := stack.Lenght()

	if len != 5 {
		t.Errorf("Lenght must be 5 got %v", len)
	}
}

type testType struct {
	name string
	age  int
}

func TestQueue_PositionOfElement(t *testing.T) {
	stack := NewStack(0)
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

	stack.StackUp(elemA)
	stack.StackUp(elemB)

	compareFunc := func(elemA interface{}, elemB interface{}) bool {
		castA := elemA.(testType)
		castB := elemB.(testType)

		if castA.name == castB.name {
			return true
		}
		return false
	}

	resultC := stack.PositionOfElement(elemC, compareFunc)
	resultB := stack.PositionOfElement(elemB, compareFunc)
	resultA := stack.PositionOfElement(elemA, compareFunc)

	if resultC != -1 {
		t.Errorf("Position of elementC must be -1 got %v", resultC)
	}
	if resultB != 0 {
		t.Errorf("Position of elementB must be 1 got %v", resultB)
	}
	if resultA != 1 {
		t.Errorf("Position of elementC must be 0 got %v", resultA)
	}
}
