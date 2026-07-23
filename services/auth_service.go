package services

import (
	"myapi/models"
	"myapi/repositories"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService() *AuthService {
	return &AuthService{
		userRepo: &repositories.UserRepository{},
	}
}

func (s *AuthService) Register(req models.User) (*models.User, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Username: req.Username,
		Password: string(hashPassword),
		Email:    req.Email,
	}

	if err := s.userRepo.Create(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *AuthService) Login(username, password string) (*models.User, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	if !s.userRepo.MatchPassword(password, user.Password) {
		return nil, ErrInvalidInput
	}

	return user, nil
}

func (s *AuthService) GenerateToken(userID uint) (string, error) {
	claims := models.Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
