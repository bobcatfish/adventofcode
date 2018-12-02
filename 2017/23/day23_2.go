package main

import "fmt"

func main() {
	bo := 109900
	c := 126900
	var primes []int
	var count int

	isPrime := true
	for b := bo; b <= c; b += 17 {
		for d := 2; d < b; d++ {
			if b%d == 0 {
				isPrime = false
			}
		}
		if isPrime {
			primes = append(primes, b)
		} else {
			count++
		}
		isPrime = true
	}

	fmt.Println(len(primes))
	fmt.Println(count)
}
