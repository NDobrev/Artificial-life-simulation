package core

type ObjLocator struct {
	location FieldPoint
	obj      FieldObject
}

func NewObjLocator(p FieldPoint, obj FieldObject) *ObjLocator {
	o := new(ObjLocator)
	o.location = FieldPoint{p.x, p.y}
	o.obj = obj
	return o
}
