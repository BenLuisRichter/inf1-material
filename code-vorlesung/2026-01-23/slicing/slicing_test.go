package slicing

import "fmt"

func Example() {
	l1 := []int{11, 22, 33, 44, 55}
	s1 := l1[2:4]

	fmt.Println(l1)
	fmt.Println(s1)

	s1[0] = 555

	fmt.Println(l1)
	fmt.Println(s1)

	s1 = append(s1, 777)

	fmt.Println(l1)
	fmt.Println(s1)

	s1 = append(s1, 999)

	fmt.Println(l1)
	fmt.Println(s1)

	s1[0] = 333

	fmt.Println(l1)
	fmt.Println(s1)

	// Output:
}

func Foo(y int) {
	y += 3

	fmt.Println(y)
}

func ExampleFoo() {
	x := 42

	Foo(x)

	fmt.Println(x)

	// Output:
}

func Bar(a []int) {
	b := []int{4, 5}
	a = append(b, 6)

	fmt.Println(a)

	// Output:
}
func Example_bar() {
	a := []int{1, 2, 3}
	fmt.Println(a)

	Bar(a)

	fmt.Println(a)

	// Output:
}
