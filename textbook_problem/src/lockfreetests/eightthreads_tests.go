package main

import (
  "fmt"
)

func eight_threads() {
  totalops := NUMOPS*EIGHTTHREADS

  fmt.Println("8 Threads - Test Results......")

  test_case(totalops, 10, 1, EIGHTTHREADS)
  test_case(totalops, 25, 2, EIGHTTHREADS)
  test_case(totalops, 50, 3, EIGHTTHREADS)
  test_case(totalops, 75, 4, EIGHTTHREADS)
  test_case(totalops, 90, 5, EIGHTTHREADS)
}
