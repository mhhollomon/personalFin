package commands

import (
	"errors"
	"fmt"
	"pf/account"
)

var listAccountsSpec = CommandSpec{Fn: listAccountsCmd}
var listAnAccountSpec = CommandSpec{Fn: listAnAccountCmd, Args: []string{"account"}}
var addAccountSpec = CommandSpec{Fn: addAccountCmnd, Args: []string{"account"}}

func listAccountsCmd([]string) error {
	fmt.Println(account.ListAccounts())
	return nil
}

func listAnAccountCmd(args []string) error {

	// commands.Execute makes sure our arg count is correct
	accountName := args[0]

	// but we need to make sure the arg makes sense
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

func addAccountCmnd(args []string) error {
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
