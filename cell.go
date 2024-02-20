package gol

type Cell struct {
	Point Point
	State CellState
}

type CellState int

const (
	Dead CellState = iota
	Life
)

func (cell Cell) isLife() bool {
	return cell.State == Life
}

func (cell Cell) isDead() bool {
	return cell.State == Dead
}

// IsUnderpopulated return value is evaluated by:
// Any live cell with fewer than two live neighbors dies, as if by underpopulation.
func (cell Cell) IsUnderpopulated(neighbors []Cell) bool {
	if cell.isDead() {
		return false
	}

	neighborsLifeCount := countLivingNeighbors(neighbors)

	return neighborsLifeCount < 2
}

// EvolveToNextGeneration return value is evaluated by:
// Any live cell with two or three live neighbors lives on to the next generation.
func (cell Cell) EvolveToNextGeneration(neighbors []Cell) bool {
	if cell.isDead() {
		return false
	}

	neighborsLifeCount := countLivingNeighbors(neighbors)

	return neighborsLifeCount == 2 || neighborsLifeCount == 3
}

// IsOverpopulated return value is evaluated by:
// Any live cell with more than three live neighbors dies, as if by overpopulation.
func (cell Cell) IsOverpopulated(neighbors []Cell) bool {
	if cell.isDead() {
		return false
	}

	neighborsLifeCount := countLivingNeighbors(neighbors)

	return neighborsLifeCount > 3
}

// ShouldBecomeAlive return value is evaluated by:
// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.
func (cell Cell) ShouldBecomeAlive(neighbors []Cell) bool {
	if cell.isLife() {
		return false
	}

	neighborsLifeCount := countLivingNeighbors(neighbors)

	return neighborsLifeCount == 3
}

func countLivingNeighbors(neighbors []Cell) int {
	lifeCount := 0
	for _, neighbor := range neighbors {
		predictionFound := neighbor.isLife()
		if predictionFound {
			lifeCount += 1
		}
	}
	return lifeCount
}
