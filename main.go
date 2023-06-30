package main

import "fmt"

func main() {
	fmt.Printf("hello world")

	n := 7
	k := Fibo(n)
	fmt.Printf("%v", k)

	fmt.Printf("down")
}

func Fibo(n int) int {
	if n < 0 {
		return -1
	} else {
		f := Fibonacci()
		result := 0
		for i := 0; i < n; i++ {
			result = f()
		}
		return result
	}
}

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
