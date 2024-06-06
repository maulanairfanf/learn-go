package handlers

import (
	"encoding/json"
	"myapi/db"
	"myapi/models"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Claims represents the JWT claims
type Claims struct {
    UserID uint `json:"user_id"`
    jwt.StandardClaims
}

// LoginRequest represents the request payload for the login endpoint
type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

// LoginResponse represents the response payload for the login endpoint
type LoginResponse struct {
	Token string `json:"token"`
}

// Login handles user authentication and issues a JWT token upon successful login
func Login(w http.ResponseWriter, r *http.Request) {
    var loginReq LoginRequest

    // Decode the JSON request body into the LoginRequest struct
    if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
        ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    // Check if the user exists
    var user models.User
    if err := db.DB.Where("username = ?", loginReq.Username).First(&user).Error; err != nil {
        ErrorResponse(w, http.StatusUnauthorized, "Invalid username or password")
        return
    }

    // Verify the password
    if !checkPasswordHash(loginReq.Password, user.Password) {
        ErrorResponse(w, http.StatusUnauthorized, "Invalid username or password")
        return
    }

    // Generate a JWT token
    token, err := generateJWTToken(user.ID)
    if err != nil {
        ErrorResponse(w, http.StatusInternalServerError, "Failed to generate token")
        return
    }

    // Send the token in the response
    SuccessResponse(w, LoginResponse{Token: token})
}

// checkPasswordHash verifies if a password matches its hashed version
func checkPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

// generateJWTToken generates a JWT token for the given user ID
func generateJWTToken(userID uint) (string, error) {
    // Create the token claims
    claims := Claims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
            IssuedAt:  time.Now().Unix(),
        },
    }

    // Generate JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    jwtSecret := []byte(os.Getenv("JWT_SECRET")) // Change this to your secret key

    signedToken, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }

    return signedToken, nil
}
