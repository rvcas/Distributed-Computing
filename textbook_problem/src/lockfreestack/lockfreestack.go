// Lucas Rosa
package lockfreestack

import (
  "unsafe"
  "sync/atomic"
)

type LockFreeStack struct {
  top int
  len int

}
