package server

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"mezink/src/business/domain"
	"mezink/src/business/usecase"
	"mezink/stdlib/db"
	logger "mezink/stdlib/log"
	"mezink/stdlib/middleware"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

type Server interface {
	Serve()
	EnableGracefulShutdown()
}

type Service struct {
	Router *mux.Router
	db     *sql.DB
}

func Init() Server {
	sqlDBm, err := db.Init()
	if err != nil {
		logger.FatalContext(context.Background(), err.Error())
	}

	dom := domain.Init(sqlDBm)
	uc := usecase.Init(dom)
	h := NewHandler(uc, sqlDBm)

	return &Service{
		Router: h.CreateRouter(),
		db:     sqlDBm,
	}
}

func (s *Service) Serve() {
	ctx := context.Background()
	s.EnableGracefulShutdown()
	middlewares := alice.New(middleware.LoggingHandler)
	logger.InfoContext(ctx, "service is running")
	logger.InfoContext(ctx, "service: rest Server mounted at [::]:8080")
	err := http.ListenAndServe(":8080", middlewares.Then(s.Router))
	if err != nil {
		logger.FatalContext(ctx, err.Error())
	}
}

func (s *Service) EnableGracefulShutdown() {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go s.handleShutdown(signalChannel)
}

func (s *Service) handleShutdown(ch chan os.Signal) {
	<-ch
	defer os.Exit(0)
	ctx := context.Background()
	duration := time.Duration(1 * time.Second)
	logger.InfoContext(ctx, "service: Signal termination received. Waiting %v seconds to shutdown.", duration.Seconds())
	IsShuttingDown = true
	time.Sleep(duration)
	s.db.Close()
	logger.InfoContext(ctx, "service: Cleaning up resources...\n")
	logger.InfoContext(ctx, "service: Bye\n")
}
