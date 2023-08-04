package main

import (
	"context"
	"fmt"
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/config"
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/repository"
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/service"
	"github.com/mereiamangeldin/One-lab-Homework-1/pkg/transport/http"
	"github.com/mereiamangeldin/One-lab-Homework-1/pkg/transport/http/handler"
	"log"
	"os"
	"os/signal"
)

func main() {
	log.Fatalln(fmt.Sprintf("Service shut down:%s", run()))
}

func run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	gracefullyShutdown(cancel)
	conf, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}
	repo, err := repository.New(conf)
	if err != nil {
		log.Fatal(err.Error())
	}
	svc, err := service.NewManager(repo)
	if err != nil {
		log.Fatal(err.Error())
	}
	h := handler.NewManager(conf, svc)
	srv := http.NewServer(conf, h)
	return srv.Run(ctx)
}

func gracefullyShutdown(c context.CancelFunc) {
	osC := make(chan os.Signal, 1)
	signal.Notify(osC, os.Interrupt)
	go func() {
		log.Print(<-osC)
		c()
	}()
}
