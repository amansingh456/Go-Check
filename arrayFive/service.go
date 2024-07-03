package arrayfive

import "fmt"

func Smaple() {
	fmt.Println("Checking Array")

	var name [3]string
	name[0] = "prince"
	name[1] = "akash"
	name[2] = "singh" // if not ful fil the array then it will store space

	fmt.Println(name)

	//? second method to declare array is
	numbers := [8]int{1, 2, 3, 4} // in number case it will store 0 on rest of the place, for boolean value it stores false and for pointers and complex types it stores nil
	fmt.Println(numbers)

	// find length
	fmt.Println(len(numbers))

}
