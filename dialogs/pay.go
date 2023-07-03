package dialogs

import (
	"errors"
	"fmt"
	"log"
	"pf/globals"
	"pf/models"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"
)

func PayDialog(acct uuid.UUID, w fyne.Window) {

	bills := models.FindBillsForAcct(acct)

	if len(bills) == 0 {
		log.Println("No bills to pay for", acct)
		return
	}

	billMap := make(map[string]*models.Bill)
	options := make([]string, 0, 10)

	for _, b := range bills {
		dateStr := b.DueDate.Format("2006-01-02")
		str := strings.TrimSpace(fmt.Sprintf("%8.2f %s", b.Amount, dateStr))

		billMap[str] = b
		options = append(options, str)
	}

	amountEntry := widget.NewEntry()
	amountEntry.SetText(fmt.Sprintf("%8.2f", billMap[options[0]].Amount))
	defaultDate := time.Now().Format("2006-01-02")
	dateEntry := widget.NewEntry()
	dateEntry.SetText(defaultDate)

	selector := widget.NewSelect(options, func(s string) {
		amountEntry.SetText(strings.TrimSpace(fmt.Sprintf("%8.2f", billMap[s].Amount)))

	})
	selector.SetSelectedIndex(0)

	box := container.New(layout.NewFormLayout(),
		widget.NewLabel("Bill"), selector,
		widget.NewLabel("Amount"), amountEntry,
		widget.NewLabel("Date"), dateEntry,
	)

	dialog.ShowCustomConfirm("Pay Bill", "Pay", "Cancel", box, func(a bool) {
		if a {

			targetBill := billMap[selector.Selected]
			dollars, str_err := strconv.ParseFloat(strings.TrimSpace(amountEntry.Text), 32)
			if str_err != nil {
				dialog.ShowError(errors.New("amount does not seem to be a number"), w)
				return
			}

			if dollars > float64(targetBill.OrigAmount) {
				dialog.ShowError(errors.New("amount is more than what is outstanding"), w)
			}

			due, t_err := time.Parse("2006-01-02", dateEntry.Text)
			if t_err != nil {
				dialog.ShowError(errors.New("date could not be parse (yyy-mm-dd)"), w)
				return
			}

			targetBill.Pay(float32(dollars), due)

			globals.RequestUpdate()
		}

	}, w)

}
