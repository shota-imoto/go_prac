package main

import (
	"encoding/json"
	"log"
	"net/http"
	"shota-imoto/go_prac/handler"
)

type Tab struct {
	Title string `json:"Name"`
	Id    int    `json:"Id"`
}

type TabInfo struct {
}

func tabHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	tab := Tab{Title: "Ventura Highway", Id: 1}

	json.NewEncoder(writer).Encode(tab)
}

func main() {
	http.HandleFunc("/tabs", handler.Handle(tabHandler))
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
