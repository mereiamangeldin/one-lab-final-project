package http

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mereiamangeldin/One-lab-Homework-1/internal/config"
	"github.com/mereiamangeldin/One-lab-Homework-1/pkg/transport/http/handler"
	"log"
	"net/http"
	"time"
)

type Server struct {
	cfg        *config.Config
	httpServer *http.Server
	handler    *handler.Manager
	router     *echo.Echo
}

func NewServer(cfg *config.Config, handler *handler.Manager) *Server {
	return &Server{cfg: cfg, handler: handler}
}

func (s *Server) Run(ctx context.Context) error {
	s.router = s.BuildEngine()
	s.InitRoutes()
	go func() {
		if err := s.router.Start(fmt.Sprintf(":%d", s.cfg.Port)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()
	<-ctx.Done()

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err := s.router.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}
	log.Print("server exited properly")
	return nil
}

func (s *Server) BuildEngine() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	return e
}
