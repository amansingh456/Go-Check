package switchseven

import "fmt"

func Sample() {
	fmt.Println("Hello from switch file")

	day := 2

	switch day {
	case 1:
		{
			fmt.Println("MONDAY")
			break
		}
	case 2:
		{
			fmt.Println("TUESDAY")
			break
		}
	case 3:
		{
			fmt.Println("WEDNESDAY")
			break
		}
	default:
		{
			fmt.Println("SUNDAY")
			break
		}

	}

}
