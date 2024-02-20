package gol

import (
	"math/rand"
	"time"
)

type GameOfLife struct {
	Width     Width
	Height    Height
	GameBoard [][]Cell
}

type Width int
type Height int

type CellGenerator func(x int, y int) Cell

func NewGameOfLife(width Width, height Height) *GameOfLife {
	return NewGameOfLifeWithGenerator(width, height, defaultCellGenerator)
}

func NewGameOfLifeWithGenerator(width Width, height Height, cellGenerator CellGenerator) *GameOfLife {
	var gameBoard [][]Cell
	for y := 0; y < int(height); y++ {
		var rowCells []Cell
		for x := 0; x < int(width); x++ {
			cell := cellGenerator(x, y)
			rowCells = append(rowCells, cell)
		}
		gameBoard = append(gameBoard, rowCells)
	}

	return &GameOfLife{
		Width:     width,
		Height:    height,
		GameBoard: gameBoard,
	}
}

func (gol *GameOfLife) Next() {
	var newGameBoard [][]Cell
	for y := 0; y < int(gol.Height); y++ {
		var rowCells []Cell
		for x := 0; x < int(gol.Width); x++ {
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
	return neighborX < 0 || neighborX >= int(gol.Width) || neighborY < 0 || neighborY >= int(gol.Height)
}

func isCurrentCell(current Cell, possibleNeighbor Cell) bool {
	return current == possibleNeighbor
}
