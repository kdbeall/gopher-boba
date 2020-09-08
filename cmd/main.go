package main

import (
	"../pkg/boba"
	"database/sql"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

func main() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	boba.Init(db)
	router := mux.NewRouter()
	router.HandleFunc("/boba", boba.GetBobas).Methods("GET")
	router.HandleFunc("/boba/{id}", boba.GetBoba).Methods("GET")
	router.HandleFunc("/boba", boba.CreateBoba).Methods("POST")
	router.HandleFunc("/boba/{id}", boba.UpdateBoba).Methods("PUT")
	router.HandleFunc("/boba/{id}", boba.DeleteBoba).Methods("DELETE")
	http.ListenAndServe(":8000", router)
}
