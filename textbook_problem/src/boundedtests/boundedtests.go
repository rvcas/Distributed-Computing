package main

import (
  "boundedstack"
  "fmt"
)

const (
  NUMOPS = 500000
  DATA = 4
  DONE = 1
)

func main() {
  fmt.Println("# Bounded Stack Size:", boundedstack.SIZE)
  fmt.Println("# Each Thread Executes", NUMOPS, "Operations\n")

  test_one()
  test_two()
  test_three()
  test_four()
}

func thread(s *boundedstack.BoundedStack, c chan int, threadID int) {
  for count := 0; count < NUMOPS; count++ {
      s.Lock.Lock()

      if (count+1)%3 == 0 {
          s.Pop()
      } else {
          s.Push(DATA)
      }

      s.Lock.Unlock()

      c <- DONE
  }
}
