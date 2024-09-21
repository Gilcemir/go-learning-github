package main

import "fmt"

type base2 struct {
	num int
}
type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}
func (b base2) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base2
	base
	str string
}

func main() {

	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.base2.num, co.str)

	fmt.Println("also num:", co.base.num)

	fmt.Println("describe:", co.base2.describe())

	type describer interface {
		describe() string
	}

	var d describer = co.base2
	fmt.Println("describer:", d.describe())
}
