package stack

import (
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	s := NewStack()

	log.Println(s.Empty())

	res, err := s.Pop()
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(res)
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	if s.Len() != 3 {
		t.Error()
	}

	a, _ := s.Pop()
	log.Println(a)
	a, _ = s.Pop()
	log.Println(a)
	/*
		a, _ = s.Pop()
		log.Println(a)
	*/

	b, e := s.Peek()
	if e != nil {
		log.Println(e.Error())
	} else {
		log.Println(b)
	}
}
