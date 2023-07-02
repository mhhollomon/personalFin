package models

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/google/uuid"
)

const accountListFileName = "accountList.json"

var AccountTypes = []string{"Payable", "Cash"}

type Account struct {
	Name    string    `json:"name"`
	ID      uuid.UUID `json:"id"`
	Type    string    `json:"type"`
	Balance float32   `json:"balance"`
}

var accountsChanged = false

var accountList = make([]*Account, 0)

func LoadAccountList() error {
	file, ferr := os.Open(accountListFileName)
	if ferr != nil {
		return ferr
	}
	defer file.Close()

	err := json.NewDecoder(file).Decode(&accountList)

	if err != nil {
		log.Println("saw an error in decoding ", err)
	}
	return err
}

func SaveAccountList() error {

	if !accountsChanged {
		return nil
	}

	log.Println("Writing account file because of changes")

	file, ferr := os.Create(accountListFileName)
	if ferr != nil {
		return ferr
	}

	defer file.Close()

	err := json.NewEncoder(file).Encode(accountList)
	accountsChanged = err != nil
	return err

}

func CountAccounts() int {
	return len(accountList)
}

func GetAccountByName(name string) *Account {

	for _, v := range accountList {
		if v.Name == name {
			return v
		}
	}

	return nil
}

func GetAccountByIndex(i int) (*Account, bool) {
	return accountList[i], true
}

func GetAccountById(i uuid.UUID) *Account {
	for _, a := range accountList {
		if a.ID == i {
			return a
		}
	}

	return nil

}

func (a *Account) updateBalance(delta float32) {
	a.Balance += delta
	accountsChanged = true
}

func AddAccount(name string, acctType string, startingBalance float32) (Account, error) {

	if acct := GetAccountByName(name); acct != nil {
		log.Println("proposed account already exists: ", name)
		return Account{}, errors.New("proposed account already exists")
	}

	found := false
	for _, v := range AccountTypes {
		if v == acctType {
			found = true
			break
		}
	}

	if !found {
		return Account{}, errors.New("invalid account type")
	}

	newAcct := &Account{Name: name, ID: uuid.New(), Type: acctType, Balance: startingBalance}

	accountList = append(accountList, newAcct)
	accountsChanged = true

	log.Printf("added account: %+v\n", newAcct)

	return *newAcct, nil
}
