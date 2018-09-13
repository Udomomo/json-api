package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/todos", TodoIndex)
	router.HandleFunc("/todos/{todoID}", TodoShow)
	log.Fatal(http.ListenAndServe(":8080", router))
}

//Index ハンドラ
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wellcome!")
}

//TodoIndex ハンドラ
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Todo Index!")
}

//TodoShow ハンドラ
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID := vars["todoID"]
	fmt.Fprintln(w, "Todo show:", todoID)
}
