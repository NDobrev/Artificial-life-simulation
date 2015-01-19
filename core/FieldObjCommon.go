package core

func IsPlankton(fo FieldObject) bool {
	if fo.GetType() > FirstPlankton && fo.GetType() < LastPlankton {
		return true
	}
	return false
}

func IsReplaceble(fo FieldObject) bool {
	if fo.GetType() > FirstReplaceble && fo.GetType() < LastReplaceble {
		return true
	}
	return false
}

func GenFuncForObjType(ot ObjType) func(FieldObject) bool {
	return func(fo FieldObject) bool {
		if fo.GetType() == ot {
			return true
		}
		return false
	}
}
