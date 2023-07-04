package main

import (
	"log"
	"os"
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

func setDataDir(dir string) {

	last_char := dir[len(dir)-1]

	if last_char != '/' && last_char != '\\' {
		dir = dir + "/"
	}

	globals.DataDir = dir
}

func main() {
	log.SetFlags(0)

	if len(os.Args) > 1 {
		setDataDir(os.Args[1])
	}

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
			acctSum := o.(*widgets.AccountSummary)
			acctSum.SetName(acct.Name)
			acctSum.SetAmount(acct.Balance)

			bill := models.FindEarliestBillForAcct(acct.ID)
			if bill != nil {
				log.Printf("A bill is due")
				acctSum.SetDueDate(bill.DueDate.Format("2006-01-02"))
			} else {
				log.Printf("A bill is Not due")
				acctSum.SetDueDate("  ")
			}
		})

	acctScroll := container.NewVScroll(accountList)

	addAccountButton := widget.NewButton("+ Add Account", func() { dialogs.AddAccountDialog(mainWindow) })

	mainWindow.SetContent(container.New(layouts.NewVFlex(500, 300), addAccountButton, acctScroll))

	globals.StartUpdater(accountList)
	defer globals.RequestClose()

	mainWindow.ShowAndRun()

}
