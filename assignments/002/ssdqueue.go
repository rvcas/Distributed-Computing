// Lucas Rosa
// lu599305
// Self-Stabilizing Distributed Queuing
package ssdqueue

type Node struct {
    next *Node
    data int
}

type Queue struct {
    head *Node
    tail *Node
    len int
}
