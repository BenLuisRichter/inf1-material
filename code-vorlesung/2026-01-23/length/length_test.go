package length

import "fmt"

type Length int

func FromCentimeters(m int) Length {
	return Length(m)
}

func FromMeters(m int) Length {
	return Length(m * 100)
}

func (l Length) Centimeters() int {
	return int(l)
}

func (l Length) Meters() int {
	return l.Centimeters() / 100
}

func (l *Length) Scale(factor int) {
	//*l = Length(*l * Length(factor))
	*l = Length(l.Centimeters() * factor)
}

func Example() {
	//var a Length = 1000000
	a := FromMeters(10000)
	var b int = 2

	fmt.Println(a)
	a.Scale(b)
	fmt.Println(a)

	fmt.Println(a.Centimeters())
	fmt.Println(a.Meters())

	// Output:
}
