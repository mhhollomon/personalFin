package layouts

import "fyne.io/fyne/v2"

type VFlex struct {
	MinHeight float32
	MinWidth  float32
}

func NewVFlex(w float32, h float32) *VFlex {
	return &VFlex{MinWidth: w, MinHeight: h}
}

func (v *VFlex) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)

	for _, o := range objects {
		childSize := o.MinSize()

		if childSize.Width > w {
			w = childSize.Width
		}

		h += childSize.Height
	}

	if w < v.MinWidth {
		w = v.MinWidth
	}

	if h < v.MinHeight {
		h = v.MinHeight
	}

	return fyne.NewSize(w, h)

}

func (v *VFlex) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(0, 0)
	lastIndex := len(objects) - 1

	for i, o := range objects {
		size := o.MinSize()

		// Fix the width of the kids to the container
		if containerSize.Width > size.Width {
			size.Width = containerSize.Width
		}

		// For all but the last child, set the height the
		// their requested Min
		// Let the last child take up the rest of the room.

		if i == lastIndex {
			heightLeft := containerSize.Height - pos.Y
			if heightLeft > size.Height {
				size.Height = heightLeft
			}
		}
		o.Resize(size)
		o.Move(pos)
		pos = pos.Add(fyne.NewPos(0, size.Height))
	}
}
