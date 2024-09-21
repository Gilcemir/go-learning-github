package main

import "fmt"

func sum(letter rune, nums ...int) {
	fmt.Print(nums, " ", string(letter), " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	sum('a', 1, 2)
	sum('b', 1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum('c', nums...)
}
