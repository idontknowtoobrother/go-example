package datastructure

func GetStructure(structureType string) interface{} {

	if structureType == "array" {
		return GetArray()
	} else if structureType == "slice" {
		return GetSlice()
	}

	return nil
}
