package dataconversiontwelve

import (
	"fmt"
	"strconv"
)

func Sample() {
	fmt.Println("cheking data conversion")

	num := 42
	fmt.Printf("num is: %T and real type of num is: %d\n", num, num)

	// data := float64(num)
	data := strconv.Itoa(num)
	fmt.Println("data is ", data)
	fmt.Printf("data is %q\n", data)
	fmt.Printf("type of data is %T\n", data)

	number := "1234"
	numUpdate, _ := strconv.Atoi(number)
	fmt.Println("data is ", numUpdate)
	fmt.Printf("type of data is %T\n", numUpdate)

}
