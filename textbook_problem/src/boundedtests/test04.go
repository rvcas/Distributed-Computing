package main

import (
  "boundedstack"
  "fmt"
  "time"
)

func test_four() {
  s := boundedstack.New()
  threadCount := 8
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

  fmt.Println("8 Thread - Test Results:")
  fmt.Println("\tExecution Time:", t2.Sub(t1))
  fmt.Println("\tStack Height:", s.Len(), "\n")

  s = nil
}
