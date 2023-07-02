package dialogs

import (
	"errors"
	"log"
	"pf/globals"
	"pf/models"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func AddAccountDialog(w fyne.Window) {

	t := widget.NewRadioGroup(models.AccountTypes, func(string) {})
	t.Selected = "Payable"

	s := ""
	sb := binding.BindString(&s)
	e := widget.NewEntryWithData(sb)

	d := ""
	db := binding.BindString(&d)
	de := widget.NewEntryWithData(db)

	box := container.New(layout.NewFormLayout(),
		widget.NewLabel("Type"), t,
		widget.NewLabel("Name"), e,
		widget.NewLabel("Balance"), de,
	)

	dialog.ShowCustomConfirm("Add Account", "Add", "Cancel", box, func(a bool) {
		if a {
			if d == "" {
				d = "0.00"
			}
			if dollars, str_err := strconv.ParseFloat(d, 32); str_err == nil {
				log.Printf("Add account balance = %f\n", dollars)
				if _, err := models.AddAccount(s, t.Selected, float32(dollars)); err == nil {
					globals.RequestUpdate()
				} else {
					dialog.ShowError(err, w)
				}
			} else {
				dialog.ShowError(errors.New("balance does not seem to be a number"), w)
			}
		}
	}, w)
}
