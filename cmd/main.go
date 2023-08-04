package main

import (
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/app"
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/config"
)

// @title           SHOP
// @version         0.0.1
// @description     API for SHOP application

// @contact.name   Merei
// @contact.email  me_amangeldin@kbtu.kz

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	cfg, err := config.InitConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	err = app.Run(cfg)
	if err != nil {
		panic(err)
	}
}
