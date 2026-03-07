package app

import (
	"context"
	lg "log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/solumD/go-service-template/config"
	v1 "github.com/solumD/go-service-template/internal/handler/v1"
	"github.com/solumD/go-service-template/internal/repository/postgres"
	"github.com/solumD/go-service-template/internal/usecase"
	httpserver "github.com/solumD/go-service-template/pkg/http_server"
	"github.com/solumD/go-service-template/pkg/logger"
	pg "github.com/solumD/go-service-template/pkg/postgres"
)

const shutdownTimeout = 10 * time.Second

func InitAndRun(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// loading config
	cfg := config.MustLoad()

	// logger initialization
	log := logger.NewLogger(cfg.LoggerLevel())
	log.Debug("debug messages are enabled")

	// database connection
	postgresConn := pg.New(cfg.PostgresDSN())
	if err := postgresConn.Ping(ctx); err != nil {
		lg.Fatalf("failed to connect to database: %v", err)
	}
	defer postgresConn.Close()

	log.Info("connected to database")

	// layers initialization
	entityRepository := postgres.NewEntityRepository(postgresConn, log)
	entityUsecase := usecase.NewEntityUsecase(entityRepository, log)
	v1Handler := v1.New(entityUsecase, log)

	// router initialization
	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)
	router.Route("/v1/entity", func(r chi.Router) {
		r.Post("/", v1Handler.CreateEntity(ctx))
		r.Get("/{id}", v1Handler.GetEntityByID(ctx))
	})

	// start of the server
	server := httpserver.New(cfg.ServerAddr(), router)
	server.Run()

	// graceful shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt

	log.Info("shutting down server")

	shutdownCtx, cancelShutdownCtx := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancelShutdownCtx()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Info("error while shutting down server", logger.Error(err))
	}

	log.Info("server stopped")
}
