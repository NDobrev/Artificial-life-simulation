package main

import (
	"Artificial-life-simulation/core"
	"Artificial-life-simulation/gui/server"
	"encoding/json"
	"fmt"
	"runtime"
	//"go/build"
	//"math"
	"net/http"
	//"sync"
)

type FieldColorRepresentation struct {
	Colors [][]uint
	//sync.Mutex
}

func MakeMatrix(width, height int) *FieldColorRepresentation {
	a := make([][]uint, height)
	for i := range a {
		a[i] = make([]uint, width)
		for j := range a[i] {
			a[i][j] = 0
		}
	}

	return &FieldColorRepresentation{Colors: a}
}

var colors *FieldColorRepresentation

type sn struct {
	update chan interface{}
	rdy    chan interface{}
	reset  chan interface{}
}

func snInit() *sn {
	result := new(sn)
	result.update = make(chan interface{})
	result.rdy = make(chan interface{})
	result.reset = make(chan interface{})
	return result
}

var req *sn

func StartGuiServer() {
	var s server.Server
	req = snInit()
	colors = MakeMatrix(140, 140)
	matrixRequest := func(w http.ResponseWriter, r *http.Request) {
		// za da moje6 da pra6ta6 kum browser-a response, 4e ima nqkvi security gluposti koit pre4at
		w.Header().Set("Access-Control-Allow-Origin", "*")
		req.update <- struct{}{}
		_ = <-req.rdy
		result, _ := json.Marshal(colors)
		w.Write(result)

	}

	ResetSimulationRequest := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Reset")
		req.reset <- struct{}{}
	}
	s.RegisterFunc("/reset", ResetSimulationRequest)
	s.RegisterFunc("/matrix", matrixRequest)
	go s.Run()
}

func BuildWall(f core.FieldBase) {
	var point core.FieldPoint
	for i := 10; i < 45; i++ {
		point.SetPoint(15, i)
		f.AddObject(point, core.NewRock())
	}
}

func SumMod20(x int, y int) int {
	return x + y
}

func main() {
	runtime.GOMAXPROCS(4)
	fmt.Println("start")
	a := core.NewLitFieldWithFillWithFunction(140, SumMod20)
	StartGuiServer()

	fmt.Println("startGUIServer done")
	fmt.Println("Init field done")

	for {
		pla := core.NewLightSensitivePlankton()
		var point core.FieldPoint
		point.SetPoint(1, 1)
		a.AddObject(point, pla)

		b := true
		for b {
			select {
			case _ = <-req.update:
				a.ColorRepresentation(colors.Colors)
				req.rdy <- struct{}{}
			case _ = <-req.reset:
				b = false
				a.ClearField()
				//default:
			}
			//fmt.Println(i)
			a.OnTick()

		}
	}
}
