// Lucas Rosa

// package must be main to run
package main

// imports for Mutex, formatting,
// time, & CPU count
import (
    "sync"
    "fmt"
    "time"
    "runtime"
)

// constants
const (
    NUMOPS = 500000
    DATA = 3
    DONE = 1
)

// nodes are elements of the queue
type Node struct {
    next *Node // next node: LL style
    Value interface{} // data
}

// the queue type definition
type Queue struct {
    head, tail *Node // head and tail: FIFO
    lock sync.Mutex // so memory doesn't get mad
    count int // node count
    encount int // enqueue count
    decount int // dequeue count
}

// spawns requested number of threads
func (q *Queue) test(numCPU int) {
    ops := NUMOPS*numCPU // total number of operations
    c := make(chan int, ops) // make channel buffered with ops

    // each thread is a function that recieves a count and a shared channel
    for i := 0; i < numCPU; i++ {
        go q.doStuff(c)
    }

    // Drain the channel
    for i := 0; i < ops; i++ {
        <-c
    }
}

// Each thread trys to execute 500000
// operations. Before each operation
// a lock must be acquired. After
// unlock a signal is sent to the channel.
func (q *Queue) doStuff(c chan int) {
    // locking to make Queue thread-safe
    for i := 0; i < NUMOPS; i++ {
        q.lock.Lock()

        // just a way to split up enqueue
        // and dequeue operations
        if (i+1)%2 == 0 {
            q.dequeue()
        } else {
            q.enqueue(DATA)
        }

        q.lock.Unlock()

        c <- DONE
    }
}

func (q *Queue) Len() int {
    return q.count
}

// add to the tail of a queue
func (q *Queue) enqueue(v interface{}) {
    if v == nil {
        panic("cannot add nil item to queue")
    }

    // Allocate space for a new node to add into the queue.
    // Set up our node to enqueue into the back of the queue.
    tmp := &Node{Value: v}

    // If the queue is NOT empty, we must set the old "last" node to point
    // to this newly created node.
    if q.tail != nil {
        q.tail.next = tmp
    }

    // Now, we must reset the back of the queue to our newly created node.
    q.tail = tmp

    // If the queue was previously empty we must ALSO set the front of the
    // queue.
    if q.head == nil {
        q.head = tmp
    }

    // update the queue's count
    q.count++
    q.encount++
}

// remove from the head of the queue
func (q *Queue) dequeue() {
    // Check the empty case.
    if q.head == nil {
        return
    }

    // Set up a temporary pointer to use to free the memory for this node.
    tmp := q.head

    // Make front point to the next node in the queue.
    q.head = q.head.next

    // If deleting this node makes the queue empty, we have to change the back
    // pointer also!
    if q.head == nil {
        q.tail = nil
    }

    // prevent memory leak
    tmp.next = nil
    tmp = nil

    // update queue count
    q.count--
    q.decount++
}

// main
func main() {
    menu()
}

// create a new queue to work with
func NewQueue() *Queue {
    return new(Queue)
}

// print queue info
func printQueueInfo(q *Queue) {
    fmt.Println("\tNumber of enqueues:", q.encount)
    fmt.Println("\tNumber of dequeues:", q.decount)
    fmt.Println("\tNodes in Queue:", q.Len())
}

func menu() {
    choice := 0
    quit := false

    for quit == false {
        fmt.Println("\n------------------")
        fmt.Println("Enter the number of threads you would like to run.\n" + "Each thread will execute 500000 operations on a Queue.")
        fmt.Println("For your consideration the number of CPUs on this machine is", runtime.NumCPU())
        fmt.Println("HIT ENTER TO QUIT!\n>")

        fmt.Scanf("%d", &choice)

        if choice == 0 {
            quit = true
        } else {
            q := NewQueue()

            t1 := time.Now()
            q.test(choice)
            t2 := time.Now()

            fmt.Println("\tResults:\n\t##########")
            fmt.Println("\tThread Count:", choice)
            fmt.Println("\tExecution Time:", t2.Sub(t1))
            printQueueInfo(q)

            q = nil
            choice = 0
        }
    }
}
