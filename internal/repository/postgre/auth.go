package postgre

import (
	"errors"
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/model"
	"gorm.io/gorm"
	"log"
)

type AuthorizationRepository struct {
	db *gorm.DB
}

func NewAuthorizationRepository(db *gorm.DB) *AuthorizationRepository {
	return &AuthorizationRepository{db: db}
}

func (r *AuthorizationRepository) GetUser(user model.AuthUser) (model.User, error) {
	var userResp model.User
	err := r.db.Where("username = ? AND password = ?", user.Username, user.Password).First(&userResp)
	if err.Error != nil {
		return model.User{}, err.Error
	}
	return userResp, nil
}

func (r *AuthorizationRepository) CreateUser(user model.User) (uint, error) {
	var checkUser model.User
	res := r.db.Where("username = ?", user.Username).First(&checkUser)
	if res.Error == nil {
		log.Println("fuck you")
		return 0, errors.New("such username already exist")
	}
	err := r.db.Create(&user)
	if err.Error != nil {
		return 0, err.Error
	}
	return user.ID, nil
}
