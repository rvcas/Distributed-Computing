// Lucas Rosa
// Self-Stabilizing distributed queuing package
package ssdqueue

// imports for mutex and formatting
import(
  "sync"
  "fmt"
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

// Queue Length
func (q *Queue) Len() int {
    return q.count
}

// add to the tail of a queue
func (q *Queue) Enqueue(v interface{}) {
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
func (q *Queue) Dequeue() {
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

// create a new queue to work with
func NewQueue() *Queue {
    return new(Queue)
}

// print queue info
func PrintQueueInfo(q *Queue) {
    fmt.Println("\tNumber of enqueues:", q.encount)
    fmt.Println("\tNumber of dequeues:", q.decount)
    fmt.Println("\tNodes in Queue:", q.Len())
}
