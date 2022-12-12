package utils

type Set[T comparable] struct {
	items map[T]bool
}

func (s *Set[T]) Add(t T) {
	if s.items == nil {
		s.items = make(map[T]bool)
	}
	s.items[t] = true
}

func (s *Set[T]) Remove(t T) {
	delete(s.items, t)
}

func (s *Set[T]) ToSlice() []T {
	result := make([]T, 0)
	for k := range s.items {
		result = append(result, k)
	}
	return result
}

func (s *Set[T]) AddSlice(slice []T) {
	if s.items == nil {
		s.items = make(map[T]bool)
	}
	for _, x := range slice {
		s.items[x] = true
	}
}

func (s *Set[T]) Contains(t T) bool {
	_, found := s.items[t]
	return found
}

func (s *Set[T]) Len() int {
	if s.items == nil {
		s.items = make(map[T]bool)
	}
	return len(s.items)
}
