package utils

import "fmt"

func Print(a ...any) (n int, err error){
	return fmt.Println(a)
}

func PrintList(a ...any) (n int, err error){
	return fmt.Printf("%v \n",a)
}
