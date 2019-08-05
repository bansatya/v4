package handler

import (
	"fmt"
	"net/http"
)

func Pong(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World!")
	respondJSON(w, http.StatusOK, `name:satya`)
}
func GetAccount(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, nil)
}
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusCreated, nil)
}

func CreateContract(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusCreated, nil)
}

func Execute(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusCreated, nil)
}

func Call(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, nil)
}
