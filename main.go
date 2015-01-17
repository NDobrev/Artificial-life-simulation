package main

import (
	"fmt"
)

type A struct {
}

func (a *A) g() {
	fmt.Println("yae")
}

type B struct {
	A
}

func main() {
	var b B
	b.g()
	b.A.g()
}
