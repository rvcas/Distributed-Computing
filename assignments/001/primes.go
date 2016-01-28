//Lucas Rosa
//COP 4520
package main

import (
	"fmt"
    "os"
	"strconv"
	"time"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

// Doesn't return a prime literaly, but the least value to be used
// to filter out non prime numbers from an array or slice
func nextPrime(a []int, factor int) (newFactor int) {
    for _, v := range a {
        if v > factor {
            newFactor = v
            break
        }
    }
    return
}

func main() {
	t1 := time.Now()

	f, err := os.Create("output.txt")
    check(err)

    defer f.Close()

    // Create slice containing values from 0 to 1 million
    a := []int{}
    for i := 0; i < 1000000; i++ {
        a = append(a, i)
    }

    done := false
    factor := 1
    for !done {
        // Get the next factor to be used as a filter for slice `a`
        factor = nextPrime(a, factor)

        var modified bool
        // Filtering slice `a`
        for i, v := range a {
            if v%factor == 0 && v > factor {
                a[i] = 0 // This is equivalent to eliminating the value
                modified = true
            }
        }

        if !modified {
            done = true
        }
    }

    var sum int
	var total int
    for _, v := range a {
        sum += v
		if v != 0 {
			total += 1
		}
    }

    t2 := time.Now()

	n2, err := f.WriteString("Execution Time: " + t2.Sub(t1).String() + " Total:" + strconv.Itoa(total) + " Sum: " + strconv.Itoa(sum) + "\n" + "Top Ten Primes:\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	var count int
	for j := len(a) - 1; j >= 0; j-- {
		if a[j] != 0 {
			n2, err := f.WriteString(strconv.Itoa(a[j]) + "\n")
			check(err)
			fmt.Printf("wrote %d bytes\n", n2)
			count += 1

			if count >= 10 {
				break
			}
		}
	}

	f.Sync()
}
