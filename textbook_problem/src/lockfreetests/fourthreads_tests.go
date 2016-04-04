package main

import (
  "fmt"
)

func four_threads() {
  totalops := NUMOPS*FOURTHREADS

  fmt.Println("4 Threads - Test Results......")

  test_case(totalops, 10, 1, FOURTHREADS)
  test_case(totalops, 25, 2, FOURTHREADS)
  test_case(totalops, 50, 3, FOURTHREADS)
  test_case(totalops, 75, 4, FOURTHREADS)
  test_case(totalops, 90, 5, FOURTHREADS)
}
