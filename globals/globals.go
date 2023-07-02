package globals

import "fyne.io/fyne/v2/widget"

var updateMessageChan = make(chan bool, 1)

func RequestUpdate() {
	updateMessageChan <- true
}

func RequestClose() {
	updateMessageChan <- false
}

func StartUpdater(l *widget.List) {
	go func() {
		for {
			msg := <-updateMessageChan
			if msg {
				l.Refresh()
			} else {
				break
			}
		}
	}()
}
