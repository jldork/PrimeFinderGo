package PrimeGo

import (
	"fmt"
	"math"
)

func FindPrime(n int) int {
	var prime int

	for i, counter := 0, 0; counter < n; i++ {
		if IsPrime(i) {
			counter++
			prime = i
		}
		fmt.Println("Check if prime!")
		fmt.Println(i, IsPrime(i))

		fmt.Println("Number of primes: ", counter)
		fmt.Println("Latest prime: ", prime)
	}

	return prime
}

func IsPrime(n int) bool {
	for i := 2; float64(i) <= math.Sqrt(float64(n)); i++ {
		if n%i == 0 {
			return false
		}
	}
	if n <= 1 {
		return false
	} else {
		return true
	}
}
