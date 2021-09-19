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
	//if head is not in the same direction as the new direction specified, change only the orientation of the head

	if gv.snakeDirection != gv.snakeHead.direction {

		gv.snakeHead.setDirection(gv.snakeDirection)
		return

	}

	gv.moveHead()
	gv.moveBody()

	return

}

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

func (gv *GameVisual) moveBody() {

	var directionToSet string
	for i, snakeBodyPart := range gv.snakeBody {
		if i == 0 {
			snakeBodyPart.part.Position1 = gv.snakeHead.part.Position2
			directionToSet = gv.snakeHead.direction
		} else {
			snakeBodyPart.part.Position1 = gv.snakeBody[i-1].part.Position2
			directionToSet = gv.snakeBody[i-1].direction
		}

		switch directionToSet {

		case leftDirection:
			snakeBodyPart.part.Position2 = fyne.Position{X: snakeBodyPart.part.Position1.X + snakeBodyPartLength, Y: snakeBodyPart.part.Position1.Y}
		case rightDirection:
			snakeBodyPart.part.Position2 = fyne.Position{X: snakeBodyPart.part.Position1.X - snakeBodyPartLength, Y: snakeBodyPart.part.Position1.Y}
		case upDirection:
			snakeBodyPart.part.Position2 = fyne.Position{X: snakeBodyPart.part.Position1.X, Y: snakeBodyPart.part.Position1.Y + snakeBodyPartLength}

		case downDirection:
			snakeBodyPart.part.Position2 = fyne.Position{X: snakeBodyPart.part.Position1.X, Y: snakeBodyPart.part.Position1.Y - snakeBodyPartLength}

		}
		snakeBodyPart.direction = directionToSet

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
		return fyne.NewPos(s.part.Position1.X-partLength, s.part.Position2.Y)

	case rightDirection:
		return fyne.NewPos(s.part.Position1.X+partLength, s.part.Position2.Y)

	default:
		err := errors.New("Unexpected new direction passed when setting direction for snake figure head/body part [from vertical to left or right]")
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
		return fyne.NewPos(s.part.Position2.X, s.part.Position1.Y-partLength)

	case downDirection:
		return fyne.NewPos(s.part.Position2.X, s.part.Position1.Y+partLength)

	default:
		err := errors.New("Unexpected new direction passed when setting direction for snake figure head/body part [from horizontal to up or down]")
		log.WithFields(log.Fields{
			"err":               err.Error(),
			"current_direction": s.direction,
			"new_direction":     direction,
		}).Error("Error turning snake figure")
		return fyne.Position{}

	}
}
