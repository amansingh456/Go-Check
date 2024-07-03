package slicessix

import "fmt"

func Sample() {
	fmt.Println("Hello from slice")
	// slice is a flexible and dynamic data structure
	numb := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(numb)
	fmt.Println(len(numb))

	numb = append(numb, 8, 9)
	fmt.Println(numb)
	fmt.Println(cap(numb))
	fmt.Println(len(numb))

	// so basiclly when we make slice normally without make function/keyword then the length and capacity will be same. but if any how if you want that your capacity is more but for if you want to store less value then can use make func/keyword

	nums := make([]bool, 3, 10)
	fmt.Println(nums)
	fmt.Println(cap(nums)) // now this is 10
	fmt.Println(len(nums)) // this is 3

	// now aading one more value in slice
	nums = append(nums, true)
	fmt.Println(nums)
	fmt.Println(cap(nums))
	fmt.Println(len(nums))

	// so when length is equal to capacity then its fine, when length it is start adding more value then capacity then capacity bemcomes genrally double or little bit more than double or less then nbit double
	nums = append(nums, true, true, true, true, true, true, true)
	fmt.Println(nums)
	fmt.Println(cap(nums))
	fmt.Println(len(nums))

	var stock = []string{} // empty slice

	stick := make([]int, 3) // empty slice with make func, genrally we don't sue capacity so remove it.
	fmt.Println(stock, stick)
}
