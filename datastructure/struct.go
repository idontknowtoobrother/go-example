package datastructure

import "github.com/bxcodec/faker/v4"

// @ Syntax
//
// type <NAME_OF_STRUCT> struct {
// 	DATA TYPE
// 	DATA TYPE
// 	...
// }
//

type Transaction struct {
	TSCID         string  // transaction id
	DateTime      string  // date time of trasaction
	Total         float64 // total money in this transaction
	ToServiceName string  // paid to who ?
}

type LoanAccount struct {
	ACCID        string        // account id
	UID          string        // user id
	CardNumber   string        // card number to use
	CCType       string        // ex. American Express
	TotalLoan    float64       // total of loan plan
	Balance      float64       // current balance
	Transactions []Transaction // every transactions of user
}

func GetLoanAccountStruct() LoanAccount {
	loanAccount1 := LoanAccount{
		ACCID:        faker.UUIDDigit(),
		UID:          faker.UUIDDigit(),
		TotalLoan:    100000.99,
		Balance:      100000.99,
		CardNumber:   faker.CCNumber(),
		CCType:       faker.CCType(),
		Transactions: []Transaction{},
	}

	return loanAccount1
}
