package iterators

import "fmt"

func PrintAllElementsPush[E comparable](s *Set[E]) {
	funcForPush := func(v E) bool {
		fmt.Println(v)
		return true
	}

	s.Push(funcForPush)
}
