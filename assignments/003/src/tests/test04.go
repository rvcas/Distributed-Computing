package main

import(
  "ssdqueue"
  "fmt"
)

func test_four() {
  q := new(ssdqueue.Node)
  fmt.Println(q)
  q = nil
}
