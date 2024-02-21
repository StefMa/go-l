package api

import (
	_ "embed"
	"html/template"
	"net/http"
)

//go:embed templates/index.html
var indexHtml string

func IndexHandler(response http.ResponseWriter, request *http.Request) {
	tmpl, err := template.New("index").Parse(indexHtml)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(response, nil)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}
