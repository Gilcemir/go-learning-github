package main

import "fmt"

func plus(a int, b int) int {

	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func plusPlusPlus(a, b, c, d int) (response int) {
	response = a + b + c + d
	return
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	res = plusPlusPlus(1, 2, 3, 4)
	fmt.Println("1 + 2 + 3 + 4 = ", res)
}
