package main

import (
	"fmt"
	"math"
)

func isPalindromeNumber(N int ) bool {
	E := N
	revN := 0
	for N != 0 {
		R := N % 10
		revN = revN * 10 + R
		N = N / 10
	}
	return revN == E
}

// https://projecteuler.net/problem=4
func euler04() {
	max_palindrome := 0.0
	for n := 999; n >= 100; n -= 1 {
		for m := 999; m >= 100; m -= 1 {
			if  isPalindromeNumber(n * m) {
				max_palindrome = math.Max(float64(max_palindrome), float64(n * m))
			}
		}
	}
	fmt.Println(max_palindrome)
}


//func main() {
//	euler04()
//}
