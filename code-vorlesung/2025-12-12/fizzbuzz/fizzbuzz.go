package main

import "fmt"

func fizzbuzz_v1() {
	for i := 0; i < 42; i++ {
		if i%5 == 0 && i%7 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("fizz")
		} else if i%7 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func fizzbuzz_v2() {
	for i := 0; i < 42; i++ {
		if i%5 != 0 && i%7 != 0 {
			fmt.Println(i)
		} else if i%5 == 0 {
			fmt.Println("fizz")
		} else if i%7 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println("fizzbuzz")
		}
	}
}

func fizzbuzz_v3() {
	for i := 0; i < 42; i++ {
		if i%5 == 0 && i%7 == 0 {
			fmt.Println("fizzbuzz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("fizz")
			continue
		}
		if i%7 == 0 {
			fmt.Println("buzz")
			continue
		}
		fmt.Println(i)
	}
}

func fizzbuzz_v4() {
	for i := 0; i < 42; i++ {
		printed_something := false
		if i%5 == 0 {
			fmt.Print("fizz")
			printed_something = true
		}
		if i%7 == 0 {
			fmt.Print("buzz")
			printed_something = true
		}
		if !printed_something {
			fmt.Print(i)
		}
		fmt.Println()
	}
}

func fizzbuzz_v5() {
	for i := 0; i < 42; i++ {
		if i%5 != 0 && i%7 != 0 {
			fmt.Println(i)
			continue
		}
		if i%5 == 0 {
			fmt.Print("fizz")
		}
		if i%7 == 0 {
			fmt.Print("buzz")
		}
		fmt.Println()
	}
}

func fizzbuzz_allgemein(x, y, max int) {
	for i := 0; i < max; i++ {
		if i%x != 0 && i%y != 0 {
			fmt.Println(i)
			continue
		}
		if i%x == 0 {
			fmt.Print("fizz")
		}
		if i%y == 0 {
			fmt.Print("buzz")
		}
		fmt.Println()
	}
}

func main() {
	//fizzbuzz_v1()
	//fizzbuzz_v2()
	//fizzbuzz_v3()
	//fizzbuzz_v4()
	//fizzbuzz_v5()

	var a int
	fmt.Print("Wähle x: ")
	fmt.Scan(&a)

	var b int
	fmt.Print("Wähle y: ")
	fmt.Scan(&b)

	var m int
	fmt.Print("Bis wohin soll gezählt werden? ")
	fmt.Scan(&m)

	fizzbuzz_allgemein(a, b, m)
}
