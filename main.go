package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Tab struct {
	Title string `json:"Name"`
	Id    int    `json:"Id"`
}

type TabInfo struct {
}

func tabHandler(writer http.ResponseWriter, request *http.Request) {
	tab := Tab{Title: "Ventura Highway", Id: 1}

	json.NewEncoder(writer).Encode(tab)
}

func main() {
	http.HandleFunc("/tabs", Handle(tabHandler))
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
