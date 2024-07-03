package pointerselevan

import "fmt"

func updateValueByRef(a *int) {
	*a = *a + 105
}

func Sample() {
	fmt.Println("Hello from pointers")

	//!
	// var num int
	// num = 2

	// var ptr *int
	// ptr = &num

	num := 2
	ptr := &num

	fmt.Println(num)
	fmt.Println("what ptr have", ptr)      // this is the only address
	fmt.Println("what ptr contains", *ptr) // if you want to use value of ptr then use astric(*)

	//! if want to update original value through pointer address then-
	janam := 101
	updateValueByRef(&janam)                  // remember we only passed address in this func
	fmt.Println("changed janam value", janam) //? 101 to 206
}
