package main

import (
	"log"
	"pf/dialogs"
	"pf/globals"
	"pf/layouts"
	"pf/models"
	"pf/widgets"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	log.SetFlags(0)

	models.LoadAccountList()
	defer models.SaveAccountList()

	models.LoadBillList()
	defer models.SaveBillList()

	application := app.New()
	mainWindow := application.NewWindow("Personal Finance")

	accountList := widget.NewList(
		func() int { return models.CountAccounts() },
		func() fyne.CanvasObject {
			log.Println("Returning Template Account")
			return widgets.NewAccountSummary(mainWindow)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			acct, _ := models.GetAccountByIndex(i)
			log.Printf("Setting data for index %d %+v", i, *acct)
			l := o.(*widgets.AccountSummary)
			l.SetName(acct.Name)
			l.SetAmount(acct.Balance)
		})

	acctScroll := container.NewVScroll(accountList)

	addAccountButton := widget.NewButton("+ Add Account", func() { dialogs.AddAccountDialog(mainWindow) })

	mainWindow.SetContent(container.New(layouts.NewVFlex(500, 300), addAccountButton, acctScroll))

	globals.StartUpdater(accountList)
	defer globals.RequestClose()

	mainWindow.ShowAndRun()

}
