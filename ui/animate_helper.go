package ui

import (
	"errors"

	"fyne.io/fyne/v2"
	log "github.com/sirupsen/logrus"
)

func (s *snakePart) moveBodyPartLeftOrUp(gv *GameVisual, snakePartIndex int, directionToSet string) {

	if snakePartIndex == 0 {
		s.part.Position1 = gv.snakeHead.part.Position2
	} else {
		s.part.Position1 = gv.snakeBody[snakePartIndex-1].part.Position2
	}

	switch directionToSet {

	case leftDirection:
		s.part.Position2 = fyne.Position{X: s.part.Position1.X + snakeBodyPartLength, Y: s.part.Position1.Y}

	case upDirection:
		s.part.Position2 = fyne.Position{X: s.part.Position1.X, Y: s.part.Position1.Y + snakeBodyPartLength}

	}

}

func (s *snakePart) moveBodyPartRightOrDown(gv *GameVisual, snakePartIndex int, directionToSet string) {

	if snakePartIndex == 0 {
		s.part.Position2 = gv.snakeHead.part.Position1
	} else {
		s.part.Position2 = gv.snakeBody[snakePartIndex-1].part.Position1
	}

	switch directionToSet {
	case rightDirection:
		s.part.Position1 = fyne.Position{X: s.part.Position2.X - snakeBodyPartLength, Y: s.part.Position2.Y}

	case downDirection:
		s.part.Position1 = fyne.Position{X: s.part.Position2.X, Y: s.part.Position2.Y - snakeBodyPartLength}

	}

}

func (s *snakePart) getPositionAfterVerticalTurn(direction string) (fyne.Position, fyne.Position) {
	var partLength float32
	if s.isHead {
		partLength = snakeHeadLength
	} else {
		partLength = snakeBodyPartLength
	}

	newDirection := direction

	switch newDirection {
	case leftDirection:
		return s.getPositionAfterVerticalLeftTurn(partLength)

	case rightDirection:
		return s.getPositionAfterVerticalRightTurn(partLength)

	default:
		err := errors.New("Unexpected new direction passed when setting direction for snake figure head/body part [from vertical to left or right]")
		log.WithFields(log.Fields{
			"err":               err.Error(),
			"current_direction": s.direction,
			"new_direction":     direction,
		}).Error("Error turning snake figure")
		return fyne.Position{}, fyne.Position{}

	}
}

func (s *snakePart) getPositionAfterHorizontalTurn(direction string) (fyne.Position, fyne.Position) {

	var partLength float32
	if s.isHead {
		partLength = snakeHeadLength
	} else {
		partLength = snakeBodyPartLength
	}

	newDirection := direction
	switch newDirection {
	case upDirection:
		return s.getPositionAfterHorizontalUpTurn(partLength)

	case downDirection:
		return s.getPositionAfterHorizontalDownTurn(partLength)

	default:
		err := errors.New("Unexpected new direction passed when setting direction for snake figure head/body part [from horizontal to up or down]")
		log.WithFields(log.Fields{
			"err":               err.Error(),
			"current_direction": s.direction,
			"new_direction":     direction,
		}).Error("Error turning snake figure")
		return fyne.Position{}, fyne.Position{}

	}
}

func (s *snakePart) getPositionAfterVerticalLeftTurn(partLength float32) (fyne.Position, fyne.Position) {
	switch s.direction {
	case upDirection:
		return fyne.NewPos(s.part.Position2.X-partLength, s.part.Position2.Y), s.part.Position2
	case downDirection:
		return fyne.NewPos(s.part.Position1.X-partLength, s.part.Position1.Y), s.part.Position1
	default:
		err := errors.New("Unexpected current direction passed when setting direction for snake figure head/body part [from vertical to left]")
		log.WithFields(log.Fields{
			"err":               err.Error(),
			"current_direction": s.direction,
			"new_direction":     leftDirection,
		}).Error("Error turning snake figure")
		return fyne.Position{}, fyne.Position{}
	}
}

func (s *snakePart) getPositionAfterVerticalRightTurn(partLength float32) (fyne.Position, fyne.Position) {

	switch s.direction {
	case upDirection:
		return s.part.Position2, fyne.NewPos(s.part.Position2.X+partLength, s.part.Position2.Y)
	case downDirection:
		return s.part.Position1, fyne.NewPos(s.part.Position1.X+partLength, s.part.Position1.Y)
	default:
		err := errors.New("Unexpected current direction passed when setting direction for snake figure head/body part [from vertical to right]")
		log.WithFields(log.Fields{
			"err":               err.Error(),
			"current_direction": s.direction,
			"new_direction":     rightDirection,
		}).Error("Error turning snake figure")
		return fyne.Position{}, fyne.Position{}
	}

}

func (s *snakePart) getPositionAfterHorizontalUpTurn(partLength float32) (fyne.Position, fyne.Position) {

	switch s.direction {
	case leftDirection:
		return fyne.NewPos(s.part.Position2.X, s.part.Position1.Y-partLength), s.part.Position2
	case rightDirection:
		return fyne.NewPos(s.part.Position1.X, s.part.Position1.Y-partLength), s.part.Position1
	default:
		err := errors.New("Unexpected current direction passed when setting direction for snake figure head/body part [from horizontal to up]")
		log.WithFields(log.Fields{
			"err":               err.Error(),
			"current_direction": s.direction,
			"new_direction":     upDirection,
		}).Error("Error turning snake figure")
		return fyne.Position{}, fyne.Position{}
	}
}

func (s *snakePart) getPositionAfterHorizontalDownTurn(partLength float32) (fyne.Position, fyne.Position) {

	switch s.direction {
	case leftDirection:
		return s.part.Position2, fyne.NewPos(s.part.Position2.X, s.part.Position2.Y+partLength)
	case rightDirection:
		return s.part.Position1, fyne.NewPos(s.part.Position1.X, s.part.Position1.Y+partLength)
	default:
		err := errors.New("Unexpected current direction passed when setting direction for snake figure head/body part [from horizontal to down]")
		log.WithFields(log.Fields{
			"err":               err.Error(),
			"current_direction": s.direction,
			"new_direction":     downDirection,
		}).Error("Error turning snake figure")
		return fyne.Position{}, fyne.Position{}
	}

}
