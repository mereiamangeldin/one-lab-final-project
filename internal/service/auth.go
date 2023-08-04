package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/model"
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/repository"
	"time"
)

const (
	salt       = "sgnrapong3g04jdfi"
	tokenTTL   = 12 * time.Hour
	signingKey = "sdgsgwg4gewsgsrhr"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId uint `json:"user_id"`
}

type AuthorizationService struct {
	Repo *repository.Repository
}

func (s *AuthorizationService) ParseToken(accessToken string) (uint, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil
}

func (s *AuthorizationService) GenerateToken(user model.AuthUser) (string, error) {
	user.Password = s.generatePassHash(user.Password)
	userResp, err := s.Repo.Auth.GetUser(user)
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: userResp.ID,
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthorizationService) CreateUser(user model.User) (uint, error) {
	user.Password = s.generatePassHash(user.Password)
	return s.Repo.Auth.CreateUser(user)
}

func (s *AuthorizationService) generatePassHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func NewAuthorizationService(repo *repository.Repository) *AuthorizationService {
	return &AuthorizationService{Repo: repo}
}

func (s *AuthorizationService) UpdatePassword(id uint, pass model.ChangePassword) error {
	pass.CurrentPassword = s.generatePassHash(pass.CurrentPassword)
	pass.NewPassword = s.generatePassHash(pass.NewPassword)
	return s.Repo.User.UpdatePassword(id, pass)
}
