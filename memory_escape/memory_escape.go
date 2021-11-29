package main

import "fmt"

type A struct {
	s string
}

func foo(s string) *A {
	a := new(A)
	a.s = s
	return a
}

func foo2() map[int]int {
	d := make(map[int]int)
	d[1] = 2
	delete(d, 1)
	return d
}

func main() {
	a := foo("hello")
	b := a.s + " world"
	c := b + "!"
	fmt.Println(c)
	d := foo2()
	fmt.Println(d)
}
