package main

import (
	"context"
	"httpsd-service/internal/config"
	"httpsd-service/internal/db"
	"httpsd-service/internal/db/metrics"
	"httpsd-service/internal/httpserver"
	"httpsd-service/internal/log"
	"httpsd-service/internal/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"

	v1 "httpsd-service/internal/controller/http/v1"
)

func main() {
	// Init logger
	logger := log.NewProductionLogger(nil)
	printLogo(logger)

	// Init config
	cfg, err := config.InitConfig()
	if err != nil {
		logger.Sugar().Fatalf("error init config: %v", err)
	}

	// Init database
	database, err := db.InitDataBase(context.Background(), cfg.RepositoryCfg)
	if err != nil {
		logger.Sugar().Fatalf("error init database: %v", err)
	}

	// Init Manager
	storage := metrics.NewStorage(database)
	manager := service.NewManager(logger, storage, []int64{1418862576}) // TODO: whitelist from config

	// Start HTTP server
	var mux = http.NewServeMux()
	v1.HandleRouts(mux, manager, logger)
	httpServer := httpserver.New(mux, httpserver.Port(cfg.ServicePort))
	logger.Sugar().Infof("server started on %s port", cfg.ServicePort)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Sugar().Warnf("app - Run - signal: %s", s.String())
	case err = <-httpServer.Notify():
		logger.Sugar().Warnf("app - Run - httpServer.Notify: %v", err)
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		logger.Sugar().Warnf("app - Run - httpServer.Shutdown: %v", err)
		return
	}
	logger.Sugar().Info("server shutdown")
}

func printLogo(logger *zap.Logger) {
	log.ClearTerminal()

	log.PrintLogo("AMS Service", []string{"DC71F5"})

	logger.Info("httpsd service is starting")
}
