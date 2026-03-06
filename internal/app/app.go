package app

import (
	"context"
	lg "log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/solumD/go-service-template/config"
	"github.com/solumD/go-service-template/internal/repository/postgres"
	"github.com/solumD/go-service-template/internal/transport"
	"github.com/solumD/go-service-template/internal/transport/handler"
	"github.com/solumD/go-service-template/internal/usecase"
	httpserver "github.com/solumD/go-service-template/pkg/http_server"
	"github.com/solumD/go-service-template/pkg/logger"
	pg "github.com/solumD/go-service-template/pkg/postgres"
)

const shutdownTimeout = 10 * time.Second

func InitAndRun(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	cfg := config.MustLoad()

	log := logger.NewLogger(cfg.LoggerLevel())
	log.Debug("debug messages are enabled")

	postgresConn := pg.New(cfg.PostgresDSN())
	if err := postgresConn.Ping(ctx); err != nil {
		lg.Fatalf("failed to connect to postgres: %v", err)
	}
	defer postgresConn.Close()

	log.Info("connected to postgres")

	entityRepository := postgres.NewEntityRepository(postgresConn, log)
	entityUsecase := usecase.NewEntityUsecase(entityRepository, log)
	handler := handler.New(entityUsecase, log)

	router := transport.NewRouter(ctx, log, handler)

	server := httpserver.New(cfg.ServerAddr(), router)
	server.Run()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt

	log.Info("shutting down server...")

	shutdownCtx, cancelShutdownCtx := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancelShutdownCtx()

	err := server.Shutdown(shutdownCtx)
	if err != nil {
		log.Info("error while shutting down server", logger.Error(err))
	}
}
