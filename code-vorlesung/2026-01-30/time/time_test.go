package time

import (
	"fmt"
	"time"
)

func Example() {
	now := time.Now().Unix()

	year := now/60/60/24/365 + 1970
	minute := (now / 60) % 60
	hour := (now / 60 / 60) % 24

	fmt.Println(now)
	fmt.Println(year)
	fmt.Println(minute)
	fmt.Println(hour)

	// Output:
}
