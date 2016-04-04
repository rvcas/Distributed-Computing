package boundedstack

import "errors"

var (
  ErrEmpty = errors.New(`stack: EMPTY`)
  ErrFull = errors.New(`stack: FULL`)
  ErrTimeout = errors.New(`Waiting Timed Out`)
)
