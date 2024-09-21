package main

import (
	"fmt"
	"reflect"
)

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func (r *rect) muteStruct(width, height int) {
	r.height = height
	r.width = width
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

	rp.muteStruct(20, 5)

	fmt.Println(rp)
	fmt.Println()
	fmt.Println(r)

	fmt.Print(reflect.TypeOf(r))
	fmt.Println()
	fmt.Print(reflect.TypeOf(rp))
}
