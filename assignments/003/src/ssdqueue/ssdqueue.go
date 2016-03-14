// Lucas Rosa
// Self-Stabilizing distributed queuing package
package ssdqueue

// imports for mutex
// import (
//   "sync"
// )

const (
  OBSERVE = false
  CORRECT = true
)

// nodes are elements of the queue
type Node struct {
  arrow *Node // next node: LL style
  find chan Node
  state bool // observe or correct
  sent int // number of finds
  phi_est int
}

func (node *Node) FindMessage() {

}

func (node *Node) ObserveMessage() {

}
