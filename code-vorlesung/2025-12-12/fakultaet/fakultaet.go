package main

func Factorial(n int) int {
	result := 1
	for i := n; i > 1; i-- {
		result = result * i
	}
	return result
}

func FactorialV2(n int) int {
	if n == 0 {
		return 1
	}
	return n * FactorialV2(n-1)
}
