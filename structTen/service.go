package structten

import "fmt"

func Sample() {

	fmt.Println("Struct hello")

	//! 1st method
	aman := Person{
		Name:      "Aman Singh Rajawat",
		Age:       24,
		isMarried: false,
	}
	fmt.Printf("aman object is %+v\n", aman)

	//! 2nd method
	var aman2 Person

	aman2.Name = "Amit"
	aman2.isMarried = true
	aman2.Age = 26

	fmt.Printf("aman2 object is %+v\n", aman2)

	//! 3rd method (with new keyword)
	var aman3 = new(Person)
	aman3.Name = "Simran Agrawal"
	aman3.isMarried = true
	aman3.Age = 22

	fmt.Printf("aman3 object is %+v\n", aman3)

}
