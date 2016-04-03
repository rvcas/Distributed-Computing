// Lucas Rosa
package boundedstack

import (
  "sync"
)

const (
  SIZE = 500000
)

type Element struct {
  Value interface{}
}

type BoundedStack struct {
  top int
  len int
  Lock sync.Mutex
  Elements [SIZE]Element
}

func New() *BoundedStack {
  return new(BoundedStack).init()
}

func (s *BoundedStack) Push(v interface{}) bool {
  if (s.full()) {
    return false
  }

  s.Elements[s.top+1].Value = v
  s.top++
  s.len++

  return true
}

func (s *BoundedStack) Pop() interface{} {
  if (s.empty()) {
    return "FAILED TO POP STACK IS EMPTY"
  }

  retval := s.Elements[s.top]

  s.top--
  s.len--

  return retval.Value
}

func (s *BoundedStack) Top() interface{} {
  return s.Elements[s.top].Value
}

func (s *BoundedStack) Len() int {
  return s.len
}

func (s *BoundedStack) init() *BoundedStack {
  s.top = -1
  s.len = 0

  return s
}

func (s *BoundedStack) full() bool {
  return (s.len == SIZE-1)
}

func (s *BoundedStack) empty() bool {
  return (s.len == 0)
}
