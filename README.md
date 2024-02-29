# go-l

A Go module implementation of [Conway's Game of Life](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life).

A demo can be found at **https://vergol.vercel.app**.

## How to use

**1.** First you need to create a new instance of the `GameOfLife` structure and put a `Height` and a `Width` into it to describe the dimensions of the game:
```go
gol "github.com/stefma/go-l"

func main() {
  gameOfLife := gol.NewGameOfLife(gol.Width(10), gol.Height(10))
}
```

**1.1** Additionally, you can use a custom `generator` function to create the game `Cell`s yourself:
```go
gameOfLife := gol.NewGameOfLifeWithGenerator(
		gol.Width(10),
		gol.Height(10),
		func(x, y int) gol.Cell { /* Custom logic here */ },
)
```

**2.0** Get the current game board by calling `gameOfLife.GameBoard` and render it somewhere:
```go
var gameBoard [][]Cell = gameOfLife.GameBoard
```

**3.0** To update the `Cell`s to the next evolution call `Next()`:
```go
gameOfLife.Next()
```

**4.0** Repeat steps 2 and 3 🙃

## Example: *Vergol*

One consumer of this module is [Vergol](vergol/), also part of this repository.
