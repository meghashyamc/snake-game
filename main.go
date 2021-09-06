package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/joho/godotenv"
	"github.com/meghashyamc/snake-game/ui"
)

func main() {
	godotenv.Load()
	gameApp := app.New()
	window := gameApp.NewWindow("Snake Game")

	window.SetContent(formGameVisual(window).Container)
	window.ShowAndRun()

}

func formGameVisual(win fyne.Window) *ui.GameVisual {

	thisGameVisual := ui.NewGameVisual()
	thisGameVisual.InitContainer()

	go thisGameVisual.Animate()
	return thisGameVisual

}
