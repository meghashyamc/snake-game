package ui

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"github.com/meghashyamc/snake-game/decisions"
)

type GameVisual struct {
	snakeDirection string
	snakeHead      *canvas.Line
	snakeBody      []*canvas.Line
	gameOverText   *canvas.Text
	Container      *fyne.Container
}

func NewGameVisual() *GameVisual {
	snakeHeadColor = theme.FocusColor()
	snakeBodyColor = theme.ForegroundColor()
	gameOverTextColor = theme.ForegroundColor()

	gameVisual := &GameVisual{snakeHead: &canvas.Line{StrokeColor: snakeHeadColor, StrokeWidth: snakeHeadWidth}, gameOverText: &canvas.Text{Alignment: gameOverTextAlignment, Color: gameOverTextColor, Text: gameOverText, TextSize: gameOverTextSize, TextStyle: gameOverTextStyle}}
	gameVisual.createNewSnakeBody()
	gameVisual.gameOverText.Hide()
	return gameVisual
}

func (gv *GameVisual) createNewSnakeBody() {
	gv.snakeBody = [](*canvas.Line){}
	for i := 0; i < numOfStartingBodyParts; i++ {
		gv.snakeBody = append(gv.snakeBody, &canvas.Line{StrokeColor: snakeBodyColor, StrokeWidth: snakeBodyWidth})
	}
}

func (gv *GameVisual) Layout(objects []fyne.CanvasObject, size fyne.Size) {

	gv.snakeHead.Move(fyne.Position{X: gv.snakeHead.Position1.X - snakeSpeed, Y: gv.snakeHead.Position1.Y})

	for i := 0; i < len(gv.snakeBody); i++ {
		gv.snakeBody[i].Move(fyne.Position{X: gv.snakeBody[i].Position1.X - snakeSpeed, Y: gv.snakeBody[i].Position1.Y})
	}
	return

}

func (gv *GameVisual) MinSize(objects []fyne.CanvasObject) fyne.Size {

	return fyne.NewSize(minGridSize, minGridSize)
}

func (gv *GameVisual) InitContainer() {

	containerParts := []fyne.CanvasObject{}
	containerParts = append(containerParts, gv.snakeHead)
	for _, snakeBodyPart := range gv.snakeBody {
		containerParts = append(containerParts, snakeBodyPart)
	}
	containerParts = append(containerParts, gv.gameOverText)

	container := fyne.NewContainer(containerParts...)

	container.Layout = gv
	gv.Container = container
	gv.setSnakeFigureInitialPositions()

}

func (gv *GameVisual) setSnakeFigureInitialPositions() {

	gv.snakeHead.Position1 = fyne.Position{X: gv.Container.MinSize().Width / 2, Y: gv.Container.MinSize().Height / 2}
	gv.snakeHead.Position2 = fyne.Position{X: gv.snakeHead.Position1.X + snakeHeadLength, Y: gv.snakeHead.Position1.Y}

	for i := 0; i < len(gv.snakeBody); i++ {

		if i == 0 {
			gv.snakeBody[i].Position1 = gv.snakeHead.Position2

		} else {
			gv.snakeBody[i].Position1 = gv.snakeBody[i-1].Position2
		}
		gv.snakeBody[i].Position2 = fyne.Position{X: gv.snakeBody[i].Position1.X + snakeBodyPartLength, Y: gv.snakeBody[i].Position1.Y}

	}
}
func (gv *GameVisual) hideSnakeFigure() {

	gv.snakeHead.Hide()
	for i := 0; i < len(gv.snakeBody); i++ {
		gv.snakeBody[i].Hide()
	}
}

func (gv *GameVisual) Animate() {
	tick := time.NewTicker(time.Second)
	for {
		select {
		case <-tick.C:
			containerSize := gv.Container.Size()
			gv.Layout(nil, containerSize)
			canvas.Refresh(gv.Container)
			if decisions.IsGameOver(&gv.snakeHead.Position1, &containerSize) {
				gv.hideSnakeFigure()
				gv.gameOverText.Show()
				canvas.Refresh(gv.Container)
				return
			}
		case val, ok := <-directionC:
			if !ok {
				panic("Error in snake direction data received")
			}
			gv.snakeDirection = val
			gv.Layout(nil, gv.Container.Size())
			canvas.Refresh(gv.Container)

		default:
			time.Sleep(1 * time.Millisecond)

		}
	}
}
