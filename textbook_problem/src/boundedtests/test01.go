package main

import (
  "boundedstack"
  "fmt"
  "time"
)

func test_one() {
  s := boundedstack.New()
  threadCount := 1
  totalOps := NUMOPS*threadCount

  c := make(chan int, totalOps)

  t1 := time.Now()

  for thread := 1; thread <= threadCount; thread++ {
    go func() {
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
    }()
  }

  t2 := time.Now()

  for i := 0; i < totalOps; i++ {
      <-c
  }

  fmt.Println("1 Thread - Test Results:")
  fmt.Println("\tExecution Time:", t2.Sub(t1), "\n")

  s = nil
}
