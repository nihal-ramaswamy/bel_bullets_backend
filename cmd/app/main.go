package main

import (
	clubsApi "bel_bullets/internal/api/clubs"
	healthcheck_api "bel_bullets/internal/api/health_check"
	middleware "bel_bullets/internal/api/middleware"
	server_config "bel_bullets/internal/server"
	utils "bel_bullets/internal/utils"

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

	strava := server.Group("/strava", middleware.StravaMiddleWare())
	{
		strava.GET("/getClubById", clubsApi.GetClubByIdHandler(utils.GetDotEnvVariable("CLUB_ID")))
	}

	server.Run(server_config.Port)
}
