package set

type Set[T AllowedTypes] struct {
	hashMap map[T]struct{}
}

func New[T AllowedTypes](capacity int) *Set[T] {
	return &Set[T]{
		hashMap: make(map[T]struct{}, capacity),
	}
}

func (s *Set[T]) Add(value T) {
	s.hashMap[value] = struct{}{}
}
func (s *Set[T]) In(value T) bool {
	_, in := s.hashMap[value]
	return in
}
func (s *Set[T]) All() []T {
	out := make([]T, 0, len(s.hashMap))
	for value := range s.hashMap {
		out = append(out, value)
	}
	return out
}
func (s *Set[T]) Len() int {
	return len(s.hashMap)
}

func (s *Set[T]) Intersect(other Set[T]) Set[T] {
	intersectionSet := New[T](max(other.Len(), s.Len()))
	a := s
	b := &other
	if s.Len() > b.Len() {
		swap(&a, &b)
	}
	for val := range a.hashMap {
		if b.In(val) {
			intersectionSet.Add(val)
		}
	}
	return *intersectionSet
}
func (s *Set[T]) Diff(other Set[T]) Set[T] {
	diffSet := New[T](max(other.Len(), s.Len()))
	for val := range s.hashMap {
		if !other.In(val) {
			diffSet.Add(val)
		}
	}
	for val := range other.hashMap {
		if !s.In(val) {
			diffSet.Add(val)
		}
	}
	return *diffSet
}

// Union returns a new Set containing all elements from both sets
func (s *Set[T]) Union(other Set[T]) Set[T] {
	unionSet := New[T](s.Len() + other.Len())
	for val := range s.hashMap {
		unionSet.Add(val)
	}
	for val := range other.hashMap {
		unionSet.Add(val)
	}
	return *unionSet
}

// Remove removes the specified value from the set
func (s *Set[T]) Remove(value T) {
	delete(s.hashMap, value)
}

// Clear removes all elements from the set
func (s *Set[T]) Clear() {
	s.hashMap = make(map[T]struct{})
}

func swap[T AllowedTypes](a, b **Set[T]) {
	t := *a
	*a = *b
	*b = t
}
