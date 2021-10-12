package checks

const (
	clearence          = float32(10)
	pointsPerIncrement = float32(10)
)
const (
	nextTick         = "next_tick"
	directionChanged = "direction_changed"
	gameOver         = "game_over"
	foodEaten        = "food_eaten"
)

const (
	snakeHeadX           = "snake_head_x"
	snakeHeadY           = "snake_head_y"
	foodParticleCentreX  = "food_particle_centre_x"
	foodParticleCentreY  = "food_particle_centre_y"
	foodParticleDiameter = "food_particle_diameter"
	gridSize             = "grid_size"
	gameScore            = "game_score"
)

type gameState map[string]float32
