package main

import "fmt"

func euler02() {
	FOUR_MILLION := 4000000
	S := 0
	value := 1
	DP := []int{}
	DP  = append(DP, 0)
	DP = append(DP, 1)
	DP = append(DP, 2)
	for v:= 3 ; v <= 100 ;v++ {
		DP = append(DP, DP[v-1] + DP[v-2])
	}
	index := 1
	for index < 100  {
		value = DP[index]
		if value >= FOUR_MILLION {break}
		if value % 2 == 0 {
			S += value
		}
		index += 1
	}
	fmt.Printf("%d \n", S)
}
//
//func main() {
//	euler02()
//}
