package core

type ObjType int

const (
	None ObjType = iota

	//Replaceble object
	FirstReplaceble
	Empty
	LightSpace
	LastReplaceble

	Movable

	Doable
	FirstPlankton
	ZooPlanktonT
	PhytoPlanktonT
	LastPlankton
	NumberOfTypes
)
