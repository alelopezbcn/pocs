package iterators

type Set[E comparable] struct {
	m map[E]struct{}
}

func New[E comparable]() *Set[E] {
	return &Set[E]{
		m: make(map[E]struct{}),
	}
}

func (s *Set[E]) Add(value E) {
	if _, exists := s.m[value]; !exists {
		s.m[value] = struct{}{}
	}
}

func (s *Set[E]) Contains(v E) bool {
	_, exists := s.m[v]
	return exists
}

func Union[E comparable](s1, s2 *Set[E]) *Set[E] {
	result := New[E]()

	// we need to loop over the internal Set field m
	for v := range s1.m {
		result.Add(v)
	}
	for v := range s2.m {
		result.Add(v)
	}
	return result
}

func (s *Set[E]) Push(f func(E) bool) {
	for v := range s.m {
		if !f(v) {
			return
		}
	}
}
