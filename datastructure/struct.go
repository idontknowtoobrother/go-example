package datastructure

import (
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v4"
)

// @ Syntax
//
// type <NAME_OF_STRUCT> struct {
// 	DATA TYPE
// 	DATA TYPE
// 	...
// }
//

type Transaction struct {
	TSCID           string  // transaction id
	DateTime        string  // date time of trasaction
	Total           float64 // total money in this transaction
	ToServiceDomain string  // paid to who ?
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

type PayBill struct {
	ToServiceDomain string  //	to service domain
	Total           float64 // total
}

// reciever arguments (OOP Style)
func (account *LoanAccount) GetBalance() float64 { // get balance
	return account.Balance
}

func (account *LoanAccount) IsMoneyEnough(total float64) bool { // check is money enough
	return account.Balance >= total
}

func (account *LoanAccount) Pay(bill PayBill) bool { // pay to service domain
	moneyEnough := account.IsMoneyEnough(bill.Total)
	if !moneyEnough {
		return false
	}

	account.Balance -= bill.Total

	newTransaction := Transaction{
		TSCID:           faker.UUIDDigit(),
		DateTime:        time.Now().String(),
		Total:           bill.Total,
		ToServiceDomain: bill.ToServiceDomain,
	}
	account.Transactions = append(account.Transactions, newTransaction)

	return true
}

// public mock data
func GenerateLoanAccount() LoanAccount {
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

func GenerateLoanAccounts(total int) map[string]LoanAccount {

	lasm := make(map[string]LoanAccount)

	for i := 0; i < total; i++ {
		totalLoad := float64(rand.Intn(100000))
		mTransction := Transaction{
			TSCID:           faker.UUIDDigit(),
			DateTime:        faker.Timestamp(),
			Total:           float64(rand.Intn(200)),
			ToServiceDomain: faker.DomainName(),
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
