package guessinggame

import (
	"fmt"
	"math/rand/v2"
)

func GuessingGame() {
	correct := rand.IntN(100) + 1
	for n := 0; n < 3; n++ {
		guess := ReadNumber()
		if guess == correct {
			fmt.Println("Richtig geraten :-)")
			return //zurückspringen zum Anfang
		}
	}

	fmt.Printf("Zu viele falsche Zahlen. Die richtige Zahl war %d!", correct)
}
