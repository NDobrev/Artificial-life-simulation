package main

import (
	"Artificial-life-simulation/core"
	"fmt"
)

func main() {
	a := core.NewField(100)
	pla := core.NewPhytoPlankton()
	var point core.FieldPoint
	point.SetPoint(1, 1)
	a.AddObject(point, pla)
	fmt.Println("--------")
	a.Print()
	fmt.Println("--------")
	for i := 0; i < 25; i++ {
		a.OnTick()
		//a.Print()
	}
	a.Print()
}
