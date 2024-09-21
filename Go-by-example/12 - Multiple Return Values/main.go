package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func vals2() (resp1, resp2 int) {
	resp1 = 5
	resp2 = 6

	return
}

func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	r1, r2 := vals2()

	fmt.Println(r1)
	fmt.Println(r2)
}
