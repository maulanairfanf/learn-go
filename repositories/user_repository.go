package repositories

import (
	"myapi/db"
	"myapi/models"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct{}

func (r *UserRepository) FindAll() ([]models.User, error) {
	var user []models.User
	err := db.DB.Find(&user).Error
	return user, err
}

func (r *UserRepository) FindByID(id int) (*models.User, error) {
	var user models.User
	err := db.DB.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) Create(user *models.User) error {
	return db.DB.Create(user).Error
}

func (r *UserRepository) Save(user *models.User) error {
	return db.DB.Save(user).Error
}

func (r *UserRepository) Delete(user *models.User) error {
	return db.DB.Delete(user).Error
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := db.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *UserRepository) MatchPassword(password, hash string) bool {
	return checkPasswordHash(password, hash)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
