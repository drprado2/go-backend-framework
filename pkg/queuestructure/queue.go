package queuestructure

type Queue struct {
	elements []interface{}
}

func NewQueue(startQueueCapacity int) *Queue {
	return &Queue{
		elements: make([]interface{}, 0, startQueueCapacity),
	}
}

func (q *Queue) Length() int {
	return len(q.elements)
}

func (q *Queue) Enqueue(element interface{}) {
	q.elements = append(q.elements, element)
}

func (q *Queue) PositionOfElement(element interface{}, comparator func(elementA interface{}, elementB interface{}) bool) int {
	for pos, elem := range q.elements {
		if comparator(element, elem) {
			return pos
		}
	}
	return -1
}

func (q *Queue) Next() interface{} {
	if len(q.elements) == 0 {
		return nil
	}

	element := q.elements[0]
	q.elements = q.elements[1:]
	return element
}
