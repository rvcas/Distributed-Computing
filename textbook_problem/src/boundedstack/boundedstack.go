// Lucas Rosa
package boundedstack

import (
  "sync"
  "time"
)

const (
  SIZE = 1000000
  TIMEOUT = time.Microsecond
)

type BoundedStack struct {
  head int
  len int
  waiters waiters
  lock sync.Mutex
  elements [SIZE]interface{}
}

func New() *BoundedStack {
  return new(BoundedStack).init()
}

func (s *BoundedStack) Push(v interface{}) error {
  s.lock.Lock()

  if (s.full() == ErrFull) {
    sema := newSema()
    s.waiters.put(sema)

    s.lock.Unlock()

    var timeoutC <-chan time.Time

    if TIMEOUT > 0 {
      timeoutC = time.After(TIMEOUT)
    }

    select {
    case <-sema.ready:
      s.elements[s.top+1] = v
      s.top++
      s.len++

      sema.response.Done()

      return nil
    case <-timeoutC:
      select {
      case sema.ready <- true:
        //
      default:
        sema.response.Done()
      }

      return ErrTimeout
    }
  }

  s.elements[s.top+1] = v

  s.top++
  s.len++

  for {
    sema := s.waiters.get()

    if sema == nil {
      break
    }

    sema.response.Add(1)

    select {
    case sema.ready <- true:
      sema.response.Wait()
    default:
      // This semaphore timed out.
    }

    if s.empty() == ErrEmpty {
      break
    }
  }

  s.lock.Unlock()

  return nil
}

func (s *BoundedStack) Pop() (interface{}, error) {
  s.lock.Lock()

  if s.empty() == ErrEmpty {
    sema := newSema()
    s.waiters.put(sema)

    s.lock.Unlock()

    var timeoutC <-chan time.Time

    if TIMEOUT > 0 {
      timeoutC = time.After(TIMEOUT)
    }

    select {
    case <-sema.ready:
      retval := s.elements[s.top]

      s.elements[s.top] = nil

      s.top--
      s.len--

      sema.response.Done()

      return retval, nil
    case <-timeoutC:
      select {
      case sema.ready <- true:
        //
      default:
        sema.response.Done()
      }

      return nil, ErrTimeout
    }
  }

  retval := s.elements[s.top]

  s.elements[s.top] = nil

  s.top--
  s.len--

  for {
    sema := s.waiters.get()

    if sema == nil {
      break
    }

    sema.response.Add(1)

    select {
    case sema.ready <- true:
      sema.response.Wait()
    default:
      // This semaphore timed out.
    }

    if s.full() == ErrFull {
      break
    }
  }

  s.lock.Unlock()

  return retval, nil
}

func (s *BoundedStack) Peek() (interface{}, error) {
  s.lock.Lock()
  defer s.lock.Unlock()

  if s.empty() != nil {
    return s.elements[s.top], nil
  }

  return nil, ErrEmpty
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

func (s *BoundedStack) full() error {
  if s.len == SIZE {
    return ErrFull
  } else {
    return nil
  }
}

func (s *BoundedStack) empty() error {
  if s.len == 0 {
    return ErrEmpty
  } else {
    return nil
  }
}
