package core

type ObjType int

const (
	None ObjType = iota
	FirstStatic
	// Members of Static family
	RockT
	LimiterT
	LastStatic

	//Replaceble object
	FirstReplaceble
	// Members of Replaceble family
	Empty
	LitSpaceT
	LastReplaceble

	FirstMovable
	//Members of Movable family
	FirstDoable
	// Members of Doable family
	FirstPlankton
	// Family Movable
	// Sub family Doable
	// Members of Plankton family
	ZooPlanktonT
	PhytoPlanktonT
	LightSensitivePlanktonT
	PredatoryPlanktonT

	LastPlankton

	SwarmUnitT

	LastMovable

	NumberOfTypes
)
