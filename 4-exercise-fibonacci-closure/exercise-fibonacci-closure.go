package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fibo := 0
	fibo2 := 1
	return func() int {
		r := fibo
		temp := fibo2
		fibo2 = fibo + fibo2
		fibo = temp
		return r
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
