package varsOne

import "fmt"

func AllVars() {

	//! normal way to declare a variable
	var name string = "Aman Singh"
	fmt.Println(name)
	var currency float64 = 67.3123
	fmt.Println(currency)
	var decided bool = false
	fmt.Println(decided)

	//* go already knows type
	var money = 6700094273847135789
	fmt.Println(money)
	var version = "1.22.1"
	fmt.Println(version)

	//? short hand notation for declaring variable
	age := 33
	fmt.Println(age)

	//! difference between Printf()-> print format  & Println()-> print line
	fmt.Println("name", name, "person", age) // expected output --> "name""Aman Singh""age"133 // response --> name Aman Singh age 133
	// println always add space in string

	fmt.Printf("age is %d\n", age) // %d is int type formatter check
	// %f is float type formatter check
	// %.3f is float type formatter check but only for 3 value after point
	// %s is string type formatter check
	fmt.Printf("name is %s\n", name)
	// %T is typeOf formatter check
	fmt.Printf("version is %T\n", version)
	fmt.Printf("version is %q\n", version) // %q is a formatter for quoted string or any falue in quaoted form ""
}
