package dialogs

import (
	"errors"
	"pf/globals"
	"pf/models"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"
)

func calcDueDate(id uuid.UUID) string {

	acctObj := models.GetAccountById(id)

	defaultDay := acctObj.DueDay

	due := time.Now()

	month := due.Month()

	if defaultDay < due.Day() {
		month += 1
	}

	due = time.Date(due.Year(), month, defaultDay, 0, 0, 0, 0, time.UTC)

	return due.Format("2006-01-02")

}

func AddBillDialog(acct uuid.UUID, w fyne.Window) {

	amountEntry := widget.NewEntry()
	dateEntry := widget.NewEntry()

	dateEntry.SetText(calcDueDate(acct))

	box := container.New(layout.NewFormLayout(),
		widget.NewLabel("Amount"), amountEntry,
		widget.NewLabel("Due Date"), dateEntry,
	)

	dialog.ShowCustomConfirm("Add Bill", "Add", "Cancel", box, func(a bool) {
		if a {
			dollars, str_err := strconv.ParseFloat(amountEntry.Text, 32)
			if str_err != nil {
				dialog.ShowError(errors.New("balance does not seem to be a number"), w)
				return
			}

			due, t_err := time.Parse("2006-01-02", dateEntry.Text)
			if t_err != nil {
				dialog.ShowError(errors.New("date could not be parse (yyy-mm-dd)"), w)
				return
			}

			models.NewBill(acct, float32(dollars), due)
			globals.RequestUpdate()
		}

	}, w)

}
