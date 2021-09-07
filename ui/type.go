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

var directionKeys = map[string]bool{"Up": true, "Down": true, "Left": true, "Right": true}
