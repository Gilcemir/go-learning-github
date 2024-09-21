package main

import (
	"fmt"
	"iter"
	"slices"
)

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {

		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

func (lst *List[T]) Filter(input func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if input(e.val) {
				if !yield(e.val) {
					return
				}
			}
		}
	}
}

func IsEven() func(resp int) bool {
	return func(arg int) bool {
		return arg%2 == 0
	}
}

func isEven(arg int) bool {
	return arg%2 == 0
}

func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1
		for {
			if !yield(a) {
				return
			}
			fmt.Println("a = ", a, " b = ", b)
			a, b = b, a+b
		}
	}
}

func main() {
	lst := List[int]{}
	lst.Push(10)
	lst.Push(12)
	lst.Push(13)
	lst.Push(23)
	lst.Push(18)

	for e := range lst.All() {
		fmt.Println(e)
	}
	fmt.Println("isEven")
	for e := range lst.Filter(isEven) {
		fmt.Print(e, " ")
	}

	fmt.Println("\n", "IsEven")
	for e := range lst.Filter(IsEven()) {
		fmt.Print(e, " ")
	}
	fmt.Println()
	all := slices.Collect(lst.All())
	fmt.Println("all:", all)

	for n := range genFib() {

		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
