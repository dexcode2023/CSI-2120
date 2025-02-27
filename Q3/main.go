//Name: Dexter Day
//Student number: 300192263
//Q3: ~2 hours

package main

import (
	"fmt"
	"time"
)

//we have to define isPrime and getPrime in order to complete the getSpecialPrime function.
//they are defined below

// a special prime is a prime number that ends
// with the specified pattern sequence
// after nTrials the function returns with a false error code
func getSpecialPrime(pattern int64, maxValue int64, nTrials int) (int64, bool) {

	var div int64
	for div = 10; pattern/div != 0; div *= 10 {

	}

	for i := 0; i < nTrials; i++ {
		n := getPrime(maxValue)
		if n%div == pattern {
			return n, true // special prime found
		}
	}

	return 0, false // we failed to find a special prime
}

// This is the version without threads
func main() {

	fmt.Println("Solution without threads.")

	var number int64
	var sp []int64
	var ok bool

	var pattern int64 = 1111       // the suffix pattern we are looking for
	var maxV int64 = 1000000000000 // maximum value for the prime
	var trials int = 3             // number of trials for each attempts
	var nPrimes int = 4            // number of special primes we want

	start := time.Now()

	for len(sp) < nPrimes {
		number, ok = getSpecialPrime(pattern, maxV, trials)
		if ok { // we have found a special prime
			sp = append(sp, number)
		}
	}

	end := time.Now()

	fmt.Println("Special prime numbers are: ", sp)
	fmt.Println("End of program.", end.Sub(start))

}
