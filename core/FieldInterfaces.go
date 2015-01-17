package core

type FieldObject interface {
	GetCopy() *FieldObject
	GetType() *ObjType
}

type DoableObject interface {
	Do(f *Field)
}
