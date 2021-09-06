package ui

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"github.com/meghashyamc/snake-game/core"
)

type GameVisual struct {
	snakeFigure  *canvas.Line
	gameOverText *canvas.Text
	Container    *fyne.Container
}

func NewGameVisual() *GameVisual {

	gameVisual := &GameVisual{snakeFigure: &canvas.Line{StrokeColor: theme.ForegroundColor(), StrokeWidth: 10}, gameOverText: &canvas.Text{Alignment: fyne.TextAlignLeading, Color: theme.ForegroundColor(), Text: "Game Over!", TextSize: 20, TextStyle: fyne.TextStyle{Bold: true}}}
	gameVisual.gameOverText.Hide()
	return gameVisual
}

func (gv *GameVisual) Layout(objects []fyne.CanvasObject, size fyne.Size) {

	gv.snakeFigure.Move(fyne.Position{X: gv.snakeFigure.Position1.X - 10, Y: gv.snakeFigure.Position1.Y})
	return

}

func (gv *GameVisual) MinSize(objects []fyne.CanvasObject) fyne.Size {

	minSize := os.Getenv("MIN_GRID_SIZE")
	minSizeInt, err := strconv.Atoi(minSize)
	if err != nil {
		panic(err)
	}
	return fyne.NewSize(float32(minSizeInt), float32(minSizeInt))
}

func (gv *GameVisual) InitContainer() {
	container := fyne.NewContainer(gv.snakeFigure, gv.gameOverText)

	container.Layout = gv
	gv.Container = container
	snakeFigureLength := os.Getenv("STARTING_SNAKE_LEN")
	snakeFigureLengthInt, err := strconv.Atoi(snakeFigureLength)
	if err != nil {
		panic(err)
	}

	gv.snakeFigure.Position1 = fyne.Position{X: container.MinSize().Width / 2, Y: container.MinSize().Height / 2}

	gv.snakeFigure.Position2 = fyne.Position{X: gv.snakeFigure.Position1.X + float32(snakeFigureLengthInt), Y: gv.snakeFigure.Position1.Y}

}

func (gv *GameVisual) Animate() {
	tick := time.NewTicker(time.Second)
	for {
		select {
		case <-tick.C:
			containerSize := gv.Container.Size()
			gv.Layout(nil, containerSize)
			canvas.Refresh(gv.Container)
			if core.IsGameOver(&gv.snakeFigure.Position1, &containerSize) {
				fmt.Println("Reached here!")
				gv.snakeFigure.Hide()
				gv.gameOverText.Show()
				canvas.Refresh(gv.Container)
				return
			}

		default:
			time.Sleep(1 * time.Millisecond)

		}
	}
}
