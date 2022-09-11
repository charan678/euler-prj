package main


func eular06() {
	N := 100
	Sqaure_of_sum_N_numbers := (N * (N + 1)) / 2
	sum_of_sqr_N_numbers := (N * (N + 1) * (2 * N + 1)) / 6

	print((Sqaure_of_sum_N_numbers * Sqaure_of_sum_N_numbers) - sum_of_sqr_N_numbers)

}


//func main() {
//	eular06()
//}
