package widgets

import (
	"log"
	"pf/dialogs"
	"pf/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

// Declare conformity with Layout interface
var _ fyne.Layout = (*accountSummaryLayout)(nil)

type accountSummaryLayout struct {
	acct *AccountSummary
}

type AccountSummary struct {
	widget.BaseWidget
	name   binding.String
	amount binding.Float

	parent fyne.Window

	nameLabel   *widget.Label
	amountLabel *widget.Label
	billButton  *widget.Button
}

const buttonSpacing = 20

func addBillCB(a *AccountSummary) {
	nStr, _ := a.name.Get()
	acctID := models.GetAccountByName(nStr).ID

	dialogs.AddBillDialog(acctID, a.parent)

}

func (a *accountSummaryLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	pos := fyne.NewPos(0, 0)

	height := a.MinSize(objects).Height

	// Name
	a.acct.nameLabel.Move(pos)

	width := size.Width * 0.3
	a.acct.nameLabel.Resize(fyne.NewSize(width, height))

	pos = pos.Add(fyne.NewPos(width, 0))

	// Amount
	a.acct.amountLabel.Move(pos)
	width = a.acct.amountLabel.MinSize().Width

	a.acct.amountLabel.Resize(fyne.NewSize(width, height))

	//Actions
	// These are layed out starting from the right
	width = a.acct.billButton.MinSize().Width

	pos = fyne.NewPos(size.Width-width-buttonSpacing, 0)

	a.acct.billButton.Move(pos)
	a.acct.billButton.Resize(fyne.NewSize(width, height))

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

	a.billButton = widget.NewButton("B", func() { addBillCB(a) })

	c := container.New(&accountSummaryLayout{a}, a.nameLabel, a.amountLabel, a.billButton)

	return widget.NewSimpleRenderer(c)
}

func NewAccountSummary(w fyne.Window) *AccountSummary {

	a := &AccountSummary{}
	a.ExtendBaseWidget(a)
	a.name = binding.NewString()
	a.SetName("Template")
	a.amount = binding.NewFloat()
	a.parent = w
	return a
}

func (a *AccountSummary) SetAmount(b float32) {
	a.amount.Set(float64(b))
}

func (a *AccountSummary) SetName(n string) {
	a.name.Set(n)
}
