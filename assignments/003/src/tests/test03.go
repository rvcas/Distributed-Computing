package main

import(
  "ssdqueue"
  "fmt"
)

func test_three() {
  q := new(ssdqueue.Node)
  fmt.Println(q)
  q = nil
}
