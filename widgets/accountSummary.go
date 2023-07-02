package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

// Declare conformity with Layout interface
var _ fyne.Layout = (*accountSummaryLayout)(nil)

type accountSummaryLayout struct{}

type AccountSummary struct {
	widget.BaseWidget
	name   binding.String
	amount binding.Float

	nameLabel   *widget.Label
	amountLabel *widget.Label
}

func (a *accountSummaryLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	pos := fyne.NewPos(0, 0)

	height := a.MinSize(objects).Height

	objects[0].Move(pos)

	width := size.Width * 0.3
	objects[0].Resize(fyne.NewSize(width, height))

	pos = pos.Add(fyne.NewPos(width, 0))

	objects[1].Move(pos)

	width = size.Width - width
	objects[1].Resize(fyne.NewSize(width, height))

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

	a.nameLabel = widget.NewLabelWithData(a.name)

	amountBinding := binding.NewSprintf("%8.2f", a.amount)
	a.amountLabel = widget.NewLabelWithData(amountBinding)

	c := container.New(&accountSummaryLayout{}, a.nameLabel, a.amountLabel)

	return widget.NewSimpleRenderer(c)
}

func NewBlankAccountSummary() *AccountSummary {

	a := &AccountSummary{}
	a.ExtendBaseWidget(a)
	a.name = binding.NewString()
	a.SetName("Template")
	a.amount = binding.NewFloat()
	return a
}

func (a *AccountSummary) SetAmount(b float32) {
	a.amount.Set(float64(b))
}

func (a *AccountSummary) SetName(n string) {
	a.name.Set(n)
}
