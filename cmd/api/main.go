package main

import (
	"flag"
	"fmt"
	"go-boilerplate/cmd/api/router"
	"go-boilerplate/config"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	env := flag.String("env", "dev", "Environment to run the application in (dev or prod)")
	flag.Parse()

	// init config
	err := config.InitConfig(*env)
	if err != nil {
		log.Fatal().Msgf("config.LoadConfig: %v", err)
	}

	// init gin handler
	switch config.Config.Env {
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	router.SetupRoutes(r)

	log.Info().Msgf("[API Server] started on port :%d", config.Config.Port)
	r.Run(fmt.Sprintf(":%d", config.Config.Port))
}
