// Lucas Rosa
package boundedstack

import (
  "sync"
)

const (
  SIZE = 1000000
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

func (s *BoundedStack) Push(v interface{}) interface{} {
  if (s.full()) {
    return "FAILED STACK IS FULL"
  }

  s.Elements[s.top+1].Value = v
  s.top++
  s.len++

  return v
}

func (s *BoundedStack) Pop() interface{} {
  if (s.empty()) {
    return "FAILED STACK IS EMPTY"
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
