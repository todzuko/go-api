package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/todzuko/go-api/utils"
)

func New() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/quests", utils.GetAllQuests).Methods("GET")
	router.HandleFunc("/quest/{id}", utils.GetQuest).Methods("GET")
	router.HandleFunc("/quest", utils.CreateQuest).Methods("POST")
	router.HandleFunc("/quest/{id}", utils.UpdateQuest).Methods("PUT")
	router.HandleFunc("/quest/{id}", utils.DeleteQuest).Methods("DELETE")

	return router
}
