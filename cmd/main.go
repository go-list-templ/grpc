package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-list-templ/grpc/config"
	"github.com/go-list-templ/grpc/internal/controller/grpc"
	"github.com/go-list-templ/grpc/internal/controller/http"
	"github.com/go-list-templ/grpc/internal/infra/persistent/postgres"
	"github.com/go-list-templ/grpc/internal/infra/persistent/redis"
	"go.uber.org/zap"
)

func main() {
	if err := run(); err != nil {
		log.Panic(err)
	}
}

// nolint:errcheck
func run() error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("starting app")
	logger.Info("initializing config")

	cfg, err := config.Load()
	if err != nil {
		logger.Panic("cant init config", zap.Error(err))
	}

	logger.Info("initializing postgres")

	pg, err := postgres.New(&cfg.DB)
	if err != nil {
		logger.Panic("cant init postgres", zap.Error(err))
	}
	pg.Close()

	logger.Info("initializing redis")

	rd, err := redis.New(&cfg.Redis)
	if err != nil {
		logger.Panic("cant init redis", zap.Error(err))
	}
	defer func() {
		if err = rd.Close(); err != nil {
			logger.Error("redis close failed", zap.Error(err))
		}
	}()

	logger.Info("initializing servers")

	grpcServer := grpc.NewServer(&cfg.Server)
	grpcServer.Start()

	healthServer := http.NewHealthServer(&cfg.Server)
	healthServer.Start()

	logger.Info("server started successfully")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case x := <-interrupt:
		logger.Info("Received a signal.", zap.String("signal", x.String()))
	case err = <-healthServer.Notify():
		logger.Error("Received an error from the health server", zap.Error(err))
	case err = <-grpcServer.Notify():
		logger.Error("Received an error from the grpc server", zap.Error(err))
	}

	logger.Info("stopping servers")

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Server.ShutdownTimeout)
	defer cancel()

	grpcServer.Stop()

	if err = healthServer.Stop(ctx); err != nil {
		logger.Error("server stopped with error", zap.Error(err))
	}

	logger.Info("The app is calling the last defers and will be stopped")

	return nil
}
