package main

import (
	"fmt"

	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
	"github.com/idontknowtoobrother/go-example/datastructure"
	"github.com/idontknowtoobrother/go-example/grade"
	"github.com/idontknowtoobrother/go-example/hexafu"
	"github.com/idontknowtoobrother/go-example/interfaces"
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

	findingKey := "U001"
	value, found := datastructure.IsKeyExistInMap(findingKey)
	if found {
		fmt.Printf("Found key=%s exist value=%s\n", findingKey, value)
	} else {
		fmt.Printf("Not found key=%s exist value=NULL\n", findingKey)
	}
	fmt.Println("TwoSum: nums={1,4,5,2} target=6 / result=", leetcode.TwoSum([]int{1, 4, 5, 2}, 6))

	exLoadAccounts := datastructure.GenerateLoanAccounts(10)
	for _, account := range exLoadAccounts {
		fmt.Println("ACCID: ", account.ACCID)
		fmt.Println("UID: ", account.UID)
		fmt.Println("CardNumber: ", account.CardNumber)
		fmt.Println("CCType: ", account.CCType)
		fmt.Println("TotalLoan: ", account.TotalLoan)
		fmt.Println("Balance: ", account.Balance)
		fmt.Println("Transactions: ", account.Transactions)
		fmt.Println("---------------------------------")
	}

	account, ok := datastructure.GetStructure("struct").(datastructure.LoanAccount)
	if ok {
		fmt.Println("Loan account struct=", account)
		fmt.Println("Loan account.ACCID=", account.ACCID)
		fmt.Println("Loan account.UID=", account.UID)
		fmt.Println("Loan account.TotalLoan=", account.TotalLoan)
		fmt.Println("Loan account.Balance=", account.Balance)
		fmt.Println("Loan account.Transactions=", account.Transactions)

		bill := datastructure.PayBill{
			ToServiceDomain: "awayfromus.dev",
			Total:           5680.23,
		}
		fmt.Println("Paying bill ", bill)
		account.Pay(bill)
		fmt.Println("Loan account.Balance=", account.Balance)
		fmt.Println("Loan account.Transactions=", account.Transactions)

	}

	// INTERFACE SECTION
	dog := interfaces.Dog{Name: "Buddy"}
	person := interfaces.Person{Name: "Alice"}

	interfaces.MakeSound(dog)
	interfaces.MakeSound(person)

	circle := interfaces.Circle{
		Radius: 30.50,
	}
	rectangle := interfaces.Rectangle{
		Width:  30,
		Height: 50,
	}
	triangle := interfaces.Triangle{
		Base:   80,
		Height: 20.5,
	}
	fmt.Println("Circle has area=", interfaces.CalculateArea(circle))
	fmt.Println("Rectangle has area=", interfaces.CalculateArea(rectangle))
	fmt.Println("Triangle has radius=", interfaces.CalculateArea(triangle))
	// END INTERFACE SECTION

}

func getGradeWithGradeMap(gradeMap map[int]string, score int) string {
	for moreThan, gradeChar := range gradeMap {
		if score >= moreThan {
			return gradeChar
		}
	}
	return "F"
}
