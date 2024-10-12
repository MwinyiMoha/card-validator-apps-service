package main

import (
	"card-validator-apps-service/internal/config"
	"log"

	"github.com/mwinyimoha/card-validator-utils/logging"
	"go.uber.org/zap"
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

	logger.Info(cfg.ServiceName)
}
