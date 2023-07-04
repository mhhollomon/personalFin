package globals

import (
	"log"

	"fyne.io/fyne/v2/widget"
)

var updateMessageChan = make(chan bool, 1)

var DataDir = "./"

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
				log.Println("REfreshing")
				l.Refresh()
			} else {
				break
			}
		}
	}()
}
