package main

func eular10() {
	U := 2000000
	DP := NthPrime(U)
	S := 0
	for i := 0 ; i < len(DP); i++ {
		if DP[i] {
			S += i
		}
	}
	println(S)
}

func main() {
	eular10()
}

