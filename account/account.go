package account

import (
	"encoding/json"
	"log"
	"os"
)

const accountListFileName = "accountList.json"

type Account struct {
	Name string `json:"name"`
}

var accountList = make(map[string]Account)

func LoadAccountList() error {
	file, ferr := os.Open(accountListFileName)
	if ferr != nil {
		return ferr
	}
	defer file.Close()

	err := json.NewDecoder(file).Decode(&accountList)
	return err
}

func SaveAccountList() error {

	file, ferr := os.Create(accountListFileName)
	if ferr != nil {
		return ferr
	}

	defer file.Close()

	err := json.NewEncoder(file).Encode(accountList)
	return err

}

func ListAccounts() []string {
	log.Println("account.ListAccounts called")
	retval := make([]string, 0)

	for _, v := range accountList {
		retval = append(retval, v.Name)
	}

	return retval

}

func GetAccount(name string) (Account, bool) {
	acct, exists := accountList[name]
	return acct, exists
}

func AddAccount(name string) (Account, bool) {
	newAcct := Account{name}

	if _, exists := GetAccount(name); exists {
		return newAcct, false
	}

	accountList[name] = newAcct
	return newAcct, true
}
