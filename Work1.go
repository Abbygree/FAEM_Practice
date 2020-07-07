package main

import "fmt"

func main() {
	println("\nВведите основание степени")
	var degree_base int64
	_, _ = fmt.Scanf("%d", &degree_base)
	println("\nВведите степень")
	var degree int
	_, _ = fmt.Scanf("%d", &degree)
	println()
	println(custom_pow(degree_base, degree))
}

func custom_pow(arg int64, power int) int64 {
	var argout int64 = 1
	for i := 0; i < power; i++ {
		argout *= arg
	}
	return argout
}
