package datastructure

func GetStructure(structureType string) interface{} {

	if structureType == "array" {
		return GetArray()
	} else if structureType == "slice" {
		return GetSlice()
	} else if structureType == "map" {
		return GetIntMap()
	} else if structureType == "struct" {
		return GenerateLoanAccount()
	}

	return nil
}
