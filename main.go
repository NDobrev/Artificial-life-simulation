package main

import (
	"fmt"
)

type A struct {
}

func (a *A) g() {
	fmt.Println("yae")
}

func main() {
	var a A
	a.g()
}
