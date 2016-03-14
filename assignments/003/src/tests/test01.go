package main

import(
  "ssdqueue"
  "fmt"
)

func test_one() {
  q := new(ssdqueue.Node)
  fmt.Println(q)
  q = nil
}
