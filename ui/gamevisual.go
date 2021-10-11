package ui

import (
	"errors"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"

	log "github.com/sirupsen/logrus"
)

type GameVisual struct {
	snakeDirection string
	snakeHead      *snakePart
	snakeBody      []*snakePart
	foodParticle   *canvas.Circle
	gameOverText   *canvas.Text
	Container      *fyne.Container
}

type snakePart struct {
	part      *canvas.Line
	isHead    bool
	direction string
}

func NewGameVisual() (*GameVisual, error) {
	initColors()

	snakeHead, err := newSnakePart(headPart)
	if err != nil {
		return nil, err
	}
	gameVisual := &GameVisual{snakeHead: snakeHead, snakeDirection: leftDirection, gameOverText: newGameOverText()}
	gameVisual.newFoodParticle(minGridSize)
	gameVisual.snakeBody = [](*snakePart){}
	for i := 0; i < numOfStartingBodyParts; i++ {
		snakeBodyPart, err := newSnakePart(bodyPart)
		if err != nil {
			return nil, err
		}
		gameVisual.snakeBody = append(gameVisual.snakeBody, snakeBodyPart)
	}
	gameVisual.gameOverText.Hide()
	return gameVisual, nil
}

func (gv *GameVisual) newFoodParticle(gridSize float32) {
	foodParticle := &canvas.Circle{FillColor: foodParticleColor, StrokeColor: foodParticleColor, Position1: getRandomPositionInGrid(gridSize)}
	foodParticle.Position2 = fyne.Position{X: foodParticle.Position1.X + enclosedSquareInsideCircleSide, Y: foodParticle.Position1.Y + enclosedSquareInsideCircleSide}
	gv.foodParticle = foodParticle

}

func newGameOverText() *canvas.Text {
	return &canvas.Text{Alignment: gameOverTextAlignment, Color: gameOverTextColor, Text: gameOverText, TextSize: gameOverTextSize, TextStyle: gameOverTextStyle}
}

func newSnakePart(partName string) (*snakePart, error) {

	switch partName {

	case headPart:
		return &snakePart{isHead: true, part: &canvas.Line{StrokeColor: snakeHeadColor, StrokeWidth: snakeHeadWidth}, direction: leftDirection}, nil

	case bodyPart:
		return &snakePart{part: &canvas.Line{StrokeColor: snakeBodyColor, StrokeWidth: snakeBodyWidth}, direction: leftDirection}, nil

	default:
		err := errors.New("Unknown snake body part requested")
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Error("Error initializing snake figure")
		return nil, err
	}

}

func (gv *GameVisual) MinSize(objects []fyne.CanvasObject) fyne.Size {

	return fyne.NewSize(minGridSize, minGridSize)
}

func (gv *GameVisual) InitContainer() {

	containerParts := []fyne.CanvasObject{}
	containerParts = append(containerParts, gv.snakeHead.part)
	for _, snakeBodyPart := range gv.snakeBody {
		containerParts = append(containerParts, snakeBodyPart.part)
	}
	containerParts = append(containerParts, gv.gameOverText)
	containerParts = append(containerParts, gv.foodParticle)

	container := fyne.NewContainer(containerParts...)

	container.Layout = gv
	gv.Container = container
	gv.setSnakeFigureInitialPositions()

}

func (gv *GameVisual) setSnakeFigureInitialPositions() {

	gv.snakeHead.part.Position1 = fyne.Position{X: gv.Container.MinSize().Width / 2, Y: gv.Container.MinSize().Height / 2}
	gv.snakeHead.part.Position2 = fyne.Position{X: gv.snakeHead.part.Position1.X + snakeHeadLength, Y: gv.snakeHead.part.Position1.Y}

	for i := 0; i < len(gv.snakeBody); i++ {

		if i == 0 {
			gv.snakeBody[i].part.Position1 = gv.snakeHead.part.Position2

		} else {
			gv.snakeBody[i].part.Position1 = gv.snakeBody[i-1].part.Position2
		}
		gv.snakeBody[i].part.Position2 = fyne.Position{X: gv.snakeBody[i].part.Position1.X + snakeBodyPartLength, Y: gv.snakeBody[i].part.Position1.Y}

	}
}
func (gv *GameVisual) hideSnakeFigure() {

	gv.snakeHead.part.Hide()
	for i := 0; i < len(gv.snakeBody); i++ {
		gv.snakeBody[i].part.Hide()
	}
}

func (gv *GameVisual) hideFood() {

	gv.foodParticle.Hide()

}

func initColors() {
	snakeHeadColor = theme.FocusColor()
	snakeBodyColor = theme.ForegroundColor()
	gameOverTextColor = theme.ForegroundColor()
}

func getRandomPositionInGrid(gridSize float32) fyne.Position {

	randSource := rand.NewSource(time.Now().UnixNano())
	randX := float32(rand.New(randSource).Intn(int(gridSize)))
	randY := float32(rand.New(randSource).Intn(int(gridSize)))

	return fyne.Position{X: randX, Y: randY}
}
