//Name: Dexter Day
//Student number: 300192263

package main

import (
	"fmt"
)

func sendInt() (ch chan int) {
	defer fmt.Println("Sender ready!")
	ch = make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			ch <- i
		}
		close(ch)
	}()
	return

}
func filterInt(src <-chan int) (dst chan int) {
	defer fmt.Println("Filter ready!")
	dst = make(chan int)
	go func() {
		defer close(dst)
		open := true
		for open {
			var i int
			i, open = <-src

			if open {
				i *= 2
				dst <- i
			}
		}
	}()
	return
}

func main() {
	src := sendInt()
	dst := filterInt(src)
	for i := range dst {

		fmt.Printf("%d ", i)
	}
	fmt.Println()
}
