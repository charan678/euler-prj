package main

import "math"

func NthPrime(N int) []bool {
	DP := []bool{}
	for i :=0; i <= N ; i++ {
		DP = append(DP, true)
	}
	DP[0] = false
	DP[1] = false
	p := 2
	for p * p < N {
		for i := p * p ; i <= N; i += p {
			DP[i] = false
		}
		p += 1
	}
	return DP
}

func eular07()  {
	N := 10001
	P := 100.0
	for int(math.Ceil(float64(P / math.Log(P)))) < N * 3 {
		P *= 2
	}
	DP := NthPrime(int(P))
	count := 0
	for index, v := range DP {
		if v {
			count += 1
			if count == 10001 {
				println(index)
				break
			}
		}
	}
}

//func main() {
//	eular07()
//}
