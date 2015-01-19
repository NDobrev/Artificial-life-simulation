package main

import (
	"Artificial-life-simulation/core"
	"Artificial-life-simulation/gui"
	//"fmt"
)

func main() {
	a := core.NewField(50)
	pla := core.NewPhytoPlankton()
	var point core.FieldPoint
	point.SetPoint(1, 1)
	a.AddObject(point, pla)
	//	a.Print()
	for i := 0; i < 150; i++ {
		if i == 10 {
			//	fmt.Println("befor")
			//a.Print()
			//a.RemoveFrom(point)
			a.AddObject(point, core.NewZooPlankton())
			//fmt.Println("after")
			//a.Print()
		}
		a.OnTick()
		//a.Print()
	}
	gui.WinGuiMain()
	//a.Print()
}
