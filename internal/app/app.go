package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/solumD/go-service-template/config"
	"github.com/solumD/go-service-template/internal/handler"
	"github.com/solumD/go-service-template/internal/repository/postgres"
	"github.com/solumD/go-service-template/internal/server"
	"github.com/solumD/go-service-template/internal/service"
	"github.com/solumD/go-service-template/pkg/logger"
	pg "github.com/solumD/go-service-template/pkg/postgres"
	"go.uber.org/zap"
)

const shutdownTimeout = 10 * time.Second

func InitAndRun(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	cfg := config.MustLoad()

	logger.Init(logger.GetCore(logger.GetAtomicLevel(cfg.LoggerLevel())))
	logger.Info("starting server")
	logger.Debug("debug messages are enabled")

	postgresConn := pg.New(cfg.PostgresDSN())
	defer postgresConn.Close()

	logger.Info("connected to postgres")

	repository := postgres.New(postgresConn)

	service := service.New(repository)

	handler := handler.New(service)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Route("/api/v1", func(r chi.Router) {
		r.Route("/entity", func(r chi.Router) {
			r.Post("/", handler.CreateEntity(ctx))
			r.Get("/{id}", handler.GetEntity(ctx))
		})
	})

	server := server.New(cfg.ServerAddr(), router)

	go func() {
		if err := server.Run(); err != nil {
			log.Fatalf("failed to run server: %v", err)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt

	logger.Info("shutting down server...")

	shutdownCtx, cancelShutdownCtx := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancelShutdownCtx()

	err := server.Shutdown(shutdownCtx)
	if err != nil {
		logger.Error("error while shutting down server", zap.Error(err))
	}
}
