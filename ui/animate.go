package ui

import (
	"errors"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/meghashyamc/snake-game/game"
	log "github.com/sirupsen/logrus"
)

func (gv *GameVisual) Animate() {
	tick := time.NewTicker(time.Second)
	for {
		select {
		case <-tick.C:
			containerSize := gv.Container.Size()
			gv.Layout(nil, containerSize)
			canvas.Refresh(gv.Container)
			if game.IsGameOver(&gv.snakeHead.part.Position1, &containerSize) {
				gv.hideSnakeFigure()
				gv.gameOverText.Show()
				canvas.Refresh(gv.Container)
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
			gv.Layout(nil, gv.Container.Size())
			canvas.Refresh(gv.Container)

		default:
			time.Sleep(1 * time.Millisecond)

		}
	}
}

//new layout of objects - on tick or as directed by the user
func (gv *GameVisual) Layout(objects []fyne.CanvasObject, size fyne.Size) {

	//if head's direction is opposite to the new direction specified
	if gv.snakeDirection == oppositeDirections[gv.snakeHead.direction] {
		return
	}
	//if head and tail are already in the same direction as the new direction specified, move further in the same direction
	if gv.snakeDirection == gv.snakeHead.direction && gv.snakeDirection == gv.snakeBody[len(gv.snakeBody)-1].direction {
		gv.snakeHead.move()

		for i := 0; i < len(gv.snakeBody); i++ {
			gv.snakeBody[i].move()
		}
		return
	}

	//if head is not in the same direction as the new direction specified, change only the orientation of the head
	if gv.snakeDirection != gv.snakeDirection {
		gv.snakeHead.setDirection(gv.snakeDirection)
		return

	}

	//if the head is in the same direction as the new direction specified but at least one body part is not in the same direction
	gv.snakeHead.move()
	turningIndices := gv.getIndicesOfBodyPartsToTurn()
	for i := 0; i < len(gv.snakeBody); i++ {
		gv.snakeBody[i].move()
	}
	for _, turningIndex := range turningIndices {
		gv.snakeBody[turningIndex].setDirection(gv.snakeDirection)
	}

}

func (gv *GameVisual) getIndicesOfBodyPartsToTurn() []int {

	indicesOfBodyPartsToTurn := []int{}
	baseDirection := gv.snakeHead.direction

	for i, snakeBodyPart := range gv.snakeBody {

		if snakeBodyPart.direction != baseDirection {
			indicesOfBodyPartsToTurn = append(indicesOfBodyPartsToTurn, i)
			baseDirection = snakeBodyPart.direction
		}
	}

	return indicesOfBodyPartsToTurn
}

func (s *snakePart) move() {

	switch s.direction {

	case leftDirection:
		s.part.Move(fyne.Position{X: s.part.Position1.X - snakeSpeed, Y: s.part.Position1.Y})
	case rightDirection:
		s.part.Move(fyne.Position{X: s.part.Position1.X + snakeSpeed, Y: s.part.Position1.Y})
	case upDirection:
		s.part.Move(fyne.Position{X: s.part.Position1.X, Y: s.part.Position1.Y - snakeSpeed})
	case downDirection:
		s.part.Move(fyne.Position{X: s.part.Position1.X, Y: s.part.Position1.Y + snakeSpeed})

	}

}

func (s *snakePart) setDirection(direction string) {

	switch s.direction {

	case upDirection, downDirection:
		s.part.Position1 = s.getPositionAfterVerticalTurn(direction)

	case leftDirection, rightDirection:
		s.part.Position1 = s.getPositionAfterHorizontalTurn(direction)

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

func (s *snakePart) getPositionAfterVerticalTurn(direction string) fyne.Position {
	var partLength float32
	if s.isHead {
		partLength = snakeHeadLength
	} else {
		partLength = snakeBodyPartLength
	}
	switch direction {
	case leftDirection:
		return fyne.NewPos(s.part.Position1.X-partLength, s.part.Position1.Y)

	case rightDirection:
		return fyne.NewPos(s.part.Position1.X+partLength, s.part.Position1.Y)

	default:
		err := errors.New("Unexpected new direction passed when setting direction for snake figure head/body part [from up to left or right]")
		log.WithFields(log.Fields{
			"err":               err.Error(),
			"current_direction": s.direction,
			"new_direction":     direction,
		}).Error("Error turning snake figure")
		return fyne.Position{}

	}
}

func (s *snakePart) getPositionAfterHorizontalTurn(direction string) fyne.Position {

	var partLength float32
	if s.isHead {
		partLength = snakeHeadLength
	} else {
		partLength = snakeBodyPartLength
	}
	switch direction {
	case upDirection:
		return fyne.NewPos(s.part.Position1.X, s.part.Position1.Y-partLength)

	case rightDirection:
		return fyne.NewPos(s.part.Position1.X, s.part.Position1.Y+partLength)

	default:
		err := errors.New("Unexpected new direction passed when setting direction for snake figure head/body part [from up to left or right]")
		log.WithFields(log.Fields{
			"err":               err.Error(),
			"current_direction": s.direction,
			"new_direction":     direction,
		}).Error("Error turning snake figure")
		return fyne.Position{}

	}
}
