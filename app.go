package main

import (
	"encoding/json"
	"database/sql"
	"net/http"
	"strconv"
	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
)


type App struct {
	Router *mux.Router,
	DB *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {}

func (a *App) Run(address string){}

func (a *App) getProduct(w http.ResponseWriter, r http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid Product ID")
		return
	}
	p := product{ID: id}
	if err := p.getProduct(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Product not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOk, p)
}

func (a *App) getProducts(w http.ResponseWriter, r http.Request) {
	count, _ := strconv.Atoi(r.FormValue["count"])
	start, _ := strconv.Atoi(r.FormValue["start"])
	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0 //this makes sense I guess
	}

	products, err := getProducts(a.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, products)
}

func (a *App) createProduct (w http.ResponseWriter, r http.Request) {
	var p product
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, 400, "Invalid request")
	}

}
