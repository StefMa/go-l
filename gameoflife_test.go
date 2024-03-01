package gol

import (
	"testing"
)

func TestNewGameOfLife(t *testing.T) {
	game := NewGameOfLife(3)

	if game.Size != 3 {
		t.Errorf("Expected game to have a width of 3")
	}

	if len(game.GameBoard) != 3 {
		t.Errorf("Expected game to have 3 rows")
	}

	if len(game.GameBoard[0]) != 3 {
		t.Errorf("Expected first row to have 3 columns")
	}

	if len(game.GameBoard[1]) != 3 {
		t.Errorf("Expected second row to have 3 columns")
	}

	if len(game.GameBoard[2]) != 3 {
		t.Errorf("Expected thrid row to have 3 columns")
	}

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			cell := game.GameBoard[y][x]
			if cell.State != Dead && cell.State != Life {
				t.Errorf("Expected cell to be dead or life")
			}
		}
	}
}

func TestGameOfLife_GetNeighbors(t *testing.T) {
	game := NewGameOfLife(3)

	neighbors := game.getNeighbors(game.GameBoard[1][1])

	if len(neighbors) != 8 {
		t.Errorf("Expected cell to have 8 neighbors")
	}

	expected := []Point{
		{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2},
		{X: 1, Y: 0} /* /Empty/ */, {X: 1, Y: 2},
		{X: 2, Y: 0}, {X: 2, Y: 1}, {X: 2, Y: 2},
	}
	for _, neighbor := range neighbors {
		if !containsPoint(expected, neighbor.Point) {
			t.Errorf("Expected %v to be part of the neightbor slice", neighbor)
		}
	}
}

func TestNewGameOfLifeWithGenerator(t *testing.T) {
	game := NewGameOfLifeWithGenerator(3, func(x int, y int) Cell {
		return Cell{
			State: Life,
		}
	})

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			cell := game.GameBoard[y][x]
			if cell.State != Life {
				t.Errorf("Expected cell to be life")
			}
		}
	}
}

func TestGameOfLife_Next_3By3(t *testing.T) {
	// Init game with the following state:
	/**
	 * 0 0 0
	 * 0 1 1
	 * 0 0 1
	 */
	game := NewGameOfLifeWithGenerator(3, func(x int, y int) Cell {
		state := Dead
		if x == 1 && y == 1 || x == 2 && y == 1 ||
			x == 2 && y == 2 {
			state = Life
		}
		return Cell{
			Point: Point{X: x, Y: y},
			State: state,
		}
	})

	game.Next()

	// Expected state:
	/**
	 * 0 0 0
	 * 0 1 1
	 * 0 1 1
	 */
	expected := [][]CellState{
		{Dead, Dead, Dead},
		{Dead, Life, Life},
		{Dead, Life, Life},
	}
	for _, cell := range game.GameBoard[0] {
		if cell.State != expected[0][cell.Point.X] {
			t.Errorf("Expected first row to be %v but is %v", expected[0], game.GameBoard[0])
		}
	}
	for _, cell := range game.GameBoard[1] {
		if cell.State != expected[1][cell.Point.X] {
			t.Errorf("Expected second row to be %v but is %v", expected[1], game.GameBoard[1])
		}
	}
	for _, cell := range game.GameBoard[2] {
		if cell.State != expected[2][cell.Point.X] {
			t.Errorf("Expected third row to be %v but is %v", expected[2], game.GameBoard[2])
		}
	}
}

