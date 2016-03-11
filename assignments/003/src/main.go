package main

import(
  "ssdqueue"
)

func main() {
  q := ssdqueue.NewQueue();

  q.Enqueue(4)
  q.Enqueue(5)
  q.Dequeue()

  ssdqueue.PrintQueueInfo(q)
}
