package errorhandlingfour

import "fmt"

func devide(a, b float64) (val float64, err error) {
	if b == 0 {
		return 1, fmt.Errorf("Denomentor should be greater than 0")
	}
	val = a / b
	return val, nil
}

func Smaple() {
	fmt.Println("Hello from error handling")
	val, _ := devide(8, 2)
	// val, err := devide(9, 0)
	// if err != nil {
	// 	fmt.Println("error while deviding")
	// }
	fmt.Println(val)
}
