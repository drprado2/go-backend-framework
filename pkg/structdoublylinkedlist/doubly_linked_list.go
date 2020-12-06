package structdoublylinkedlist

import (
	"fmt"
	"reflect"
)

type List struct {
	lenght             int
	head               *Node
	equalityComparator func(elementA interface{}, elementB interface{}) bool
	sortComparator     func(elementA interface{}, elementB interface{}) int
	isOrdered          bool
	dataType           reflect.Type
}

type Node struct {
	next, prev *Node
	data       interface{}
}

func NewList(
	equalityComparator func(elementA interface{}, elementB interface{}) bool,
	dataType reflect.Type,
) *List {
	return &List{
		lenght:             0,
		head:               nil,
		equalityComparator: equalityComparator,
		sortComparator:     nil,
		isOrdered:          false,
		dataType:           dataType,
	}
}

func NewSortedList(
	equalityComparator func(elementA interface{}, elementB interface{}) bool,
	sortComparator func(elementA interface{}, elementB interface{}) int,
	dataType reflect.Type,
) *List {
	return &List{
		lenght:             0,
		head:               nil,
		equalityComparator: equalityComparator,
		sortComparator:     sortComparator,
		isOrdered:          true,
		dataType:           dataType,
	}
}

type ListIterator struct {
	Current     interface{}
	currentNode *Node
	list        *List
}

func (i *ListIterator) Next() bool {
	if i.list.lenght == 0 {
		return false
	}
	if i.currentNode == nil {
		i.Current = i.list.head.data
		i.currentNode = i.list.head
		return true
	}
	if i.list.head != i.currentNode.next {
		i.Current = i.currentNode.next.data
		i.currentNode = i.currentNode.next
		return true
	}
	return false
}

func (l *List) checkElementType(element interface{}) error {
	elemType := reflect.TypeOf(element)
	if elemType != l.dataType {
		return fmt.Errorf("The type %v of the element is different of the type %v of the list", elemType, l.dataType)
	}
	return nil
}

func (l *List) addOrdered(element interface{}) {
	for current := l.head; ; current = current.next {
		if l.sortComparator(element, current.data) < 0 {
			node := &Node{
				next: current,
				prev: current.prev,
				data: element,
			}
			current.prev.next = node
			current.prev = node
			if current == l.head {
				l.head = node
			}
			return
		}
		if current.next == l.head {
			break
		}
	}

	l.addLast(element)
}

func (l *List) addLast(element interface{}) {
	node := &Node{
		next: l.head,
		prev: l.head.prev,
		data: element,
	}
	l.head.prev.next = node
	l.head.prev = node
}

func (l *List) addHead(element interface{}) {
	node := &Node{
		next: nil,
		prev: nil,
		data: element,
	}
	node.next = node
	node.prev = node
	l.head = node
}

func (l *List) Add(elements ...interface{}) error {
	for _, element := range elements {
		if err := l.checkElementType(element); err != nil {
			return err
		}

		l.lenght++

		if l.head == nil {
			l.addHead(element)
		} else if l.isOrdered {
			l.addOrdered(element)
		} else {
			l.addLast(element)
		}
	}

	return nil
}

func (l *List) removeNode(node *Node) {
	l.lenght--
	if l.head == node {
		l.head = nil
		return
	}

	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *List) Exists(element interface{}) bool {
	for current := l.head; ; current = current.next {
		if l.equalityComparator(current.data, element) {
			return true
		}
		if current.next == l.head {
			return false
		}
	}
}

func (l *List) Remove(element interface{}) bool {
	for current := l.head; ; current = current.next {
		if l.equalityComparator(current.data, element) {
			l.removeNode(current)
			return true
		}
		if current.next == l.head {
			return false
		}
	}
}

func (l *List) Unshift() interface{} {
	if l.lenght == 0 {
		return nil
	}
	l.lenght--
	if l.head.next == l.head {
		result := l.head
		l.head = nil
		return result.data
	}
	result := l.head
	l.head.prev.next = l.head.next
	l.head.next.prev = l.head.prev
	l.head = l.head.next
	return result.data
}

func (l *List) Pop() interface{} {
	if l.lenght == 0 {
		return nil
	}
	l.lenght--
	if l.head.next == l.head {
		result := l.head
		l.head = nil
		return result.data
	}
	result := l.head.prev
	l.head.prev.prev.next = l.head
	l.head.prev = l.head.prev.prev
	return result.data
}

func (l *List) ToIterator() *ListIterator {
	return &ListIterator{
		Current:     nil,
		currentNode: nil,
		list:        l,
	}
}

func (l *List) Lenght() int {
	return l.lenght
}
