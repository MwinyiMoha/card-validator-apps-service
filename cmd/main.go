package main

import (
	"card-validator-apps-service/internal/adaptors/api"
	"card-validator-apps-service/internal/adaptors/repository"
	"card-validator-apps-service/internal/config"
	"card-validator-apps-service/internal/core/application"
	protos "card-validator-apps-service/internal/gen"
	"card-validator-apps-service/internal/helpers"
	"card-validator-apps-service/internal/validation"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/mwinyimoha/card-validator-utils/logging"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	logger, err := logging.NewLoggerConfig().BuildLogger()
	if err != nil {
		log.Fatal(err)
	}

	defer logger.Sync()

	cfg, err := config.New()
	if err != nil {
		logger.Fatal("could not initialize app config", zap.String("original_error", err.Error()))
	}

	repo, err := repository.NewRepository(cfg)
	if err != nil {
		logger.Fatal("could not initialize app repository", zap.String("original_error", err.Error()))
	}

	val := validation.New()
	svc := application.NewService(repo, val)
	srv := api.NewServer(svc)

	s := grpc.NewServer()
	protos.RegisterAppsServiceServer(s, srv)
	reflection.Register(s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.ServerPort))
	if err != nil {
		logger.Fatal("could not create network listener", zap.String("original_error", err.Error()))
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	go func(c chan os.Signal) {
		logger.Info(
			"starting gRPC server",
			zap.String("service_name", cfg.ServiceName),
			zap.String("service_version", cfg.ServiceVersion),
			zap.Int("port", cfg.ServerPort),
		)

		if err := s.Serve(lis); err != nil {
			logger.Error("could not start server", zap.String("original_error", err.Error()))
			c <- syscall.SIGTERM
		}
	}(ch)

	received := <-ch

	func() {
		logger.Info("initiating graceful shutdown", zap.String("OS signal", received.String()))

		s.GracefulStop()

		ctx, cancel := helpers.NewTimeoutContext(cfg.DefaultTimeout)
		defer cancel()

		if err := repo.Disconnect(ctx); err != nil {
			logger.Fatal("could not close database connection", zap.String("original_error", err.Error()))
		}
	}()
}
