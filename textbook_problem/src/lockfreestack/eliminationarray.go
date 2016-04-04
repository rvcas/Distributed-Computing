package lockfreestack

import (
  "math/rand"
  "time"
)

const (
  DURATION =  time.Millisecond()
)

type eliminationarray struct {
  exchangers []exchanger
  random *rand.Rand
}

func newEliminationArray(capacity int) *eliminationarray {
  exchangers = make([]exchanger, capacity, capacity)

  for exchanger := 0; exchanger < capacity; exchanger++ {
    exchangers[exchanger] = newExchanger()
  }

  random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func (elim *eliminationarray) visit(value T, chunk int) (T, error) {
  index := random.Intn(chunk)

  return (exchangers[index].exchange(value, DURATION)
}
