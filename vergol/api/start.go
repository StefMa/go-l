package api

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	gol "github.com/stefma/go-l"
)

func StartHandler(response http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	rows := request.Form["rows"]
	cols := request.Form["cols"]
	if len(rows) == 0 || len(cols) == 0 {
		http.Error(response, "rows and cols are required", http.StatusBadRequest)
		return
	}
	rowsInt, err := strconv.Atoi(rows[0])
	if err != nil {
		http.Error(response, "rows must be an integer", http.StatusBadRequest)
		return
	}
	colsInt, err := strconv.Atoi(cols[0])
	if err != nil {
		http.Error(response, "cols must be an integer", http.StatusBadRequest)
		return
	}

	gameOfLife := gol.NewGameOfLife(gol.Width(rowsInt), gol.Height(colsInt))
	gameStateInBase64 := newGameStateInBase64ForStart(gameOfLife)

	http.Redirect(response, request, "/next?state="+gameStateInBase64, http.StatusFound)
}

func newGameStateInBase64ForStart(gameOfLife *gol.GameOfLife) string {
	gameBoardJson, _ := json.Marshal(gameOfLife.GameBoard)
	compressedJson, err := compressStringForStart(string(gameBoardJson))
	if err != nil {
		return string(gameBoardJson)
	}
	base64Encoded := base64.StdEncoding.EncodeToString([]byte(compressedJson))
	urlEncoded := url.QueryEscape(base64Encoded)
	return urlEncoded
}

func compressStringForStart(input string) (string, error) {
	var in bytes.Buffer
	b := []byte(input)

	w := zlib.NewWriter(&in)
	_, err := w.Write(b)
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