func TestGameOfLife_Next_TwoTimes_3By3(t *testing.T) {
	// Init game with the following state:
	/**
	 * 1 0 1
	 * 0 1 0
	 * 0 0 0
	 */
	game := NewGameOfLifeWithGenerator(3, func(x int, y int) Cell {
		state := Dead
		if x == 0 && y == 0 || x == 2 && y == 0 ||
			x == 1 && y == 1 {
			state = Life
		}
		return Cell{
			Point: Point{X: x, Y: y},
			State: state,
		}
	})

	game.Next()

	// Expected state:
	/**
	 * 0 1 0
	 * 0 1 0
	 * 0 0 0
	 */
	expected := [][]CellState{
		{Dead, Life, Dead},
		{Dead, Life, Dead},
		{Dead, Dead, Dead},
	}
	for _, cell := range game.GameBoard[0] {
		if cell.State != expected[0][cell.Point.X] {
			t.Errorf("Expected first row to be %v but is %v", expected[0], game.GameBoard[0])
		}
	}
	for _, cell := range game.GameBoard[1] {
		if cell.State != expected[1][cell.Point.X] {
			t.Errorf("Expected second row to be %v but is %v", expected[1], game.GameBoard[1])
		}
	}
	for _, cell := range game.GameBoard[2] {
		if cell.State != expected[2][cell.Point.X] {
			t.Errorf("Expected third row to be %v but is %v", expected[2], game.GameBoard[2])
		}
	}

	game.Next()

	// Expected state:
	/**
	 * 0 0 0
	 * 0 0 0
	 * 0 0 0
	 */
	expected = [][]CellState{
		{Dead, Dead, Dead},
		{Dead, Dead, Dead},
		{Dead, Dead, Dead},
	}
	for _, cell := range game.GameBoard[0] {
		if cell.State != expected[0][cell.Point.X] {
			t.Errorf("Expected first row to be %v but is %v", expected[0], game.GameBoard[0])
		}
	}
	for _, cell := range game.GameBoard[1] {
		if cell.State != expected[1][cell.Point.X] {
			t.Errorf("Expected second row to be %v but is %v", expected[1], game.GameBoard[1])
		}
	}
	for _, cell := range game.GameBoard[2] {
		if cell.State != expected[2][cell.Point.X] {
			t.Errorf("Expected third row to be %v but is %v", expected[2], game.GameBoard[2])
		}
	}
}

func TestGameOfLife_Next_5By5(t *testing.T) {
	// Init game with the following state:
	/**
	 * 0 1 0 1 0
	 * 0 1 0 1 0
	 * 1 0 1 0 0
	 * 0 1 0 0 1
	 * 0 0 1 0 0
	 */
	game := NewGameOfLifeWithGenerator(5, func(x int, y int) Cell {
		state := Dead
		if x == 1 && y == 0 || x == 3 && y == 0 ||
			x == 1 && y == 1 || x == 3 && y == 1 ||
			x == 0 && y == 2 || x == 2 && y == 2 ||
			x == 1 && y == 3 || x == 4 && y == 3 ||
			x == 2 && y == 4 {
			state = Life
		}
		return Cell{
			Point: Point{X: x, Y: y},
			State: state,
		}
	})

	game.Next()

	// Expected state:
	/**
	 * 0 0 0 0 0
	 * 1 1 0 1 0
	 * 1 0 1 1 0
	 * 0 1 1 1 0
	 * 0 0 0 0 0
	 */
	expected := [][]CellState{
		{Dead, Dead, Dead, Dead, Dead},
		{Life, Life, Dead, Life, Dead},
		{Life, Dead, Life, Life, Dead},
		{Dead, Life, Life, Life, Dead},
		{Dead, Dead, Dead, Dead, Dead},
	}
	for _, cell := range game.GameBoard[0] {
		if cell.State != expected[0][cell.Point.X] {
			t.Errorf("Expected first row to be %v but is %v", expected[0], game.GameBoard[0])
		}
	}
	for _, cell := range game.GameBoard[1] {
		if cell.State != expected[1][cell.Point.X] {
			t.Errorf("Expected second row to be %v but is %v", expected[1], game.GameBoard[1])
		}
	}
	for _, cell := range game.GameBoard[2] {
		if cell.State != expected[2][cell.Point.X] {
			t.Errorf("Expected third row to be %v but is %v", expected[2], game.GameBoard[2])
		}
	}
	for _, cell := range game.GameBoard[3] {
		if cell.State != expected[3][cell.Point.X] {
			t.Errorf("Expected fourth row to be %v but is %v", expected[3], game.GameBoard[3])
		}
	}
	for _, cell := range game.GameBoard[4] {
		if cell.State != expected[4][cell.Point.X] {
			t.Errorf("Expected fifth row to be %v but is %v", expected[4], game.GameBoard[4])
		}
	}
}

func containsPoint(points []Point, point Point) bool {
	for _, c := range points {
		if c == point {
			return true
		}
	}
	return false
}
