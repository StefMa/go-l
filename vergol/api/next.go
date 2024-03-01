package api

import (
	"bytes"
	"compress/zlib"
	_ "embed"
	"encoding/base64"
	"encoding/binary"
	"html/template"
	"io"
	"net/http"

	gol "github.com/stefma/go-l"
)

//go:embed templates/next.html
var nextHtml string

func NextHandler(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	state := request.Form["state"]
	if len(state) == 0 {
		http.Error(response, "state is required", http.StatusBadRequest)
		return
	}
	stateDecoded, err := base64.StdEncoding.DecodeString(state[0])
	if err != nil {
		http.Error(response, "state must be base64 encoded", http.StatusBadRequest)
		return
	}
	stateDecoded, err = decompressStringForNext(stateDecoded)
	if err != nil {
		http.Error(response, "state must be a valid game board", http.StatusBadRequest)
		return
	}
	gameBoard, err := bytesToGameBoard(stateDecoded)
	if err != nil {
		http.Error(response, "can't convert bytes to game board", http.StatusBadRequest)
		return
	}
	gameOfLife := gol.NewGameOfLifeWithGenerator(
		gol.Size(len(gameBoard[0])),
		func(x, y int) gol.Cell { return gameBoard[y][x] },
	)
	currentGameBoard := gameBoard

	gameOfLife.Next()

	nextGameStateInBase64 := newGameStateInBase64ForNext(gameOfLife)

	tmpl, err := template.New("next").Parse(nextHtml)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	nextHtmlData := NextHtmlData{
		NextGameStateInBase64: nextGameStateInBase64,
		CurrentGameBoard:      gameBoardToStringBoard(currentGameBoard),
	}
	err = tmpl.Execute(response, nextHtmlData)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

func gameBoardToStringBoard(gameBoard [][]gol.Cell) [][]string {
	stringBoard := make([][]string, len(gameBoard))
	for y, row := range gameBoard {
		stringBoard[y] = make([]string, len(row))
		for x, cell := range row {
			stringBoard[y][x] = htmlCell(cell).String()
		}
	}
	return stringBoard
}

type htmlCell gol.Cell

func (htmlCell htmlCell) String() string {
	if htmlCell.State == gol.Life {
		return "👪"
	}
	return "💀"
}

type NextHtmlData struct {
	NextGameStateInBase64 string
	CurrentGameBoard      [][]string
}

func newGameStateInBase64ForNext(gameOfLife *gol.GameOfLife) string {
	buf := new(bytes.Buffer)
	for _, row := range gameOfLife.GameBoard {
		for _, cell := range row {
			binary.Write(buf, binary.LittleEndian, int32(cell.Point.X))
			binary.Write(buf, binary.LittleEndian, int32(cell.Point.Y))
			binary.Write(buf, binary.LittleEndian, int8(cell.State))
		}
	}
	compressedJson, err := compressStringForNext(buf.Bytes())
	if err != nil {
		panic("compressing failed")
	}
	return base64.StdEncoding.EncodeToString([]byte(compressedJson))
}

func compressStringForNext(input []byte) (string, error) {
	var in bytes.Buffer

	w := zlib.NewWriter(&in)
	_, err := w.Write(input)
	if err != nil {
		return "", err
	}
	err = w.Close()
	if err != nil {
		return "", err
	}

	x := string(in.Bytes())
	return x, nil
}

func decompressStringForNext(input []byte) ([]byte, error) {
	b := bytes.NewReader(input)

	r, err := zlib.NewReader(b)
	if err != nil {
		return []byte{}, err
	}
	defer r.Close()

	decompressed, err := io.ReadAll(r)
	if err != nil {
		return []byte{}, err
	}

	return decompressed, nil
}

func bytesToGameBoard(data []byte) ([][]gol.Cell, error) {
	buf := bytes.NewReader(data)
	var gameBoard [][]gol.Cell
	for {
		var x, y int32
		var state int8
		err := binary.Read(buf, binary.LittleEndian, &x)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		err = binary.Read(buf, binary.LittleEndian, &y)
		if err != nil {
			return nil, err
		}
		err = binary.Read(buf, binary.LittleEndian, &state)
		if err != nil {
			return nil, err
		}
		cell := gol.Cell{Point: gol.Point{X: int(x), Y: int(y)}, State: gol.CellState(state)}
		if int(y) >= len(gameBoard) {
			gameBoard = append(gameBoard, []gol.Cell{})
		}
		gameBoard[y] = append(gameBoard[y], cell)
	}
	return gameBoard, nil
}
