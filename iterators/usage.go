package iterators

import "fmt"

func PrintAllElementsPush[E comparable](s *Set[E]) {
	funcForPush := func(v E) bool {
		fmt.Println(v)
		return true
	}

	s.Push(funcForPush)
}

func PrintAllElementsPull[E comparable](s *Set[E]) {
	next, stop := s.Pull()
	defer stop()
	for v, ok := next(); ok; v, ok = next() {
		fmt.Println("Reading value:", v)
	}
}
