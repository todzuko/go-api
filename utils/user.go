package utils

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/todzuko/go-api/models"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []models.User
	models.DB.Select("id", "name", "age", "email", "created_at", "updated_at").Find(&users)

	var userResponses []models.UserResponse
	for _, user := range users {
		userResponse := models.UserResponse{
			ID:      user.ID,
			Name:    user.Name,
			Age:     user.Age,
			Email:   user.Email,
			Created: user.CreatedAt,
			Updated: user.UpdatedAt,
		}
		userResponses = append(userResponses, userResponse)
	}
	json.NewEncoder(w).Encode(userResponses)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In GetUser")
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	id := mux.Vars(r)["id"]
	if err := models.DB.Where("id = ?", id).Select("id", "name", "age", "email", "created_at", "updated_at").First(&user).Error; err != nil {
		RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}
	userResponse := models.UserResponse{
		ID:      user.ID,
		Name:    user.Name,
		Age:     user.Age,
		Email:   user.Email,
		Created: user.CreatedAt,
		Updated: user.UpdatedAt,
	}
	json.NewEncoder(w).Encode(userResponse)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var input models.UserInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Validation Error")
		return
	}
	password := input.Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to save the data")
		return
	}

	user := &models.User{
		Name:     input.Name,
		Age:      input.Age,
		Email:    input.Email,
		Password: string(hashedPassword),
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

	var input models.UserInput

	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &input)

	validate = validator.New()
	err := validate.Struct(input)

	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Validation Error")
		return
	}
	password := input.Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to save the data")
		return
	}

	user.Name = input.Name
	user.Age = input.Age
	user.Email = input.Email
	user.Password = string(hashedPassword)

	models.DB.Save(&user)

	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var user models.User
	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}
	models.DB.Delete(&user)
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(user)

}
