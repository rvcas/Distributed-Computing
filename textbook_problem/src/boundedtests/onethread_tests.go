package main

import (
  "boundedstack"
  "fmt"
  "time"
)

func one_thread() {
  totalops := NUMOPS*ONETHREAD

  fmt.Println("1 Thread - Test Results......")

  one_thread_case_one(totalops)
  one_thread_case_two(totalops)
  one_thread_case_three(totalops)
  one_thread_case_four(totalops)
  one_thread_case_five(totalops)
}

func one_thread_case_one(totalops int) {
  todo := generate_tasks(10)

  s := boundedstack.New()

  c := make(chan int, totalops)

  t1 := time.Now()

  for threadID := 1; threadID <= ONETHREAD; threadID++ {
    go thread(s, c, threadID, todo)
  }

  for i := 0; i < totalops; i++ {
      <-c
  }

  t2 := time.Now()

  fmt.Println("\t- Case 1: 10% POP | 90% PUSH")
  fmt.Println("\t\tStack Height:", s.Len())
  fmt.Println("\t\tExecution Time:", t2.Sub(t1), "\n")


  s = nil
}

func one_thread_case_two(totalops int) {
  todo := generate_tasks(25)

  s := boundedstack.New()

  c := make(chan int, totalops)

  t1 := time.Now()

  for threadID := 1; threadID <= ONETHREAD; threadID++ {
    go thread(s, c, threadID, todo)
  }

  for i := 0; i < totalops; i++ {
      <-c
  }

  t2 := time.Now()

  fmt.Println("\t- Case 2: 25% POP | 75% PUSH")
  fmt.Println("\t\tStack Height:", s.Len())
  fmt.Println("\t\tExecution Time:", t2.Sub(t1), "\n")


  s = nil
}

func one_thread_case_three(totalops int) {
  todo := generate_tasks(50)

  s := boundedstack.New()

  c := make(chan int, totalops)

  t1 := time.Now()

  for threadID := 1; threadID <= ONETHREAD; threadID++ {
    go thread(s, c, threadID, todo)
  }

  for i := 0; i < totalops; i++ {
      <-c
  }

  t2 := time.Now()

  fmt.Println("\t- Case 3: 50% POP | 50% PUSH")
  fmt.Println("\t\tStack Height:", s.Len())
  fmt.Println("\t\tExecution Time:", t2.Sub(t1), "\n")


  s = nil
}

func one_thread_case_four(totalops int) {
  todo := generate_tasks(75)

  s := boundedstack.New()

  c := make(chan int, totalops)

  t1 := time.Now()

  for threadID := 1; threadID <= ONETHREAD; threadID++ {
    go thread(s, c, threadID, todo)
  }

  for i := 0; i < totalops; i++ {
      <-c
  }

  t2 := time.Now()

  fmt.Println("\t- Case 4: 75% POP | 25% PUSH")
  fmt.Println("\t\tStack Height:", s.Len())
  fmt.Println("\t\tExecution Time:", t2.Sub(t1), "\n")


  s = nil
}

func one_thread_case_five(totalops int) {
  todo := generate_tasks(90)

  s := boundedstack.New()

  c := make(chan int, totalops)

  t1 := time.Now()

  for threadID := 1; threadID <= ONETHREAD; threadID++ {
    go thread(s, c, threadID, todo)
  }

  for i := 0; i < totalops; i++ {
      <-c
  }

  t2 := time.Now()

  fmt.Println("\t- Case 5: 90% POP | 10% PUSH")
  fmt.Println("\t\tStack Height:", s.Len())
  fmt.Println("\t\tExecution Time:", t2.Sub(t1), "\n")


  s = nil
}
