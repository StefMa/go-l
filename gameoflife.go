package gol

import (
	"math/rand"
	"time"
)

type GameOfLife struct {
	Size      Size
	GameBoard [][]Cell
}

type Size int

type CellGenerator func(x int, y int) Cell

func NewGameOfLife(size Size) *GameOfLife {
	return NewGameOfLifeWithGenerator(size, defaultCellGenerator)
}

func NewGameOfLifeWithGenerator(size Size, cellGenerator CellGenerator) *GameOfLife {
	var gameBoard [][]Cell
	for y := 0; y < int(size); y++ {
		var rowCells []Cell
		for x := 0; x < int(size); x++ {
			cell := cellGenerator(x, y)
			rowCells = append(rowCells, cell)
		}
		gameBoard = append(gameBoard, rowCells)
	}

	return &GameOfLife{
		Size:      size,
		GameBoard: gameBoard,
	}
}

func (gol *GameOfLife) Next() {
	var newGameBoard [][]Cell
	for y := 0; y < int(gol.Size); y++ {
		var rowCells []Cell
		for x := 0; x < int(gol.Size); x++ {
			cell := gol.GameBoard[y][x]
			neighbors := gol.getNeighbors(cell)
			newState := cell.State
			if cell.IsUnderpopulated(neighbors) {
				newState = Dead
			} else if cell.EvolveToNextGeneration(neighbors) {
				newState = Life
			} else if cell.IsOverpopulated(neighbors) {
				newState = Dead
			} else if cell.ShouldBecomeAlive(neighbors) {
				newState = Life
			}
			newCell := Cell{
				Point: cell.Point,
				State: newState,
			}
			rowCells = append(rowCells, newCell)
		}
		newGameBoard = append(newGameBoard, rowCells)
	}
	gol.GameBoard = newGameBoard
}

func defaultCellGenerator(x int, y int) Cell {
	return Cell{
		Point: Point{
			X: x,
			Y: y,
		},
		State: createRandomCellState(),
	}
}

func createRandomCellState() CellState {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := r.Intn(2)
	if randomNumber == 0 {
		return Dead
	} else {
		return Life
	}
}

func (gol *GameOfLife) getNeighbors(cell Cell) []Cell {
	cellX := cell.Point.X
	cellY := cell.Point.Y
	var neighbors []Cell
	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			neighborX := cellX + y
			neighborY := cellY + x
			if gol.isPointOutOfBounds(neighborX, neighborY) {
				continue
			}
			possibleNeighbor := gol.GameBoard[neighborY][neighborX]
			if isCurrentCell(cell, possibleNeighbor) {
				continue
			}
			neighbors = append(neighbors, possibleNeighbor)
		}
	}
	return neighbors
}

func (gol *GameOfLife) isPointOutOfBounds(neighborX int, neighborY int) bool {
	return neighborX < 0 || neighborX >= int(gol.Size) || neighborY < 0 || neighborY >= int(gol.Size)
}

func isCurrentCell(current Cell, possibleNeighbor Cell) bool {
	return current == possibleNeighbor
}
