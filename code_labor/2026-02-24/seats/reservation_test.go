package seats

import "fmt"

func ExampleReservation_Occupied() {
	r1 := Reservation{
		id:          "42a",
		origin:      "Mannheim",
		destination: "Hamburg",
	}

	route := []string{
		"Karlsruhe",
		"Mannheim",
		"Frankfurt",
		"Kassel",
		"Göttingen",
		"Hamburg",
		"Kiel",
	}

	fmt.Println(r1.Occupied(0, route))
	fmt.Println(r1.Occupied(2, route))

	// Output:
	// false
	// true
}

func ExampleReservation_Text() {
	r1 := Reservation{
		id:          "42a",
		origin:      "Mannheim",
		destination: "Hamburg",
	}

	route := []string{
		"Karlsruhe",
		"Mannheim",
		"Frankfurt",
		"Kassel",
		"Göttingen",
		"Hamburg",
		"Kiel",
	}

	fmt.Println(r1.Text(0, route))
	fmt.Println(r1.Text(2, route))
	fmt.Println(r1.Text(5, route))

	// Output:
	// 42a: Mannheim - Hamburg
	// 42a: Mannheim - Hamburg
	//
}
