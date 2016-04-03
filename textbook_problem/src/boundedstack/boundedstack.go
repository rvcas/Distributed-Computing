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
  lock sync.Mutex
  Elements [SIZE]Element
}

func New() *BoundedStack {
  return new(BoundedStack).init()
}

func (s *BoundedStack) Push(v interface{}) (retval interface{}) {
  s.lock.Lock()
  defer s.lock.Unlock()

  if (s.full()) {
    return "FAILED STACK IS FULL"
  }

  retval, s.Elements[s.top+1].Value = v, v

  s.top++
  s.len++

  return
}

func (s *BoundedStack) Pop() (retval interface{}) {
  s.lock.Lock()
  defer s.lock.Unlock()

  if (s.empty()) {
    return "FAILED STACK IS EMPTY"
  }

  retval = s.Elements[s.top].Value

  s.Elements[s.top].Value = nil

  s.top--
  s.len--

  return
}

func (s *BoundedStack) Top() (retval interface{}) {
  s.lock.Lock()
  defer s.lock.Unlock()

  retval = s.Elements[s.top].Value

  return
}

func (s *BoundedStack) Len() int {
  s.lock.Lock()
  defer s.lock.Unlock()

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
