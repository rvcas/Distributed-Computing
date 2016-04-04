package main

import (
  "fmt"
)

func one_thread() {
  totalops := NUMOPS*ONETHREAD

  fmt.Println("1 Thread - Test Results......")

  test_case(totalops, 10, 1, ONETHREAD)
  test_case(totalops, 25, 2, ONETHREAD)
  test_case(totalops, 50, 3, ONETHREAD)
  test_case(totalops, 75, 4, ONETHREAD)
  test_case(totalops, 90, 5, ONETHREAD)
}
