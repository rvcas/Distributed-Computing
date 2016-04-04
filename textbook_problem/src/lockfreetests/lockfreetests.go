package main

import (
  "lockfreestack"
  "fmt"
  "math/rand"
  "os"
  "time"
)

func main() {
  if len(os.Args) > 1 {
    if (os.Args[1] == "-v" || os.Args[1] == "--verbose") {
      Verbose = true
    }
  }

  if !Verbose {
    fmt.Println("# Lock-Free Bounded Stack Size:", boundedstack.SIZE)
    fmt.Println("# Each Thread Executes", NUMOPS, "Operations\n")
  }

  one_thread()
  two_threads()
  four_threads()
  eight_threads()
}

func thread(s *lockfreestack.LockFreeStack, threadID int, todo *[NUMOPS]int) {
  for task := 0; task < NUMOPS; task++ {
    if todo[task] == POP {
      v, err := s.Pop()

      if Verbose {
        if err != nil {
          fmt.Println("Thread", threadID, ":", err)
        } else {
          fmt.Println("Thread", threadID, ": pop", v)
        }
      }
    } else {
      err := s.Push(DATA)

      if Verbose {
        if err != nil {
          fmt.Println("Thread", threadID, ":", err)
        } else {
          fmt.Println("Thread", threadID, ": push", DATA)
        }
      }
    }
  }
}

func test_case(totalops int, percentpop int, caseNum int, numThreads int) {
  todo := generate_tasks(percentpop)

  s := lockfreestack.New()

  c := make(chan int, totalops)

  t1 := time.Now()

  for threadID := 1; threadID <= numThreads; threadID++ {
    go thread(s, c, threadID, &todo)
  }

  for i := 0; i < totalops; i++ {
      <-c
  }

  t2 := time.Now()

  if !Verbose {
    fmt.Println("\t- Case", caseNum, ":", percentpop, "POP |", (100-percentpop), "PUSH")
    fmt.Println("\t\tStack Height :", s.Len())
    fmt.Println("\t\tExecution Time :", t2.Sub(t1), "\n")
  }


  s = nil
}

func generate_tasks(percentpop int) [NUMOPS]int {
  var todo [NUMOPS]int

  r := rand.New(rand.NewSource(time.Now().UnixNano()))

  for task := 0; task < NUMOPS; task++ {
    randNum := r.Intn(101)

    if randNum <= percentpop {
      todo[task] = POP
    } else if randNum <= 100 {
      todo[task] = PUSH
    } else {
      todo[task] = PUSH
    }
  }

  return todo
}
