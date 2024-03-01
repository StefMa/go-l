package api

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/binary"
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
	size := request.Form["size"]
	if len(size) == 0 {
		http.Error(response, "size is required", http.StatusBadRequest)
		return
	}
	sizeInt, err := strconv.Atoi(size[0])
	if err != nil {
		http.Error(response, "size must be an integer", http.StatusBadRequest)
		return
	}

	gameOfLife := gol.NewGameOfLife(gol.Size(sizeInt))
	gameStateInBase64 := newGameStateInBase64ForStart(gameOfLife)

	http.Redirect(response, request, "/next?state="+gameStateInBase64, http.StatusFound)
}

func newGameStateInBase64ForStart(gameOfLife *gol.GameOfLife) string {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, int32(gameOfLife.Size))
	for _, row := range gameOfLife.GameBoard {
		for _, cell := range row {
			if cell.State == gol.Life {
				binary.Write(buf, binary.LittleEndian, int32(cell.Point.X))
				binary.Write(buf, binary.LittleEndian, int32(cell.Point.Y))
			}
		}
	}
	compressed, err := compressStringForStart(buf.Bytes())
	if err != nil {
		panic("compressing failed")
	}
	base64Encoded := base64.StdEncoding.EncodeToString([]byte(compressed))
	urlEncoded := url.QueryEscape(base64Encoded)
	return urlEncoded
}

func compressStringForStart(input []byte) (string, error) {
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

	return string(in.Bytes()), nil
}
