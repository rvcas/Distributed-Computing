package main

import (
  "boundedstack"
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

  fmt.Println("# Bounded Stack Size:", boundedstack.SIZE)
  fmt.Println("# Each Thread Executes", NUMOPS, "Operations\n")

  one_thread()
  two_threads()
  four_threads()
  eight_threads()
}

func thread(s *boundedstack.BoundedStack, c chan int, threadID int, todo [NUMOPS]int) {
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

      c <- DONE
  }
}

func generate_tasks(percentpop int) [NUMOPS]int {
  var todo [NUMOPS]int

  r := rand.New(rand.NewSource(time.Now().UnixNano()))

  for task := 0; task < NUMOPS; task++ {
    randNum := r.Intn(101)

    if (randNum <= percentpop) {
      todo[task] = POP
    } else {
      todo[task] = PUSH
    }
  }

  return todo
}
