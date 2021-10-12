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
	gameApp.SetIcon(loadIcon())
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

func loadIcon() fyne.Resource {
	res, err := fyne.LoadResourceFromPath("icon.png")
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("Error loading app icon")
		os.Exit(1)
	}

	return res
}
