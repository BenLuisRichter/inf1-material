package rekursion

import "fmt"

func Ack(m, n int) int {
	if m == 0 {
		return n + 1
	}
	if n == 0 {
		return Ack(m-1, 1)
	}
	return Ack(m-1, Ack(m, n-1))
}

func ExampleAck() {
	// m == 0
	fmt.Println(Ack(0, 0))
	fmt.Println(Ack(0, 1))
	fmt.Println(Ack(0, 2))

	// n == 0
	fmt.Println(Ack(1, 0))
	fmt.Println(Ack(2, 0))
	fmt.Println(Ack(3, 0))

	// Diagonale
	fmt.Println(Ack(1, 1))
	fmt.Println(Ack(2, 2))
	fmt.Println(Ack(3, 3))
	fmt.Println(Ack(3, 15))

	// Output:
}
