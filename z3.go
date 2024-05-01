package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getPrimeDivisors(11, 22))
}

func getPrimeDivisors(min int, max int) (res []int) {
	for i := min; i <= max; i++ {
		if isPrime(i) {
			res = append(res, i)
		}
	}

	return res
}

func isPrime(n int) bool {
	if n < 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
