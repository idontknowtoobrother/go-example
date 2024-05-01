package main

import (
	"fmt"

	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
	"github.com/idontknowtoobrother/go-example/datastructure"
	"github.com/idontknowtoobrother/go-example/grade"
	"github.com/idontknowtoobrother/go-example/hexafu"
	"github.com/idontknowtoobrother/go-example/leetcode"
	"github.com/idontknowtoobrother/go-example/loop"
	"github.com/idontknowtoobrother/go-example/pracconst"
)

func main() {

	fullname := "Jakkrit"
	var lastname string = "Chaopron"

	id := uuid.New()

	fmt.Printf("FULLNAME=%s\nLASTNAME=%s\nMACADDRESS=%v\nUUID=%v", fullname, lastname, faker.MacAddress(), id)

	fmt.Printf("\nGetRandomInt=%d", hexafu.GetRandomInt(100))

	fullname = "Chaopron"
	fmt.Printf("\nNEW_FULLNAME=%s", fullname)

	pracconst.Example()

	gradeInfo := fmt.Sprintf("GRADE=%s", grade.GetGrade(50))
	fmt.Println(gradeInfo)

	loop.ForLoop()
	loop.DoWhileLoop()
	loop.WhileLoop()

	fmt.Println()
	fmt.Println(datastructure.GetStructure("array"))
	fmt.Println(datastructure.GetStructure("slice"))
	datastructure.ConvertArrayToSlice()

	gradeMap := datastructure.GetIntMap()
	myScore := 20
	myGrade := getGradeWithGradeMap(gradeMap, myScore)
	fmt.Println("\nGetIntMap=", gradeMap)
	fmt.Println("My score=", myScore)
	fmt.Println("My Grade=", myGrade)

	datastructure.RemoveItemInMapExample()

	fmt.Println("TwoSum: nums={1,4,5,2} target=6 / result=", leetcode.TwoSum([]int{1, 4, 5, 2}, 6))
}

func getGradeWithGradeMap(gradeMap map[int]string, score int) string {
	for moreThan, gradeChar := range gradeMap {
		if score >= moreThan {
			return gradeChar
		}
	}
	return "F"
}
