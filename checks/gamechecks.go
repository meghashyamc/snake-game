package checks

func CheckGameState(gameStateMap gameState) string {

	if gameStateMap.isGameOver() {
		return gameOver
	}

	if gameStateMap.isFoodEaten() {

		return foodEaten
	}
	return ""

}
func (gs gameState) isGameOver() bool {

	if hasSnakeReachedExtremes(gs[snakeHeadX], gs[gridSize]) || hasSnakeReachedExtremes(gs[snakeHeadY], gs[gridSize]) {
		return true
	}

	return false

}

func (gs gameState) isFoodEaten() bool {

	rectangleLeftX, rectangleRightX, rectangleTopY, rectangleBottomY := gs.getEnclosingRectangleForCircle(gs[foodParticleCentreX], gs[foodParticleCentreY])

	return isSnakeHeadBetweenCoords(rectangleLeftX, rectangleRightX, gs[snakeHeadX]) && isSnakeHeadBetweenCoords(rectangleTopY, rectangleBottomY, gs[snakeHeadY])

}

func (gs gameState) getEnclosingRectangleForCircle(centreX, centreY float32) (float32, float32, float32, float32) {

	radius := gs[foodParticleDiameter] / 2
	return centreX - radius, centreX + radius, centreY - radius, centreY + radius
}

func isSnakeHeadBetweenCoords(coord1, coord2, snakeHeadCoord float32) bool {

	return coord1 <= snakeHeadCoord+clearence && coord2 >= snakeHeadCoord-clearence
}

func hasSnakeReachedExtremes(snakeHeadPosition, gridLength float32) bool {

	return snakeHeadPosition <= 0 || snakeHeadPosition >= gridLength

}
