package main

import (
  "fmt"
)

func two_threads() {
  totalops := NUMOPS*TWOTHREADS

  fmt.Println("2 Threads - Test Results......")

  test_case(totalops, 10, 1, TWOTHREADS)
  test_case(totalops, 25, 2, TWOTHREADS)
  test_case(totalops, 50, 3, TWOTHREADS)
  test_case(totalops, 75, 4, TWOTHREADS)
  test_case(totalops, 90, 5, TWOTHREADS)
}
