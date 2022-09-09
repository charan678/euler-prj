package main

import "fmt"

func euler01() {
	MUL3 := 3
	MUL5 := 5
	S := 0
	for i := 0; i <= 1000; i++ {
		if i % MUL3 == 0 || i % MUL5 == 0 {
			S += i
		}
	}

	fmt.Println(S)
}
