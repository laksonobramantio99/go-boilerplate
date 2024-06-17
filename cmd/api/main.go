package main

import (
	"fmt"
	"go-boilerplate/cmd/api/router"
	"go-boilerplate/config"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	env := "dev" // or "prod"

	// init config
	err := config.LoadConfig(env)
	if err != nil {
		log.Fatal().Msgf("config.LoadConfig: %v", err)
	}

	// init gin handler
	r := gin.Default()
	router.SetupRoutes(r)
	r.Run(fmt.Sprintf(":%d", config.Config.Port))
}
