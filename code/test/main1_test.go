package main

func Example() {
	i, j := 50, 2701

	p := &i //zeigt auf i
	//fmt.Println(*p) //gibt p aus, welcher auf i zeigt, also 42
	*p = 25 //setzt den pointer auf 21
	//fmt.Println(i)  //gibt durch p veränderte value von i aus, also 21

	p = &j       //p zeigt auf j
	*p = *p / 37 //setzt neuen p wert auf wert von p /37, also 2701/37 = 73
	//fmt.Println(j)

	type Person struct {
		Name string
		Age  int
	}

	//a := Person{Name: "Max", Age: 30}
	//fmt.Println(a.Name)

	//Output:
}
