package dialogs

import (
	"errors"
	"fmt"
	"pf/globals"
	"pf/models"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"
)

func ModifyAccountDialog(acct uuid.UUID, w fyne.Window) {

	acctObj := models.GetAccountById(acct)

	nameEntry := widget.NewEntry()
	nameEntry.SetText(acctObj.Name)

	dayEntry := widget.NewEntry()
	dayEntry.SetText(fmt.Sprintf("%d", acctObj.DueDay))

	box := container.New(layout.NewFormLayout(),
		widget.NewLabel("Name"), nameEntry,
		widget.NewLabel("Usual Due Day"), dayEntry,
	)

	dialog.ShowCustomConfirm("Update Account", "Apply", "Cancel", box, func(a bool) {
		if a {

			needUpdate := false

			day, str_err := strconv.ParseUint(strings.TrimSpace(dayEntry.Text), 10, 32)
			if str_err != nil || day > 31 || day == 0 {
				dialog.ShowError(errors.New("day must be a number between 1 and 31"), w)
				return
			}

			newName := strings.TrimSpace(nameEntry.Text)
			if newName != acctObj.Name {
				newAcct := models.GetAccountByName(newName)
				if newAcct == nil {
					acctObj.SetName(newName)
					needUpdate = true
				}
			}

			if day != uint64(acctObj.DueDay) {
				acctObj.SetDueDay(int(day))
				needUpdate = true
			}

			if needUpdate {
				globals.RequestUpdate()
			}
		}

	}, w)

}
