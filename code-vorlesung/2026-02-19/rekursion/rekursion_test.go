package rekursion

import "fmt"

func CountDown(n int) {
	if n <= 0 {
		return
	}
	fmt.Println(n)
	CountDown(n - 1)
}

func CountUp(n int) {
	if n <= 0 {
		return
	}
	CountUp(n - 1)
	fmt.Println(n)
}

func ExampleCountUp() {
	CountUp(3)

	// Output:
}

func CountDownUp(n int) {
	if n <= 0 {
		return
	}
	fmt.Println(n)
	CountDownUp(n - 1)
	fmt.Println(n)
}

func ExampleCountDownUp() {
	CountDownUp(3)

	// Output:
}

func CountUp2(n, i int) {
	if i == n {
		return
	}
	fmt.Println(i + 1)
	CountUp2(n, i+1)
}

func CountUp1(n int) {
	CountUp2(n, 0)
}

func ExampleCountUp1() {
	CountUp1(3)

	// Output:
}

var l []int

func CountDownListGlobal(n int) {
	if n <= 0 {
		return
	}

	l = append(l, n)
	CountDownListGlobal(n - 1)
}

func CountDownList(n int) []int {
	if n <= 0 {
		return []int{}
	}

	result := CountDownList(n - 1)
	result = append([]int{n}, result...)
	return result

	// return append([]int{n}, CountDownList(n-1)...)
}

func Example() {
	CountDown(3)

	//fmt.Println(l)
	CountDownListGlobal(5)
	CountDownListGlobal(5)
	fmt.Println(l)

	fmt.Println(CountDownList(7))

	// Output:
}
