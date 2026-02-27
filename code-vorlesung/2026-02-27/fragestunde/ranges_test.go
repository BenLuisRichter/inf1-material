package fragestunde

import "fmt"

func Example() {
	list := []int{5, 15, 25, 23, 35}

	// Jede Position zusammen mit ihrem Wert ausgeben.
	for i := 0; i < len(list); i++ {
		fmt.Printf("%v: %v | ", i, list[i])
	}
	fmt.Println()

	// Jede Position zusammen mit ihrem Wert ausgeben.
	for i := range list {
		fmt.Printf("%v: %v | ", i, list[i])
	}
	fmt.Println()

	// Jede Position zusammen mit ihrem Wert ausgeben.
	for i, v := range list {
		fmt.Printf("%v: %v | ", i, v)
	}
	fmt.Println()

	// Nur die Werte ausgeben.
	for _, v := range list {
		fmt.Printf("%v | ", v)
	}
	fmt.Println()

	// Output:
}

func Example_append_v1() {
	list := []int{0}

	for i := 0; i < len(list); i++ {
		list = append(list, i+1)
	}

	// Output:
}

func Example_append_v2() {
	list := []int{0}

	for i := range list {
		list = append(list, i+1)
	}
	fmt.Println(list)

	// Output:
}
