package checksums

import "fmt"

// Sum erwartet eine Liste von Zahlen und berechnet deren Summe.
func Sum(numbers []int) int {
	result := 0
	for i, n := range numbers {
		if i%2 == 0 {
			result += n
		} else {
			result += 3 * n
		}

		// Als Einzeiler: result += (n + 2*n*(i+1)%2)
	}
	return result
}

// 978-0-345-39180-3

// Ean erwartet eine Liste von Ziffern und liefert die
// EAN-Prüfziffer dazu.
func Ean(digits []int) int {
	return (10 - (Sum(digits) % 10)) % 10
}

func ExampleEan() {
	fmt.Println(Ean([]int{9, 7, 8, 0, 3, 4, 5, 3, 9, 1, 8, 0}))

	// Output:
	// 3
}
