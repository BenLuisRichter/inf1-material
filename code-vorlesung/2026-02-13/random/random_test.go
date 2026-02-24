package random

import "fmt"

func ExampleRandom_demo() {
	r := New(42)

	fmt.Println(r.Int())
	fmt.Println(r.Int())
	fmt.Println(r.Int())
	fmt.Println(r.Int())
	fmt.Println(r.Int())
	fmt.Println(r.Int())

	// Output:
	// 42
	// 76
	// 77
	// 92
	// 46
	// 11
}

func ExampleRandom_GetCircle() {
	r := New(42)
	c := r.GetCircle()
	fmt.Println(len(c))
	fmt.Println(c)

	// Output:
	// 15
	// [42 76 77 92 46 11 12 14 19 36 29 84 5 2 0]
}
