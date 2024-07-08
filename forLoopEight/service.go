package forloopeight

import "fmt"

func Sample() {
	fmt.Println("Go lang  has only one loop - forLoop")

	for i := 1; i <= 1; i++ {
		fmt.Println(i)
	}

	// for infinte loop like (while loop)
	counter := 0
	for {
		fmt.Println("loop started")
		counter++
		if counter == 2 {
			break
		}
	}

	// range keyword in for loop
	//! range keyword in forloop will go to value when it is on slices, arrays, object, interface, strings

	newSlice := []string{"aman", "singh", "rajawat", "anuj", "kumar", "mishra"}
	for ind, val := range newSlice {
		fmt.Println(ind, val)
	}

	//
	data := "hello world"
	for index, value := range data {
		fmt.Printf("index is : %d and value is %c\n", index, value)
		fmt.Println(value) // this is showing unicode of the char
	}
}
