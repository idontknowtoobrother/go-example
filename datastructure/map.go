package datastructure

import (
	"fmt"
	"math/rand"

	"github.com/bxcodec/faker/v4"
)

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

func IsKeyExistInMap(uid string) (string, bool) {

	// checking is valid key in map ?
	userNameMap := make(map[string]string)
	userNameMap["U001"] = "John"
	userNameMap["U002"] = "Jane"

	value, ok := userNameMap[uid]
	return value, ok
}

func GetLoanAccountsMap(total int) map[string]LoanAccount {

	lasm := make(map[string]LoanAccount)

	for i := 0; i < total; i++ {
		totalLoad := float64(rand.Intn(100000))
		mTransction := Transaction{
			TSCID:         faker.UUIDDigit(),
			DateTime:      faker.Timestamp(),
			Total:         float64(rand.Intn(200)),
			ToServiceName: faker.DomainName(),
		}

		mockLoanAccount := LoanAccount{
			ACCID:      faker.UUIDDigit(),
			UID:        faker.UUIDDigit(),
			CardNumber: faker.CCNumber(),
			CCType:     faker.CCType(),
			TotalLoan:  totalLoad,
			Balance:    totalLoad - mTransction.Total,
			Transactions: []Transaction{
				mTransction,
			},
		}

		lasm[mockLoanAccount.ACCID] = mockLoanAccount
	}

	return lasm
}
