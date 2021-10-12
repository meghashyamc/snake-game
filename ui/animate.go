package ui

import (
	"errors"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/meghashyamc/snake-game/checks"
	log "github.com/sirupsen/logrus"
)

func (gv *GameVisual) Animate() {
	tick := time.NewTicker(time.Second)
	for {
		select {
		case <-tick.C:
			gv.update(nextTick)

			gameState := checks.CheckGameState(gv.formGameStateMap())
			if gameState != "" && gv.update(gameState) {
				return
			}

		case val, ok := <-directionC:
			if !ok {
				err := errors.New("Direction data channel is closed and empty")
				log.WithFields(log.Fields{
					"err": err.Error(),
				}).Error("Error animating snake figure")
				return
			}
			gv.snakeDirection = val
			gv.update(directionChanged)

		default:
			time.Sleep(1 * time.Millisecond)

		}
	}
}

func (gv *GameVisual) update(gameState string) bool {
	containerSize := gv.Container.Size()

	var quitGame bool
	switch gameState {
	case nextTick:
		gv.Layout(nil, containerSize)

	case gameOver:
		gv.hideSnakeFigure()
		gv.hideFood()
		gv.gameOverText.Show()
		quitGame = true
	case foodEaten:
		gv.foodParticle.Move(getRandomPositionInGrid(containerSize.Width))
		gv.growSnakeBody()
	case directionChanged:
		gv.Layout(nil, gv.Container.Size())

	}

	canvas.Refresh(gv.Container)
	return quitGame
}

//new layout of objects - on tick or as directed by the user
func (gv *GameVisual) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	//if head's direction is opposite to the new direction specified
	if gv.snakeDirection == oppositeDirections[gv.snakeHead.direction] {
		return
	}
	//if head is not in the same direction as the new direction specified, change only the orientation of the head

	if gv.snakeDirection != gv.snakeHead.direction {

		gv.snakeHead.setDirection(gv.snakeDirection)
		return

	}

	gv.moveHead()
	gv.moveBody()

	return

}

//moves snake head in the direction of the head
func (gv *GameVisual) moveHead() {

	switch gv.snakeHead.direction {

	case leftDirection:
		gv.snakeHead.part.Move(fyne.Position{X: gv.snakeHead.part.Position1.X - snakeSpeed, Y: gv.snakeHead.part.Position1.Y})
	case rightDirection:
		gv.snakeHead.part.Move(fyne.Position{X: gv.snakeHead.part.Position1.X + snakeSpeed, Y: gv.snakeHead.part.Position1.Y})
	case upDirection:
		gv.snakeHead.part.Move(fyne.Position{X: gv.snakeHead.part.Position1.X, Y: gv.snakeHead.part.Position1.Y - snakeSpeed})
	case downDirection:
		gv.snakeHead.part.Move(fyne.Position{X: gv.snakeHead.part.Position1.X, Y: gv.snakeHead.part.Position1.Y + snakeSpeed})

	}

}

//moves a snake part in the direction of the previous snake part/head
func (gv *GameVisual) moveBody() {

	var directionToSet string
	for i, snakeBodyPart := range gv.snakeBody {
		if i == 0 {
			directionToSet = gv.snakeHead.direction
		} else {
			directionToSet = gv.snakeBody[i-1].direction
		}
		switch directionToSet {
		case leftDirection, upDirection:
			snakeBodyPart.moveBodyPartLeftOrUp(gv, i, directionToSet)

		case rightDirection, downDirection:
			snakeBodyPart.moveBodyPartRightOrDown(gv, i, directionToSet)
		default:
			err := errors.New("Unknown current direction passed when moving  snake figure body part")
			log.WithFields(log.Fields{
				"err":               err.Error(),
				"current_direction": snakeBodyPart.direction,
				"new_direction":     directionToSet,
			}).Error("Error moving snake figure during turn")
			os.Exit(1)

		}

		snakeBodyPart.direction = directionToSet

	}

}

//changes the direction of the snake head/snake part (currently only the head's direction is being changed like this)
func (s *snakePart) setDirection(direction string) {

	switch s.direction {

	case upDirection, downDirection:
		s.part.Position1, s.part.Position2 = s.getPositionAfterVerticalTurn(direction)

	case leftDirection, rightDirection:
		s.part.Position1, s.part.Position2 = s.getPositionAfterHorizontalTurn(direction)

	default:
		err := errors.New("Unknown current direction passed when setting direction for snake figure head/body part")
		log.WithFields(log.Fields{
			"err":               err.Error(),
			"current_direction": s.direction,
			"new_direction":     direction,
		}).Error("Error turning snake figure")
		os.Exit(1)
	}

	if s.part.Position1.Subtract(fyne.Position{}).IsZero() {
		os.Exit(1)
	}

	s.direction = direction

}

func (gv *GameVisual) formGameStateMap() map[string]float32 {
	gameStateMap := make(map[string]float32)
	gameStateMap[snakeHeadX] = gv.snakeHead.part.Position1.X
	gameStateMap[snakeHeadY] = gv.snakeHead.part.Position1.Y
	gameStateMap[foodParticleCentreX] = (gv.foodParticle.Position1.X + gv.foodParticle.Position2.X) / 2
	gameStateMap[foodParticleCentreY] = (gv.foodParticle.Position1.Y + gv.foodParticle.Position2.Y) / 2
	gameStateMap[foodParticleDiameter] = foodDiameter
	gameStateMap[gridSize] = gv.Container.Size().Width
	return gameStateMap
}

func (gv *GameVisual) growSnakeBody() {

	snakeBodyPart, _ := newSnakePart(bodyPart)
	gv.snakeBody = append(gv.snakeBody, snakeBodyPart)
	lastIndex := len(gv.snakeBody) - 1
	gv.Container.Objects = append(gv.Container.Objects, gv.snakeBody[lastIndex-1].part)
	gv.snakeBody[lastIndex].part.Position1 = gv.snakeBody[lastIndex-1].part.Position2
	gv.snakeBody[lastIndex].part.Position2 = fyne.Position{X: gv.snakeBody[lastIndex].part.Position1.X + snakeBodyPartLength, Y: gv.snakeBody[lastIndex].part.Position1.Y}

}
