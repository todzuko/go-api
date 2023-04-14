package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID      uint      `json:"id"`
	Name    string    `json:"name"`
	Age     int       `json:"age"`
	Email   string    `json:"email"`
	Created time.Time `json:"created_at"`
	Updated time.Time `json:"updated_at"`
}

type UserInput struct {
	Name     string `json:"name" validate:"required"`
	Age      int    `json:"age" validate:"required,gt=0"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
