package main

import (
	"github.com/euler-prj/utils"
	"math"
)

func isPrime(N int) bool {
	DP := []bool{}
	for i :=0; i <= N ; i++ {
		DP = append(DP, true)
	}
	p := 2
	for p * p < N {
		for i := p * p ; i <= N; i += p {
			DP[i] = false
		}
		p += 1
	}
	return DP[N]
}

func euler03() {
	N := 600851475143.0
	output := []int{}
	for i := 3; i <= int(math.Sqrt(N)); i += 2 {
		if int(N) % i == 0 {
			if isPrime(i) {
				output = append(output, i)
			}
		}
	}
	utils.Print(output[len(output)-1])
}

//func main() {
//	euler03()
//}
