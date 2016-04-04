package lockfreestack

import (
  "unsafe"
  "sync/atomic"
  "time"
)

const (
  EMPTY = 0
  WAITING = 1
  BUSY = 2
)

type exchanger struct {
  slot *slot
}

func newExchanger() (e *exchanger) {
  return &exchanger{slot: newSlot()}
}

func (e *exchanger) exchange(myItem *LockFreeStack, timeout time.Duration) (*LockFreeStack, error) {
  nanos := timeout.Nanoseconds()
  timeBound := Time.UnixNano() + nanos
  stampHolder := []int{EMPTY}

  for true {
    if Time.UnixNano() > timeBound {
      return nil, ErrTimeout
    }

    yrItem := e.slot.get(stampHolder)
    stamp := stampHolder[0]

    switch stamp {
    case EMPTY:
      if e.slot.compareAndSet(yrItem, myItem, EMPTY, WAITING) {
        for Time.UnixNano() < timeBound {
          yrItem = e.slot.get(stampHolder)

          if stampHolder[0] == BUSY {
            e.slot.set(nil, EMPTY)

            return yrItem
          }
        }

        if e.slot.compareAndSet(yrItem, nil, WAITING, EMPTY) {
          return nil, ErrTimeout
        } else {
          yrItem = e.slot.get(stampHolder)
          e.slot.set(nil, EMPTY)

          return yrItem
        }
      }

      break
    case WAITING:
      if e.slot.compareAndSet(yrItem, myItem, WAITING, BUSY) {
        return yrItem
      }

      break
    case BUSY:
      break
    default:
      // claimed to be impossible by book
    }
  }
}
