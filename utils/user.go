package utils

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/todzuko/go-api/models"
	"io/ioutil"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	models.DB.Find(&users)

	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	id := mux.Vars(r)["id"]
	if err := models.DB.Where("id = ?", id).First(user).Error; err != nil {
		RespondWithError(w, http.StatusNotFound, "Quest not found")
		return
	}

	json.NewEncoder(w).Encode(user)
}

type UserInput struct {
	Name     string `json:"name" validate:"required"`
	Age      int    `json:"age" validate:"required,gt=0"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input UserInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Validation Error")
		return
	}

	user := &models.User{
		Name:     input.Name,
		Age:      input.Age,
		Email:    input.Email,
		Password: input.Password,
	}

	models.DB.Create(user)

	json.NewEncoder(w).Encode(user)
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var user models.User

	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}

	var input UserInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Validation Error")
		return
	}

	user.Name = input.Name
	user.Age = input.Age
	user.Email = input.Email
	user.Password = input.Password

	models.DB.Save(&user)

	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var user models.User
	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		RespondWithError(w, http.StatusNotFound, "Quest not found")
		return
	}
	models.DB.Delete(&user)
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(user)

}
