package structdoublylinkedlist

import (
	"reflect"
	"testing"
)

type testElement struct {
	id   int
	name string
}

func testElementEquals(elemA interface{}, elemB interface{}) bool {
	castA, AOk := elemA.(testElement)
	castB, BOk := elemB.(testElement)

	if !AOk || !BOk {
		return false
	}

	return castA.id == castB.id
}

func testElementSortComparator(elemA interface{}, elemB interface{}) int {
	castA := elemA.(testElement)
	castB := elemB.(testElement)

	if castA.id < castB.id {
		return -1
	}
	if castA.id == castB.id {
		return 0
	}
	return 1
}

func TestList_Add(t *testing.T) {
	list := NewList(testElementEquals, reflect.TypeOf(testElement{}))
	elemA := testElement{
		id:   1,
		name: "Adriano",
	}
	elemB := testElement{
		id:   2,
		name: "Bruno",
	}
	list.Add(elemA, elemB)

	headId := list.head.data.(testElement).id
	if headId != elemA.id {
		t.Errorf("Head id must be 1 got %v", headId)
	}
	secondId := list.head.next.data.(testElement).id
	if secondId != elemB.id {
		t.Errorf("Second id must be 2 got %v", secondId)
	}
	if list.lenght != 2 {
		t.Errorf("List lenght must be 2 got %v", list.lenght)
	}
}

func TestList_AddOrdered(t *testing.T) {
	list := NewSortedList(testElementEquals, testElementSortComparator, reflect.TypeOf(testElement{}))
	elemA := testElement{
		id:   1,
		name: "Adriano",
	}
	elemB := testElement{
		id:   2,
		name: "Bruno",
	}
	elemC := testElement{
		id:   3,
		name: "Bruno",
	}
	elemD := testElement{
		id:   4,
		name: "Bruno",
	}
	list.Add(elemC, elemB, elemD, elemA)

	if list.lenght != 4 {
		t.Errorf("List lenght must be 2 got %v", list.lenght)
	}

	elementsExpected := []int{1, 2, 3, 4}
	iterator := list.ToIterator()
	for iterator.Next() {
		element := iterator.Current.(testElement)
		if element.id != elementsExpected[0] {
			t.Errorf("Element id shoud be %v got %v", elementsExpected[0], element.id)
		}
		elementsExpected = elementsExpected[1:]
	}
	if len(elementsExpected) > 0 {
		t.Errorf("Elements comparator should be empty")
	}
}

func TestList_AddDifferentType(t *testing.T) {
	list := NewSortedList(testElementEquals, testElementSortComparator, reflect.TypeOf(testElement{}))

	expectedError := "The type int of the element is different of the type structdoublylinkedlist.testElement of the list"
	if err := list.Add(23); err == nil || err.Error() != expectedError {
		t.Errorf("Error on add element shoud be %s got %v", expectedError, err.Error())
	}
}

func TestList_Remove(t *testing.T) {
	list := NewList(testElementEquals, reflect.TypeOf(testElement{}))
	elemA := testElement{
		id:   1,
		name: "Adriano",
	}
	elemB := testElement{
		id:   2,
		name: "Pedro",
	}

	list.Add(elemA)
	if list.Remove(elemB) {
		t.Error("Shoud not remove elementB")
	}
	if !list.Remove(elemA) {
		t.Error("Shoud remove elementA")
	}
	if list.lenght != 0 {
		t.Errorf("List lenght should be 0 got %v", list.lenght)
	}
	if list.head != nil {
		t.Errorf("Head list should be nil got %v", list.head.data)
	}
}

func TestList_Pop(t *testing.T) {
	list := NewList(testElementEquals, reflect.TypeOf(testElement{}))

	if elem := list.Pop(); elem != nil {
		t.Errorf("Pop in empty list shoud return nil got %v", elem)
	}
	if list.lenght != 0 {
		t.Errorf("List lenght shoud be 0 got %v", list.lenght)
	}

	elemA := testElement{
		id:   1,
		name: "Adriano",
	}
	elemB := testElement{
		id:   2,
		name: "Pedro",
	}
	list.Add(elemA)

	elemARemoved := list.Pop()
	if elemARemoved == nil {
		t.Error("Pop should not be nil")
	}
	if elemARemoved.(testElement).id != elemA.id {
		t.Errorf("Element A removed id should be 1 got %v", elemARemoved.(testElement).id)
	}
	if list.lenght != 0 {
		t.Errorf("List lenght shoud be 0 got %v", list.lenght)
	}

	list.Add(elemA)
	list.Add(elemB)

	firstPop := list.Pop()

	if firstPop == nil {
		t.Error("First pop should not be nil")
	}
	if firstPop.(testElement).id != elemB.id {
		t.Errorf("Element B removed id should be 2 got %v", firstPop.(testElement).id)
	}
	if list.lenght != 1 {
		t.Errorf("List lenght shoud be 1 got %v", list.lenght)
	}

	secondPop := list.Pop()

	if secondPop == nil {
		t.Error("Second pop should not be nil")
	}
	if secondPop.(testElement).id != elemA.id {
		t.Errorf("Element A removed id should be 1 got %v", secondPop.(testElement).id)
	}
	if list.lenght != 0 {
		t.Errorf("List lenght shoud be 0 got %v", list.lenght)
	}
}

func TestList_Unshift(t *testing.T) {
	list := NewList(testElementEquals, reflect.TypeOf(testElement{}))

	if elem := list.Unshift(); elem != nil {
		t.Errorf("Unshift in empty list shoud return nil got %v", elem)
	}
	if list.lenght != 0 {
		t.Errorf("List lenght shoud be 0 got %v", list.lenght)
	}

	elemA := testElement{
		id:   1,
		name: "Adriano",
	}
	elemB := testElement{
		id:   2,
		name: "Pedro",
	}
	list.Add(elemA)

	elemARemoved := list.Unshift()
	if elemARemoved == nil {
		t.Error("Pop should not be nil")
	}
	if elemARemoved.(testElement).id != elemA.id {
		t.Errorf("Element A removed id should be 1 got %v", elemARemoved.(testElement).id)
	}
	if list.lenght != 0 {
		t.Errorf("List lenght shoud be 0 got %v", list.lenght)
	}

	list.Add(elemA)
	list.Add(elemB)

	firstUnshift := list.Unshift()

	if firstUnshift == nil {
		t.Error("First Unshift should not be nil")
	}
	if firstUnshift.(testElement).id != elemA.id {
		t.Errorf("Element A removed id should be 1 got %v", firstUnshift.(testElement).id)
	}
	if list.lenght != 1 {
		t.Errorf("List lenght shoud be 1 got %v", list.lenght)
	}

	secondUnshift := list.Unshift()

	if secondUnshift == nil {
		t.Error("Second pop should not be nil")
	}
	if secondUnshift.(testElement).id != elemB.id {
		t.Errorf("Element B removed id should be 2 got %v", secondUnshift.(testElement).id)
	}
	if list.lenght != 0 {
		t.Errorf("List lenght shoud be 0 got %v", list.lenght)
	}
}

func TestList_Exists(t *testing.T) {
	list := NewList(testElementEquals, reflect.TypeOf(testElement{}))

	elemA := testElement{
		id:   1,
		name: "Adriano",
	}
	elemB := testElement{
		id:   2,
		name: "Pedro",
	}
	list.Add(elemA)

	if !list.Exists(elemA) {
		t.Error("Element A must exists")
	}
	if list.Exists(elemB) {
		t.Error("Element B must not exists")
	}
}
