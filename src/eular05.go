package main

import (
	"math"
)

func primeFactors(N int) map[int]int {
	mapping := make(map[int]int)
	for N % 2 == 0 {
		if _, ok := mapping[2]; ok {
			mapping[2] += 1
		} else {
			mapping[2] = 1
		}
		N = N / 2
	}
	for i :=3; i <= int(math.Sqrt(float64(N))) ; i += 2 {
		for N % i == 0 {
			if _, ok := mapping[i]; ok {
				mapping[i] += 1
			} else {
				mapping[i] = 1
			}
			N = N / i
		}
	}
	if N > 2 {
		mapping[N] = 1
	}
	return mapping
}

func eular05() {
	primeMap := make(map[int]int)
	output := 1
	for i := 1; i <= 20; i++ {
		outputMap := primeFactors(i)
		for k, v := range outputMap {
			if _,ok := primeMap[k]; ok {
				if outputMap[k] > primeMap[k] {
					primeMap[k] = outputMap[k]
				}
			} else {
				primeMap[k] = v
			}
		}
	}
	for k, v := range(primeMap) {
		print(k, v)
		output *= int(math.Pow(float64(k) ,float64(v)))
	}
	println("Output =" , output)
}

//func main() {
//	eular05()
//}
