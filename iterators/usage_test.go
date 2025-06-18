package iterators

import (
	"testing"
)

func TestPush(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(100)
	s.Add(101)
	s.Add(102)
	s.Add(1000)

	PrintAllElementsPush(s)
}

func TestPull(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(100)
	s.Add(101)
	s.Add(102)
	s.Add(1000)

	PrintAllElementsPull(s)
}
