package main

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/meghashyamc/snake-game/ui"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	gameApp := app.New()
	window := gameApp.NewWindow("Snake Game")

	newGameVisual, err := formGameVisual(window)
	if err != nil {
		os.Exit(1)
	}
	window.SetContent(newGameVisual.Container)
	go ui.ListenToInput(window)
	window.ShowAndRun()

}

func formGameVisual(win fyne.Window) (*ui.GameVisual, error) {

	thisGameVisual, err := ui.NewGameVisual()
	if err != nil {
		return nil, err
	}
	thisGameVisual.InitContainer()

	go thisGameVisual.Animate()
	return thisGameVisual, nil

}
