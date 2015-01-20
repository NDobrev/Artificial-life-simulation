package main

import (
	"Artificial-life-simulation/core"
	"Artificial-life-simulation/gui/server"
	"encoding/json"
	"fmt"
	"net/http"
)

type FieldColorRepresentation struct {
	Colors [][]int
}

func StartGuiServer(f core.FieldBase) {
	var s server.Server
	listen := make(chan interface{})
	request := make(chan [][]int)

	matrixRequest := func(http.ResponseWriter, *http.Request) {
		listen <- 1
		matr := <-request
		result, _ := json.Marshal(matr)
		fmt.Print(result)
	}
	s.RegisterFunc("/matrix", matrixRequest)
	go s.Run()
}

func main() {
	a := core.NewLitField(50, 10)
	StartGuiServer(a)

	pla := core.NewPhytoPlankton()
	var point core.FieldPoint
	point.SetPoint(1, 1)
	a.AddObject(point, pla)

	for i := 0; i < 1500; i++ {
		if i == 10 {
			a.RemoveFrom(point)

		}
		a.OnTick()

	}
	a.Print()
}
