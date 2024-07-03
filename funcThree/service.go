package funcThree

import "fmt"

func add(a, b int) (int, error) {
	return a + b, nil
}

func sub(a, b int) (result int, err error) {
	result = a - b
	return result, nil
}

func Smaple() {
	val, err := add(4, 5)
	if err != nil {
		fmt.Println("err while adding")
	}
	fmt.Println(val)

	value, err := sub(4, 5)
	if err != nil {
		fmt.Println("err while adding")
	}
	fmt.Println(value)
}
