package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
)

const (
	minGridSize            = float32(500)
	snakeHeadWidth         = float32(15)
	snakeBodyWidth         = float32(10)
	gameOverText           = "Game Over!"
	gameOverTextSize       = float32(20)
	snakeBodyPartLength    = float32(33)
	snakeHeadLength        = float32(33)
	numOfStartingBodyParts = 3
	headPart               = "head"
	bodyPart               = "body"
	upDirection            = "Up"
	downDirection          = "Down"
	leftDirection          = "Left"
	rightDirection         = "Right"
)

var (
	snakeHeadColor    color.Color
	snakeBodyColor    color.Color
	gameOverTextColor color.Color
)
var (
	gameOverTextAlignment = fyne.TextAlignLeading
	gameOverTextStyle     = fyne.TextStyle{Bold: true}
	snakeSpeed            = float32(10)
)

var (
	directionKeys      = map[string]bool{upDirection: true, downDirection: true, leftDirection: true, rightDirection: true}
	oppositeDirections = map[string]string{upDirection: downDirection, downDirection: upDirection, leftDirection: rightDirection, rightDirection: leftDirection}
)
