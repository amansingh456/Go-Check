package stringsthirteen

import (
	"fmt"
	"strings"
)

func Smaple() {
	fmt.Println("hello world from string")
	data := "apple,banana,ornage"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	str := "one two three four five six seven ei8 Two" // i want to know how many times "two" repeated
	count := strings.Count(str, "two")
	fmt.Println(count)

	str = "   Hello AMAN    Singh"
	val := strings.TrimSpace(str)
	fmt.Println(val)

}
