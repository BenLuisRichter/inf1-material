package main

import "fmt"

// Rekursive Funktion zur Fakultätsberechnung
func faktoriael(n int) int {
	if n == 0 { // Basisbedingung
		return 1
	}
	return n * faktoriael(n-1) // Rekursiver Aufruf
}

func countDown(number int) {

	if number > 0 {
		fmt.Println(number)

		// recursive call
		countDown(number - 1)

	} else {
		// ends the recursive function
		fmt.Println("Countdown Stops")
	}

}

func main() {

	fmt.Println(faktoriael(5)) // Ausgabe: 120

	countDown(3)

	//for i := 1; i <= 3; i++ {
	//	fmt.Println("Line", i)
	//}

	names := []string{"Alice", "Bob", "Charlie"}

	//for _, name := range names {
	//	fmt.Println(name)
	//}

	//for i := 0; i < len(names); i++ {
	//	fmt.Println("value of i is:", i)
	//}

	//for i, v := range names {
	//	fmt.Printf("the value at index %v is %v \n", i, v)
	//}

	//for _, value := range names {
	//	fmt.Printf("the value is %v \n", value)
	//}

	names[1] = "Luigi"             //Tauscht "Bob" an Stelle 1 mit "Luigi"
	fmt.Println(names, len(names)) //Gibt Liste "names" und Länge von names aus

	var scores = []int{100, 50, 60} //initiiert slice scores mit type int
	//var ages = [3]int{20, 25, 30}		//array mit fester länge 3
	scores[2] = 25              //tauscht wert an stelle 2 (0,1,2,3,...) in scores mit 25
	scores = append(scores, 85) //beschreibt scores neu mit scores und 85 angehängt

	fmt.Println(scores, len(scores))

	rangeOne := names[1:3]  //gibt Teil von slice names von stelle 1 bis (ohne) 3 aus
	rangeTwo := names[2:]   //gibt Teil von slice names ab stelle 2 bis zum ende des slice aus
	rangeThree := names[:3] //-,,- bis (ohne) stelle 3
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "Cooper")
	fmt.Println(rangeOne)
}
