package main

import (
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/app"
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/config"
)

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
