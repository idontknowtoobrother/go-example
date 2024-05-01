package datastructure

import "fmt"

func GetIntMap() map[int]string {

	sampleGradeMap := make(map[int]string)
	sampleGradeMap[80] = "A"
	sampleGradeMap[70] = "B"
	sampleGradeMap[60] = "C"
	sampleGradeMap[50] = "D"

	return sampleGradeMap
}

func RemoveItemInMapExample() {

	userUUIDBankMap := make(map[string]int)

	userUUIDBankMap["John"] = 10000
	userUUIDBankMap["Rich"] = 550000
	userUUIDBankMap["SuperRich"] = 99999999
	userUUIDBankMap["Billionare"] = 9999999999

	fmt.Println("UUID-USER-BANK-MAP (Initialized)=", userUUIDBankMap)
	fmt.Println("Removing ['SuperRich'] ...")
	delete(userUUIDBankMap, "SuperRich")
	fmt.Println("Removed ['SuperRich'] :D")
	fmt.Println("UUID-USER-BANK-MAP (After removed)=", userUUIDBankMap)
}

func IsKeyExistInMap(uid string) bool {

	// checking is valid key in map ?
	userNameMap := make(map[string]string)
	userNameMap["U001"] = "John"
	userNameMap["U002"] = "Jane"

	_, ok := userNameMap[uid]
	return ok
}

func GetUserMap() {

}
