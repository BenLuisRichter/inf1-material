package main

import "fmt"

func ExampleFactorial() {
	fmt.Println(Factorial(0))
	fmt.Println(Factorial(1))
	fmt.Println(Factorial(2))
	fmt.Println(Factorial(3))
	fmt.Println(Factorial(4))
	fmt.Println(Factorial(5))

	// Output:
	// 1
	// 1
	// 2
	// 6
	// 24
	// 120
}

func ExampleFactorialV2() {
	fmt.Println(FactorialV2(0))
	fmt.Println(FactorialV2(1))
	fmt.Println(FactorialV2(2))
	fmt.Println(FactorialV2(3))
	fmt.Println(FactorialV2(4))
	fmt.Println(FactorialV2(5))

	// Output:
	// 1
	// 1
	// 2
	// 6
	// 24
	// 120
}
