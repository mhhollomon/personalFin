package commands

import (
	"errors"
	"fmt"
	"pf/account"
)

func listAccountsCmd([]string) error {
	fmt.Println(account.ListAccounts())
	return nil
}

func listAnAccountCmd(args []string) error {
	if len(args) < 1 {
		return errors.New("no account given")
	}
	accountName := args[0]

	if accountName == "" {
		return errors.New("no account given")
	}

	acct, ok := account.GetAccount(accountName)

	if !ok {
		return errors.New("account does not exist")
	}

	fmt.Println(acct)

	return nil
}

func AddAccountCmnd(args []string) error {
	if len(args) < 1 {
		return errors.New("no account given")
	}
	accountName := args[0]

	if accountName == "" {
		return errors.New("no account given")
	}

	acct, ok := account.AddAccount(accountName)

	if !ok {
		return errors.New("account already exists")
	}

	fmt.Println(acct)

	return nil
}
