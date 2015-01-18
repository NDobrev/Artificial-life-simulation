package core

type ObjType int

const (
	Empty ObjType = iota

	Movable
	Doable
	FirstPlankton
	ZooPlanktonT
	PhytoPlanktonT
	LastPlankton
	NumberOfTypes
)
