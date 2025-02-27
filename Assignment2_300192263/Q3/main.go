//Name: Dexter Day
//Student number: 300192263
//Q3: ~2 hours

package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

//we have to define isPrime and getPrime in order to complete the getSpecialPrime function.
//they are defined below

// isPrime to test if a number is prime using math.Sqrt optimization
func isPrime(n int64) bool {
	if n < 2 || n%2 == 0 {
		return false
	}
	if n == 2 {
		return true
	}
	n2 := int64(math.Sqrt(float64(n)))
	for i := int64(3); i <= n2; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// getPrime to generate below max val
func getPrime(maxValue int64) int64 {
	for {
		n := rand.Int63n(maxValue) // Generate a random number
		if isPrime(n) {
			return n
		}
	}
}

// a special prime is a prime number that ends
// with the specified pattern sequence
// after nTrials the function returns with a false error code
func getSpecialPrime(pattern int64, maxValue int64, nTrials int, stop <-chan bool, wg *sync.WaitGroup) <-chan int64 {

	out := make(chan int64)

	go func() {
		defer close(out)
		defer wg.Done()

		var div int64
		for div = 10; pattern/div != 0; div *= 10 {

		}
		for {
			select {
			case <-stop:
				return
			default:
				for i := 0; i < nTrials; i++ {
					n := getPrime(maxValue)
					if n%div == pattern {
						select {
						case out <- n:
							return
						case <-stop:
							return
						}
					}
				}

			}
		}
	}()
	return out
}

// This is the version without threads
func main() {

	fmt.Println("Solution with Threads.")

	var sp []int64

	var pattern int64 = 1111       // the suffix pattern we are looking for
	var maxV int64 = 1000000000000 // maximum value for the prime
	var trials int = 3             // number of trials for each attempts
	var nPrimes int = 4            // number of special primes we want

	start := time.Now()
	//waitgroup so we can sync threads
	var wg sync.WaitGroup

	//stop channel to end computation when numbers are found
	stop := make(chan bool)

	wg.Add(8)
	ch0 := getSpecialPrime(pattern, maxV, trials, stop, &wg)
	ch1 := getSpecialPrime(pattern, maxV, trials, stop, &wg)
	ch2 := getSpecialPrime(pattern, maxV, trials, stop, &wg)
	ch3 := getSpecialPrime(pattern, maxV, trials, stop, &wg)
	ch4 := getSpecialPrime(pattern, maxV, trials, stop, &wg)
	ch5 := getSpecialPrime(pattern, maxV, trials, stop, &wg)
	ch6 := getSpecialPrime(pattern, maxV, trials, stop, &wg)
	ch7 := getSpecialPrime(pattern, maxV, trials, stop, &wg)

	for len(sp) < nPrimes {
		select {
		case x := <-ch0:
			if x != 0 {
				sp = append(sp, x)
			}
		case x := <-ch1:
			if x != 0 {
				sp = append(sp, x)
			}
		case x := <-ch2:
			if x != 0 {
				sp = append(sp, x)
			}
		case x := <-ch3:
			if x != 0 {
				sp = append(sp, x)
			}
		case x := <-ch4:
			if x != 0 {
				sp = append(sp, x)
			}
		case x := <-ch5:
			if x != 0 {
				sp = append(sp, x)
			}
		case x := <-ch6:
			if x != 0 {
				sp = append(sp, x)
			}
		case x := <-ch7:
			if x != 0 {
				sp = append(sp, x)
			}
		}

	}
	close(stop)
	wg.Wait()
	end := time.Now()

	fmt.Println("Special prime numbers are: ", sp)
	fmt.Println("End of program.", end.Sub(start))

}
