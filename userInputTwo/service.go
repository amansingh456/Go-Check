package userInputTwo

import (
	"bufio"
	"fmt"
	"os"
)

func Sample() {
	fmt.Println("What is your name")
	var name string
	//? 1st method

	// _, err := fmt.Scan(&name) //fmt.Scan, with this there is a bug it only reads upto white space when it gets white space it will stop reading.
	// if err != nil {
	// 	fmt.Println("error while storgin")
	// }

	//? 2nd method
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // this means reader will read upto new line

	fmt.Println("hello mr.", name)
	fmt.Println("hello mr.", input)
}