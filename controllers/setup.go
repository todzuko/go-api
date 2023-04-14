package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/todzuko/go-api/utils"
)

func New() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/quests", utils.GetAllQuests).Methods("GET")
	router.HandleFunc("/quests/{id}", utils.GetQuest).Methods("GET")
	router.HandleFunc("/quests", utils.CreateQuest).Methods("POST")
	router.HandleFunc("/quests/{id}", utils.UpdateQuest).Methods("PUT")
	router.HandleFunc("/quests/{id}", utils.DeleteQuest).Methods("DELETE")

	router.HandleFunc("/users", utils.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", utils.GetUser).Methods("GET")
	router.HandleFunc("/users", utils.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", utils.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", utils.DeleteUser).Methods("DELETE")

	return router
}
