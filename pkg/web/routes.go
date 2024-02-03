package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	controller "go-jwt/pkg/web/controller/users"
	"go-jwt/pkg/web/middleware"
)

func StartWebServer(port string) {
	r := gin.Default()

	r.POST("/signup", controller.SignUp)
	r.POST("/login", controller.Login)
	r.GET("/validate", middleware.IsAuthorized, controller.Validate)

	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Error().Msgf("Failed to start webserver: %s", err.Error())
	}
}
