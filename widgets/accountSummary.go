package widgets

import (
	"log"
	"pf/dialogs"
	"pf/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/google/uuid"
)

// Declare conformity with Layout interface
var _ fyne.Layout = (*accountSummaryLayout)(nil)

type accountSummaryLayout struct {
	acct *AccountSummary
}

type AccountSummary struct {
	widget.BaseWidget
	name    binding.String
	amount  binding.Float
	dueDate binding.String

	parent fyne.Window

	nameLabel     *widget.Label
	amountLabel   *widget.Label
	dateLabel     *widget.Label
	billButton    *widget.Button
	payButton     *widget.Button
	accountButton *widget.Button
}

const buttonSpacing = 10
const buttonMargin = 20

func (a *AccountSummary) getAcctID() uuid.UUID {
	nStr, _ := a.name.Get()
	acctID := models.GetAccountByName(nStr).ID

	return acctID
}
func addBillCB(a *AccountSummary) {
	dialogs.AddBillDialog(a.getAcctID(), a.parent)
}

func addPayCB(a *AccountSummary) {
	dialogs.PayDialog(a.getAcctID(), a.parent)
}

func modAcctCB(a *AccountSummary) {
	dialogs.ModifyAccountDialog(a.getAcctID(), a.parent)
}

func (a *accountSummaryLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	pos := fyne.NewPos(0, 0)

	height := a.MinSize(objects).Height
	width := size.Width * 0.2

	// Name
	a.acct.nameLabel.Move(pos)
	a.acct.nameLabel.Resize(fyne.NewSize(width, height))

	pos = pos.Add(fyne.NewPos(width, 0))

	// Amount
	a.acct.amountLabel.Move(pos)
	a.acct.amountLabel.Resize(fyne.NewSize(width, height))

	pos = pos.Add(fyne.NewPos(width, 0))

	// Due Date
	a.acct.dateLabel.Move(pos)
	a.acct.dateLabel.Resize(fyne.NewSize(width, height))

	//--- Actions
	// These are layed out starting from the right

	// Bill
	width = a.acct.billButton.MinSize().Width

	pos = fyne.NewPos(size.Width-width-buttonMargin, 0)

	a.acct.billButton.Move(pos)
	a.acct.billButton.Resize(fyne.NewSize(width, height))

	// Pay
	width = a.acct.payButton.MinSize().Width

	pos = pos.Subtract(fyne.NewSize(width+buttonSpacing, 0))

	a.acct.payButton.Move(pos)
	a.acct.payButton.Resize(fyne.NewSize(width, height))

	// Pay
	width = a.acct.accountButton.MinSize().Width

	pos = pos.Subtract(fyne.NewSize(width+buttonSpacing, 0))

	a.acct.accountButton.Move(pos)
	a.acct.accountButton.Resize(fyne.NewSize(width, height))

}

func (a *accountSummaryLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w := float32(0)
	h := float32(0)

	for _, o := range objects {
		size := o.MinSize()
		w += size.Width
		if h < size.Height {
			h = size.Height
		}
	}

	return fyne.NewSize(w, h)
}

func (a *AccountSummary) CreateRenderer() fyne.WidgetRenderer {

	log.Println("Creating renderer for a summary")

	a.nameLabel = widget.NewLabelWithData(a.name)

	amountBinding := binding.NewSprintf("%8.2f", a.amount)
	a.amountLabel = widget.NewLabelWithData(amountBinding)

	a.dateLabel = widget.NewLabelWithData(a.dueDate)

	a.billButton = widget.NewButton("B", func() { addBillCB(a) })
	a.payButton = widget.NewButton("P", func() { addPayCB(a) })
	a.accountButton = widget.NewButton("A", func() { modAcctCB(a) })

	c := container.New(&accountSummaryLayout{a},
		a.nameLabel, a.amountLabel, a.dateLabel,
		a.billButton, a.payButton, a.accountButton,
	)

	return widget.NewSimpleRenderer(c)
}

func NewAccountSummary(w fyne.Window) *AccountSummary {

	a := &AccountSummary{}
	a.ExtendBaseWidget(a)

	a.name = binding.NewString()
	a.SetName("Template")

	a.amount = binding.NewFloat()

	a.dueDate = binding.NewString()
	a.SetDueDate("yyyy-mm-dd")

	a.parent = w
	return a
}

func (a *AccountSummary) SetAmount(b float32) {
	a.amount.Set(float64(b))
}

func (a *AccountSummary) SetName(n string) {
	a.name.Set(n)
}

func (a *AccountSummary) SetDueDate(d string) {
	a.dueDate.Set(d)
}
