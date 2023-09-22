package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func multi(a int, b int) int {
	return a * b
}

func main() {
	result := plus(1, 2)
	fmt.Println("1+2=", result)

	result1 := multi(4, 6)
	fmt.Println("4*6:", result1)

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	c, _ := vals()
	fmt.Println(c)

	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

}

// ===================================================

// Multi Return Values

func vals() (int, int) {
	return 3, 8

}

// ======================================================

// Variadic Functions -->
// Variadic functions can be called with any number of trailing arguments. For example, fmt.Println is a common variadic function.

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
