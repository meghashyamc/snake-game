package ui

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

var directionC = make(chan string)

func ListenToInput(window fyne.Window) {

	if deskCanvas, ok := window.Canvas().(desktop.Canvas); ok {
		deskCanvas.SetOnKeyDown(func(key *fyne.KeyEvent) {

			if directionKeys[string(key.Name)] {
				directionC <- string(key.Name)
			}
		})

	} else {
		panic(errors.New("Desktop not detected"))
	}
}
