// A concurrent prime sieve
package main

import (
	"fmt"
	"time"
)

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int, start int) {
	for i := start; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func main() {
	t1 := time.Now()
	sum := 0
	total := 0

	ch := make(chan int) // Create a new channel.
	go Generate(ch, 2)      // Launch Generate goroutine.
	for i := 0; i < 10; i++ {
		prime := <-ch
		sum += prime
		total++

		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}

	t2 := time.Now()
	fmt.Println("execution time:", t2.Sub(t1))
	fmt.Println("sum:", sum)
	fmt.Println("total number of primes:", total)
}
