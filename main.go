package main

import (
	"errors"
	"log"
	"pf/account"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func addDialog(w fyne.Window, l *widget.List) {
	s := "Account Name"
	sb := binding.BindString(&s)
	e := widget.NewEntryWithData(sb)

	dialog.ShowCustomConfirm("Add Account", "Add", "Cancel", e, func(a bool) {
		if a {
			_, okay := account.AddAccount(s)
			if okay {
				l.Refresh()
			} else {
				dialog.ShowError(errors.New("Could not add account"), w)
			}
		}
	}, w)
}

func main() {
	log.SetFlags(0)
	account.LoadAccountList()

	application := app.New()
	mainWindow := application.NewWindow("Hello World")

	accountList := widget.NewList(
		func() int { return account.CountAccounts() },
		func() fyne.CanvasObject { return widget.NewLabel("template") },
		func(i widget.ListItemID, o fyne.CanvasObject) {
			acct, _ := account.GetAccountById(i)
			o.(*widget.Label).SetText(acct.Name)
		})

	acctScroll := container.NewVScroll(accountList)
	acctScroll.SetMinSize(fyne.Size{Width: 300, Height: 200})

	addAccountButton := widget.NewButton("+ Add Account", func() { addDialog(mainWindow, accountList) })

	mainWindow.SetContent(container.NewVBox(addAccountButton, acctScroll))
	mainWindow.ShowAndRun()

	account.SaveAccountList()

}
