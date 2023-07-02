package account

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

const accountListFileName = "accountList.json"

var AccountTypes = []string{"Payable", "Cash"}

type Account struct {
	Name    string  `json:"name"`
	ID      int     `json:"id"`
	Type    string  `json:"type"`
	Balance float32 `json:"balance"`
}

type AccountFile struct {
	LastID      int       `json:"last_id"`
	AccountList []Account `json:"accounts"`
}

var accounts AccountFile

var accountsChanged = false

var nameList = make([]string, 0)

func LoadAccountList() error {
	file, ferr := os.Open(accountListFileName)
	if ferr != nil {
		return ferr
	}
	defer file.Close()

	err := json.NewDecoder(file).Decode(&accounts)

	if err == nil {
		for _, v := range accounts.AccountList {
			nameList = append(nameList, v.Name)
		}

	} else {
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

	err := json.NewEncoder(file).Encode(accounts)
	accountsChanged = err != nil
	return err

}

func ListAccounts() []string {
	log.Println("account.ListAccounts called")
	retval := make([]string, 0)

	retval = append(retval, nameList...)

	return retval

}

func CountAccounts() int {
	return len(nameList)
}

func GetAccount(name string) (Account, bool) {

	for _, v := range accounts.AccountList {
		if v.Name == name {
			return v, true
		}
	}

	return Account{}, false
}

func GetAccountByIndex(i int) (Account, bool) {

	account, exists := GetAccount(nameList[i])

	return account, exists
}

func AddAccount(name string, acctType string, startingBalance float32) (Account, error) {

	if _, exists := GetAccount(name); exists {
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

	accounts.LastID++
	newAcct := Account{Name: name, ID: accounts.LastID, Type: acctType, Balance: startingBalance}

	accounts.AccountList = append(accounts.AccountList, newAcct)
	accountsChanged = true
	nameList = append(nameList, name)

	log.Printf("added account: %+v\n", newAcct)

	return newAcct, nil
}
