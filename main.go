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
	for i := 0; i < 1500; i++ {
		if i == 10 {
			a.RemoveFrom(point)
			a.AddObject(point, core.NewZooPlankton())
		}
		a.OnTick()

	}
	a.Print()
	gui.WinGuiMain()

}
