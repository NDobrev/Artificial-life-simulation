package core

import (
	"fmt"
	"math/rand"
)

type Limiter struct {
	chunkSize  int
	polulation int
}

func (l Limiter) GetType() ObjType {
	return LimiterT
}

func NewLimiter(cs int, population int) *Limiter {
	result := new(Limiter)
	result.chunkSize = cs
	result.polulation = population
	return result
}

func (l Limiter) Restruct(f FieldBase, myLocation FieldPoint) {
	plankton := f.GetAllWithTypeInSquare(IsPlankton, FieldPoint{myLocation.x, myLocation.y}, l.chunkSize)
	cnt := len(plankton) - l.polulation
	fmt.Println("Population plankton is ", cnt)
	for cnt > 0 {
		r := rand.Int() % len(plankton)
		fo, _ := f.LookAt(plankton[r].location)
		if IsPlankton(fo) {
			if rand.Int()%fo.(CellBase).GetDeadTime()*5 < fo.(CellBase).GetAge()*fo.(CellBase).GetAge() {
				f.RemoveFrom(plankton[rand.Int()%len(plankton)].location)
			}
		}
		cnt--
	}
}
