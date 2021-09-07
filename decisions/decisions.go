package decisions

import (
	"fyne.io/fyne/v2"
)

func IsGameOver(snakeHeadPosition *fyne.Position, gridSize *fyne.Size) bool {

	if hasSnakeReachedExtremes(snakeHeadPosition.X, gridSize.Width) || hasSnakeReachedExtremes(snakeHeadPosition.Y, gridSize.Height) {
		return true
	}

	return false

}

func hasSnakeReachedExtremes(snakeHeadPosition, gridLength float32) bool {

	return snakeHeadPosition <= 0 || snakeHeadPosition >= gridLength

}
