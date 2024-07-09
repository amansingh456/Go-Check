package deferfourteen

import "fmt"

func Sample() {
	fmt.Println("Hello world from defer")

	//! defer keyword is used to late call for evrerfunc and it works on LIFO principal

	fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	defer fmt.Println(4)

}
