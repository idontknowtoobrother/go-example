package datastructure

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
	TotalLoan    float64       // total of loan plan
	Balance      float64       // current balance
	Transactions []Transaction // every transactions of user
}

func GetLoanAccountStruct() LoanAccount {
	loanAccount1 := LoanAccount{
		ACCID:        "ACCID105",
		UID:          "UID102",
		TotalLoan:    100000,
		Balance:      100000,
		Transactions: []Transaction{},
	}

	return loanAccount1
}
