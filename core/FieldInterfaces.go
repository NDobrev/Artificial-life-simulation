package core

type FieldObject interface {
	GetCopy() *FieldObject
	GetType() *ObjType
}
