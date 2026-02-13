package duration

import "fmt"

type Duration int

// Seconds erwartet ein int, das als Sekunden interpretiert wird.
// Liefert eine Duration für diese Sekunden-Anzahl.
func Seconds(m int) Duration {
	return Duration(m)
}

// Minutes erwartet ein int, das als Minuten interpretiert wird.
// Liefert eine Duration für diese Minuten-Anzahl.
func Minutes(m int) Duration {
	return Duration(m * 60)
}

// ToSeconds liefert die Duration als Sekunden.
func (s Duration) ToSeconds() int {
	return int(s)
}

// ToMinutes liefert die Duration als Minuten.
func (s Duration) ToMinutes() int {
	return int(s) / 60
}

// SecondsToMinutes erwartet eine Sekunden-Anzahl
// und liefert die entsprechenden Minuten.
func SecondsToMinutes(s int) int {
	return int(Seconds(s).ToMinutes())
}

func Example() {
	fmt.Println(Seconds(55))
	fmt.Println(Minutes(360))
	fmt.Println(SecondsToMinutes(360))

	// Output:
	// 55
	// 21600
	// 6
}
