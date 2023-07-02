package main

import (
	"log"
	"pf/account"
	"pf/dialogs"
	"pf/layouts"
	"pf/widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	log.SetFlags(0)
	account.LoadAccountList()

	application := app.New()
	mainWindow := application.NewWindow("Personal Finance")

	accountList := widget.NewList(
		func() int { return account.CountAccounts() },
		func() fyne.CanvasObject { return widgets.NewBlankAccountSummary() },
		func(i widget.ListItemID, o fyne.CanvasObject) {
			acct, _ := account.GetAccountByIndex(i)
			l := o.(*widgets.AccountSummary)
			l.SetName(acct.Name)
			l.SetAmount(acct.Balance)
		})

	acctScroll := container.NewVScroll(accountList)

	addAccountButton := widget.NewButton("+ Add Account", func() { dialogs.AddAccountDialog(mainWindow, accountList) })

	mainWindow.SetContent(container.New(layouts.NewVFlex(500, 300), addAccountButton, acctScroll))
	mainWindow.ShowAndRun()

	account.SaveAccountList()

}
