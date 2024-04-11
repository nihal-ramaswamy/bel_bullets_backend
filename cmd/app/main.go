package main

import (
	healthcheck_api "bel_bullets/internal/api/health_check"
	server_config "bel_bullets/internal/server"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	server_config := server_config.Default()
	gin.SetMode(server_config.GinMode)
	server := gin.Default()
	server.Use(cors.New(server_config.Cors))

	// Routes
	server.GET("/health_check", healthcheck_api.HealthCheckHandler())

	server.Run(server_config.Port)
}
