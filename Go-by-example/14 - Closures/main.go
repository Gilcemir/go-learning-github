package main

import (
	"fmt"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func intSeqArg(arg int) func() int {
	return func() int {
		arg++
		return arg
	}
}

func intSeqArgRef(arg *int) func() int {
	return func() int {
		*arg++
		return *arg
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	fmt.Println("----- /// -------")

	myClosure := intSeqArg(5)

	fmt.Println(myClosure())
	fmt.Println(myClosure())
	fmt.Println(myClosure())

	fmt.Println("----- /// -------")

	myInteger := 10

	myClosureWithReference := intSeqArgRef(&myInteger)

	fmt.Println(myClosureWithReference())
	fmt.Println(myClosureWithReference())
	fmt.Println(myClosureWithReference())

	fmt.Println("My Integer", myInteger)

}
