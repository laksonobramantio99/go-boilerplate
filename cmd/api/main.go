package main

import (
	"context"
	"flag"
	"fmt"
	"go-boilerplate/client/postgres"
	"go-boilerplate/client/redis"
	"go-boilerplate/cmd/api/middleware"
	"go-boilerplate/cmd/api/router"
	"go-boilerplate/config"
	"go-boilerplate/handler"
	"go-boilerplate/repo"
	"go-boilerplate/usecase"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	env := flag.String("env", "dev", "Environment to run the application in (dev or prod)")
	flag.Parse()

	// Init log
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Init config
	err := config.InitConfig(*env)
	if err != nil {
		log.Fatal().Msgf("config.LoadConfig: %v", err)
	}

	// Init repo, usecase, handler, client and others dependency
	mainHandler := initDependency()

	// Run HTTP server
	switch config.Config.Env {
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.Use(middleware.CorrelationIDMiddleware())
	router.SetupRoutes(r, mainHandler)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Config.Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Msgf("listen: %s", err)
		}
	}()
	log.Info().Msgf("[API Server] Started on port :%d", config.Config.Port)

	gracefulShutdown(srv)
}

func initDependency() *handler.MainHandler {
	// Init postgres DB
	db, err := postgres.InitPostgres()
	if err != nil {
		log.Fatal().Msgf("postgres.InitPostgres: %v", err)
	}

	// Init Redis
	_, err = redis.InitRedisClient()
	if err != nil {
		log.Fatal().Msgf("redis.InitRedisClient: %v", err)
	}

	// Init repo
	bookRepo := repo.NewBookRepo(postgres.DB.Master, db.Slave)

	// Init usecase
	bookUc := usecase.NewUsecase(bookRepo)

	// Init handler
	bookHandler := handler.NewBookHandler(bookUc)

	mainHandler := handler.NewMainHandler(bookHandler)
	return mainHandler
}

func gracefulShutdown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("[API Server] Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("[API Server] Server forced to shutdown: %v", err)
	}

	log.Info().Msg("[API Server] Server exiting")
}
