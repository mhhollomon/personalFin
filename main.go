package main

import (
	"log"
	"pf/account"
	"pf/layouts"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func addDialog(w fyne.Window, l *widget.List) {

	t := widget.NewRadioGroup(account.AccountTypes, func(string) {})
	s := "Account Name"
	sb := binding.BindString(&s)
	e := widget.NewEntryWithData(sb)

	box := container.NewVBox(t, e)

	dialog.ShowCustomConfirm("Add Account", "Add", "Cancel", box, func(a bool) {
		if a {
			_, err := account.AddAccount(s, t.Selected)
			if err == nil {
				l.Refresh()
			} else {
				dialog.ShowError(err, w)
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
			acct, _ := account.GetAccountByIndex(i)
			l := o.(*widget.Label)
			l.SetText(acct.Name)
		})

	acctScroll := container.NewVScroll(accountList)

	addAccountButton := widget.NewButton("+ Add Account", func() { addDialog(mainWindow, accountList) })

	mainWindow.SetContent(container.New(layouts.NewVFlex(300, 200), addAccountButton, acctScroll))
	mainWindow.ShowAndRun()

	account.SaveAccountList()

}
