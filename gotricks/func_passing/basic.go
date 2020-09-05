package main

import (
	"fmt"
)

type A struct {
	Foo string
}

func (a A) APrint() {
	fmt.Printf(a.Foo)
}

func simple(f func(a A)) {
	y := A{"foobar"}
	f(y)
}

func main() {
	simple(A.APrint)
	y := map[string]bool{"hello": true}
	for i := range y {
		fmt.Print(i)
	}
	x := A{"foobar"}
	x.APrint()
}
