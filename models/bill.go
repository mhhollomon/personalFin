package models

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

const billListFileName = "billList.json"

var billsChanged = false

type Bill struct {
	AccountID  uuid.UUID `json:"acct_id"`
	OrigAmount float32   `json:"orig_amt"`
	Amount     float32   `json:"amount"`
	DueDate    time.Time `json:"date"`
	ID         uuid.UUID `json:"uuid"`
}

var billList = make([]*Bill, 0)

func (b *Bill) Pay(amount float32, _ time.Time) {

	b.Amount -= amount
	GetAccountById(b.AccountID).updateBalance(-amount)

	billsChanged = true

}

func NewBill(acct uuid.UUID, amount float32, due time.Time) *Bill {
	b := &Bill{
		AccountID:  acct,
		OrigAmount: amount,
		Amount:     amount,
		DueDate:    due,
		ID:         uuid.New(),
	}

	billList = append(billList, b)

	GetAccountById(acct).updateBalance(amount)

	billsChanged = true

	return b
}

func FindBillsForAcct(acct uuid.UUID) []*Bill {
	retval := make([]*Bill, 0, 10)
	for _, b := range billList {
		if b.AccountID == acct && b.Amount > 0 {
			retval = append(retval, b)
		}
	}
	return retval
}

func FindEarliestBillForAcct(acct uuid.UUID) *Bill {
	var retval *Bill

	for _, b := range billList {
		if b.AccountID == acct && b.Amount > 0 {
			if retval == nil || retval.DueDate.After(b.DueDate) {
				retval = b
			}
		}
	}

	return retval
}

func LoadBillList() error {

	file, ferr := os.Open(billListFileName)
	if ferr != nil {
		return ferr
	}
	defer file.Close()

	err := json.NewDecoder(file).Decode(&billList)

	if err != nil {
		log.Println("saw an error in decoding ", err)
	}
	return err
}

func SaveBillList() error {

	if !billsChanged {
		return nil
	}

	log.Println("Writing bill file because of changes")

	file, ferr := os.Create(billListFileName)
	if ferr != nil {
		return ferr
	}

	defer file.Close()

	err := json.NewEncoder(file).Encode(billList)
	billsChanged = err != nil
	return err

}
