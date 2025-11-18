package genericstack

type Stack[T any] struct {
	elements []T
}

func StackBuilder[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(v T) {
	s.elements = append(s.elements, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}

	last := s.elements[len(s.elements)-1]

	s.elements = s.elements[:len(s.elements)-1]

	return last, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}
