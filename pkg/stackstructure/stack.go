package stackstructure

type Stack struct {
	elements []interface{}
}

func NewStack(startCapacity int) *Stack {
	return &Stack{
		elements: make([]interface{}, 0, startCapacity),
	}
}

func (s *Stack) StackUp(element interface{}) {
	s.elements = append(s.elements, element)
}

func (s *Stack) Unstack() interface{} {
	if len(s.elements) == 0 {
		return nil
	}

	elem := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return elem
}

func (s *Stack) CopyToSlice() []interface{} {
	newSlice := make([]interface{}, 0, len(s.elements))
	for i, elem := range s.elements {
		newSlice[i] = elem
	}
	return newSlice
}

func (s *Stack) Lenght() int {
	return len(s.elements)
}

func (s *Stack) PositionOfElement(element interface{}, comparator func(elementA interface{}, elementB interface{}) bool) int {
	for pos, elem := range s.elements {
		if comparator(element, elem) {
			return len(s.elements) - pos - 1
		}
	}
	return -1
}
