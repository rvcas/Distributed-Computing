package main

import(
  "ssdqueue"
  "fmt"
)

func test_two() {
  q := new(ssdqueue.Node)
  fmt.Println(q)
  q = nil
}
