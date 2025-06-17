package main

import "github.com/alelopezbcn/pocs/iterators"

func main() {
	s := iterators.New[int]()
	s.Add(1)
	s.Add(100)
	s.Add(1000)

	iterators.PrintAllElementsPush(s)
}
