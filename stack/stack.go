package stack

//基本数据结构 栈

import (
	"errors"
	"sync"
)

type Stack struct {
	stack []interface{}
	len   int
	lock  sync.Mutex
}

func NewStack() *Stack {
	s := &Stack{}
	s.stack = make([]interface{}, 0)
	s.len = 0
	return s
}

func (s *Stack) Len() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.len
}

func (s *Stack) Empty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.len == 0
}

func (s *Stack) Pop() (el interface{}, err error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.len == 0 {
		err = errors.New("stack is null")
	} else {
		el, s.stack = s.stack[0], s.stack[1:]
		s.len--
	}
	return
}

func (s *Stack) Push(el interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	prepend := make([]interface{}, 1)
	prepend[0] = el
	s.stack = append(prepend, s.stack...)
	s.len++
}

func (s *Stack) Peek() (interface{}, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.len == 0 {
		return nil, errors.New("stack is null")
	} else {
		return s.stack[0], nil
	}
}
