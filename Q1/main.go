//Name: Dexter Day
//Student number: 300192263
//Q3: ~45 mins

package main

import (
	"fmt"
	"sync"
)

// greatest common divisor
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

//a)
// func main() {
// 	x := []int{100, 15, 18, 10, 25, 90, 22, 60}
// 	y := []int{60, 5, 72, 5, 250, 108, 44, 3600}
// 	var z [8]int

// 	//adding done buffer
// 	done := make(chan bool, len(z))
// 	defer close(done)

// 	// parallel loop
// 	for i := range z {
// 		go func(i int) {
// 			z[i] = gcd(x[i], y[i])
// 			done <- true
// 		}(i)
// 	}

// 	for range z {
// 		<-done
// 	}

// 	// print result
// 	for _, v := range z {
// 		fmt.Println(v)
// 	}
// }

// b)
func main() {
	x := []int{100, 15, 18, 10, 25, 90, 22, 60}
	y := []int{60, 5, 72, 5, 250, 108, 44, 3600}
	var z [8]int

	//adding done buffer
	var wg sync.WaitGroup
	wg.Add(8)

	// parallel loop
	for i := range z {
		go func(i int) {
			z[i] = gcd(x[i], y[i])
			wg.Done()

		}(i)
	}
	wg.Wait()

	// print result
	for _, v := range z {
		fmt.Println(v)
	}
}
