package core

import (
	//"fmt"
	"math"
	"math/rand"
)

type SwarmUnit struct {
}

func (su SwarmUnit) GetType() ObjType {
	return SwarmUnitT
}

func NewSwarmUnit() *SwarmUnit {
	return new(SwarmUnit)
}

func (su SwarmUnit) Do(f FieldBase, myLocation FieldPoint) {
	swarm := f.GetAllWithTypeInSquare(AllType, FieldPoint{myLocation.x, myLocation.y}, 7)
	var hor, vert int
	for idx := range swarm {
		if swarm[idx].obj.GetType() == SwarmUnitT {
			hor += int(math.Pow(-1, float64((myLocation.x-swarm[idx].location.x)%2)))
			vert += int(math.Pow(-1, float64((myLocation.y-swarm[idx].location.y)%2)))
		}
	}

	//fmt.Println("hor: %i vert: %i", hor, vert)
	if math.Abs(float64(hor)) > math.Abs(float64(vert)) {
		for idx := range swarm {
			//	fmt.Println((myLocation.x - swarm[idx].location.x) * hor)
			//	fmt.Println(myLocation, swarm[idx])
			if IsReplaceble(swarm[idx].obj) && (myLocation.x-swarm[idx].location.x)*hor > 0 {
				s := ((myLocation.x - swarm[idx].location.x) * hor) / ((myLocation.x - swarm[idx].location.x) * hor)
				f.MoveFromTo(myLocation, FieldPoint{myLocation.x + s, myLocation.y})
				//	fmt.Println("MYRDAI")
				return
			}
		}
	}
	if math.Abs(float64(hor)) < math.Abs(float64(vert)) {
		for idx := range swarm {
			if IsReplaceble(swarm[idx].obj) && (myLocation.y-swarm[idx].location.y)*vert > 0 {
				s := ((myLocation.y - swarm[idx].location.y) * vert) / ((myLocation.y - swarm[idx].location.y) * vert)
				f.MoveFromTo(myLocation, FieldPoint{myLocation.x, myLocation.y + s})
				return
			}
		}
	}
	free := f.GetAllWithTypeInSquare(IsReplaceble, FieldPoint{myLocation.x, myLocation.y}, 3)
	if len(free) > 0 {
		f.MoveFromTo(myLocation, free[rand.Int()%len(free)].location)
	}
}
