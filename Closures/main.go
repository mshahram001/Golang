package main

import "fmt"

func intseq() func() int {
	i := 0
	return func() int {
		i += 2
		return i
	}
}

func main() {
	nextInt := intseq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := intseq()
	fmt.Println(newInt())
}
