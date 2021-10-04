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
	headPart               = "head"
	bodyPart               = "body"
	gameOverTextSize       = float32(20)
	snakeBodyPartLength    = float32(33)
	snakeHeadLength        = float32(33)
	foodDiameter           = float32(10)
	numOfStartingBodyParts = 3
)
const (
	upDirection    = "Up"
	downDirection  = "Down"
	leftDirection  = "Left"
	rightDirection = "Right"
)

const ()

var (
	snakeHeadColor    color.Color
	snakeBodyColor    color.Color
	gameOverTextColor color.Color
	foodParticleColor = color.RGBA{255, 255, 0, 0}
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
