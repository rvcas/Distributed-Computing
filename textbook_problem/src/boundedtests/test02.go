package main

import (
  "boundedstack"
  "fmt"
  "time"
)

func test_two() {
  s := boundedstack.New()
  threadCount := 2
  totalOps := NUMOPS*threadCount

  c := make(chan int, totalOps)

  t1 := time.Now()

  for threadID := 1; threadID <= threadCount; threadID++ {
    go thread(s, c, threadID)
  }

  for i := 0; i < totalOps; i++ {
      <-c
  }

  t2 := time.Now()

  fmt.Println("2 Threads - Test Results:")
  fmt.Println("\tExecution Time:", t2.Sub(t1))
  fmt.Println("\tStack Height:", s.Len(), "\n")

  s = nil
}
