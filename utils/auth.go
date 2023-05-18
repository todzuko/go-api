package utils

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/todzuko/go-api/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Credentials struct {
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Claims struct {
	Username string `json:"username"`
	User     uint   `json:"user"`
	jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Bad request")
		return
	}
	var user models.User
	if err := models.DB.Where("email = ?", creds.Email).Select("password").First(&user).Error; err != nil {
		RespondWithError(w, http.StatusNotFound, "User not found")
		return
	}
	hashedPassword := []byte(user.Password)
	password := []byte(creds.Password)
	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		RespondWithError(w, http.StatusUnauthorized, "Wrong password or email")
		return
	}
	duration, err := strconv.Atoi(os.Getenv("JWT_DURATION"))
	expirationTime := time.Now().Add(time.Duration(duration) * time.Second)
	if expirationTime == time.Now() {
		RespondWithError(w, http.StatusUnauthorized, "Wrong password or email")
		return
	}

	token, err := generateJWT(user.ID, user.Name, expirationTime)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Jwt error")
	}

	json.NewEncoder(w).Encode(token)
}

func generateJWT(userId uint, username string, expirationTime time.Time) (string, error) {
	claims := &Claims{
		Username: username,
		User:     userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := []byte(os.Getenv("API_SECRET"))
	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func Logout(w http.ResponseWriter, r *http.Request) {

}

func RefreshToken(w http.ResponseWriter, r *http.Request) {

}
