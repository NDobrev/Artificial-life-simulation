package main

import (
	"Artificial-life-simulation/core"
	"Artificial-life-simulation/gui/server"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"runtime"
	"strconv"
	//	"time"
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

type ObjectRequestData struct {
	x, y int64
	t    core.ObjType
}

type sn struct {
	update chan interface{}
	rdy    chan interface{}
	reset  chan interface{}
	setReq chan ObjectRequestData
}

func snInit() *sn {
	result := new(sn)
	result.update = make(chan interface{})
	result.rdy = make(chan interface{})
	result.reset = make(chan interface{})
	result.setReq = make(chan ObjectRequestData)
	return result
}

var req *sn

func SendRequest(r *http.Request, t core.ObjType) {
	var ord ObjectRequestData
	r.ParseForm()
	ord.x, _ = strconv.ParseInt(r.PostFormValue("x"), 0, 64)
	ord.y, _ = strconv.ParseInt(r.PostFormValue("y"), 0, 64)
	ord.t = t
	fmt.Println(ord)
	req.setReq <- ord
}

func StartGuiServer() {
	var s server.Server
	req = snInit()
	colors = MakeMatrix(Size, Size)
	matrixRequest := func(w http.ResponseWriter, r *http.Request) {
		// za da moje6 da pra6ta6 kum browser-a response, 4e ima nqkvi security gluposti koit pre4at
		w.Header().Set("Access-Control-Allow-Origin", "*")
		req.update <- struct{}{}
		_ = <-req.rdy
		result, _ := json.Marshal(colors)
		w.Write(result)

	}

	ResetSimulationRequest := func(w http.ResponseWriter, r *http.Request) {
		req.reset <- struct{}{}
	}

	CommonPhytoPlankton := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("New CommonPhytoPlankton")
		SendRequest(r, core.PhytoPlanktonT)
	}

	LightSensitivePlankton := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("New LightSensitivePlankton")
		SendRequest(r, core.LightSensitivePlanktonT)
	}

	ZooPlankton := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("New ZooPlankton")
		SendRequest(r, core.ZooPlanktonT)
	}

	PredatoryPlankton := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("New PredatoryPlankton")
		SendRequest(r, core.PredatoryPlanktonT)
	}

	Limiter := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("New Limiter")
		SendRequest(r, core.LimiterT)
	}

	Swarm := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("New Limiter")
		SendRequest(r, core.SwarmUnitT)
	}

	s.RegisterFunc("/Swarm", Swarm)
	s.RegisterFunc("/Limiter", Limiter)
	s.RegisterFunc("/CommonPhytoPlankton", CommonPhytoPlankton)
	s.RegisterFunc("/LightSensitivePlankton", LightSensitivePlankton)
	s.RegisterFunc("/ZooPlankton", ZooPlankton)
	s.RegisterFunc("/PredatoryPlankton", PredatoryPlankton)
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
	r := float64((70-x)*(70-x) + (70-y)*(70-y))
	r = math.Sqrt(r)
	//fmt.Println(40 - r)

	return int(r)
}

func Execute(f core.FieldBase, ord ObjectRequestData) {
	fp := core.NewFieldPoint(int(ord.x), int(ord.y))
	var fo core.FieldObject
	switch ord.t {
	case core.ZooPlanktonT:
		fo = core.NewZooPlankton()
	case core.PhytoPlanktonT:
		fo = core.NewPhytoPlankton()
	case core.PredatoryPlanktonT:
		fo = core.NewPredatoryPlankton()
	case core.LightSensitivePlanktonT:
		fo = core.NewLightSensitivePlankton()
	case core.LitSpaceT:
		fo = core.NewLitPlace(SumMod20(int(ord.x), int(ord.y)))
	case core.RockT:
		fo = core.NewRock()
	case core.Empty:
		fo = core.NewEmptyPlace()
	case core.LimiterT:
		fo = core.NewLimiter(140, 3500)
	case core.SwarmUnitT:
		fo = core.NewSwarmUnit()
	}
	f.RemoveFrom(*fp)
	f.AddObject(*fp, fo)
}

var Size int

func main() {
	Size = 150
	runtime.GOMAXPROCS(4)
	fmt.Println("start")
	a := core.NewLitFieldWithFillWithFunction(Size, SumMod20)
	StartGuiServer()

	fmt.Println("startGUIServer done")
	fmt.Println("Init field done")
	//<-time.after(time.Nanosecond * 1)
	for {
		b := true
		a.ClearField()
		for b {
			select {
			case _ = <-req.update:
				fmt.Println("UpdateReq")
				a.ColorRepresentation(colors.Colors)
				req.rdy <- struct{}{}
			case _ = <-req.reset:
				b = false
				a.ClearField()
			case obj := <-req.setReq:
				Execute(a, obj)
				//case <-time.After(time.Nanosecond * 1):
				//	fmt.Println("NextInteration")
			}
			//fmt.Println(i)
			a.OnTick()

		}
	}
}
