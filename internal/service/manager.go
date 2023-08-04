package service

import (
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/config"
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/repository"
	"github.com/mereiamangeldin/One-lab-Homework-1/pkg/jwttoken"
)

type Manager struct {
	Repository repository.Repository
	Token      *jwttoken.JWTToken
	Config     *config.Config
}

func New(repository repository.Repository, token *jwttoken.JWTToken, config *config.Config) *Manager {
	return &Manager{
		Repository: repository,
		Token:      token,
		Config:     config,
	}
}
