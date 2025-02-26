package main

import (
	"fmt"
	"math"
	"sync"
)

func intGenerator(wg *sync.WaitGroup, stop <-chan bool) <-chan int64 {
	intStream := make(chan int64)
	go func() {
		defer func() { wg.Done() }()
		defer close(intStream)
		var i int64
		for {
			i++
			select {
			case <-stop:
				return
			case intStream <- i:
			}
		}
	}()
	return intStream
}

// a) making the maresenne generator for the pipeline and also waiting properly for 15
func mersennegen(inte <-chan int64, stop chan bool) <-chan int64 {
	result := make(chan int64)
	go func() {
		defer close(result)
		for x := range inte {

			//here is a step to ensure that only the proper amount of numbers is generated before the generation is stopped

			if x == 16 {
				stop <- true
				break
			}
			smallres := int64(math.Pow(2, float64(x)) - 1)
			result <- smallres

		}

	}()
	return result

}

func main() {
	stop := make(chan bool)
	defer close(stop)

	var wg sync.WaitGroup
	wg.Add(1)

	intChannel := intGenerator(&wg, stop)

	resultfinale := mersennegen(intChannel, stop)

	for mercenne := range resultfinale {
		fmt.Printf("%d ", mercenne)
	}
	// stop <- true // stop the thread
	wg.Wait() // synchronisation
}
