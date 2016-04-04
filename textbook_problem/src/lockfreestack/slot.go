package lockfreestack

import (
  "sync/time"
)

type slot struct {
  reference *LockFreeStack
  stamp int
}

func newSlot() (slot *slot) {
  return &slot{reference: nil, stamp: 0}
}

func (slot *slot) get(stampHolder []int) *LockFreeStack {
  // returns the current value of the reference
}

func (slot *slot) set(reference *LockFreeStack, stamp int) {
  //
}

func (slot *slot) compareAndSet(expected *LockFreeStack, new *LockFreeStack, exStamp int, newStamp int) bool {
  // Atomically sets the value of both the reference and stamp to the given update values if the current reference is == to the expected reference and the current stamp is equal to the expected stamp.

}
