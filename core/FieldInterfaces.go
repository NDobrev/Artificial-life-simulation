package core

type FieldObject interface {
	GetCopy() FieldObject
	GetType() ObjType
}

type DoableObject interface {
	Do(f FieldBase, myLocation FieldPoint)
}

type FieldBase interface {
	AddObject(p FieldPoint, obj FieldObject) bool
	LookAt(p FieldPoint) (FieldObject, error)
	GetAllWithType(t ObjType) []*ObjLocator
	checkCorectPoint(p FieldPoint) bool
	GetAllWithTypeInSquare(comp func(FieldObject) bool, center FieldPoint, size int) []*ObjLocator
	MoveFromTo(from FieldPoint, to FieldPoint) bool
	RemoveFrom(from FieldPoint)
	ClearField()
	OnTick()
	Print()
}
